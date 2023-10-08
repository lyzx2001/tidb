// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package globalstats

import (
	"fmt"

	"github.com/pingcap/errors"
	"github.com/pingcap/tidb/infoschema"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/model"
	"github.com/pingcap/tidb/sessionctx"
	"github.com/pingcap/tidb/statistics"
	"github.com/pingcap/tidb/table"
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/logutil"
	"github.com/tiancaiamao/gp"
	"go.uber.org/zap"
)

const (
	// TiDBGlobalStats represents the global-stats for a partitioned table.
	TiDBGlobalStats = "global"
	// MaxPartitionMergeBatchSize indicates the max batch size for a worker to merge partition stats
	MaxPartitionMergeBatchSize = 256
)

// GlobalStats is used to store the statistics contained in the global-level stats
// which is generated by the merge of partition-level stats.
// It will both store the column stats and index stats.
// In the column statistics, the variable `num` is equal to the number of columns in the partition table.
// In the index statistics, the variable `num` is always equal to one.
type GlobalStats struct {
	Hg                    []*statistics.Histogram
	Cms                   []*statistics.CMSketch
	TopN                  []*statistics.TopN
	Fms                   []*statistics.FMSketch
	MissingPartitionStats []string
	Num                   int
	Count                 int64
	ModifyCount           int64
}

func newGlobalStats(histCount int) *GlobalStats {
	globalStats := new(GlobalStats)
	globalStats.Num = histCount
	globalStats.Count = 0
	globalStats.Hg = make([]*statistics.Histogram, globalStats.Num)
	globalStats.Cms = make([]*statistics.CMSketch, globalStats.Num)
	globalStats.TopN = make([]*statistics.TopN, globalStats.Num)
	globalStats.Fms = make([]*statistics.FMSketch, globalStats.Num)

	return globalStats
}

type (
	getTableByPhysicalIDFunc    func(is infoschema.InfoSchema, physicalID int64) (table.Table, bool)
	loadTablePartitionStatsFunc func(tableInfo *model.TableInfo, partitionDef *model.PartitionDefinition) (*statistics.Table, error)
	// GlobalStatusHandler is used to handle the global-level stats.
)

// MergePartitionStats2GlobalStats merge the partition-level stats to global-level stats based on the tableInfo.
func MergePartitionStats2GlobalStats(
	sc sessionctx.Context,
	gpool *gp.Pool,
	opts map[ast.AnalyzeOptionType]uint64,
	is infoschema.InfoSchema,
	globalTableInfo *model.TableInfo,
	isIndex bool,
	histIDs []int64,
	allPartitionStats map[int64]*statistics.Table,
	getTableByPhysicalIDFn getTableByPhysicalIDFunc,
	loadTablePartitionStatsFn loadTablePartitionStatsFunc,
) (globalStats *GlobalStats, err error) {
	externalCache := false
	if allPartitionStats != nil {
		externalCache = true
	}

	partitionNum := len(globalTableInfo.Partition.Definitions)
	if len(histIDs) == 0 {
		for _, col := range globalTableInfo.Columns {
			// The virtual generated column stats can not be merged to the global stats.
			if col.IsVirtualGenerated() {
				continue
			}
			histIDs = append(histIDs, col.ID)
		}
	}

	// Initialized the globalStats.
	globalStats = newGlobalStats(len(histIDs))

	// Slice Dimensions Explanation
	// First dimension: Column or Index Stats
	// Second dimension: Partition Tables
	// Because all topN and histograms need to be collected before they can be merged.
	// So we should store all the partition-level stats first, and merge them together.
	allHg := make([][]*statistics.Histogram, globalStats.Num)
	allCms := make([][]*statistics.CMSketch, globalStats.Num)
	allTopN := make([][]*statistics.TopN, globalStats.Num)
	allFms := make([][]*statistics.FMSketch, globalStats.Num)
	for i := 0; i < globalStats.Num; i++ {
		allHg[i] = make([]*statistics.Histogram, 0, partitionNum)
		allCms[i] = make([]*statistics.CMSketch, 0, partitionNum)
		allTopN[i] = make([]*statistics.TopN, 0, partitionNum)
		allFms[i] = make([]*statistics.FMSketch, 0, partitionNum)
	}

	skipMissingPartitionStats := sc.GetSessionVars().SkipMissingPartitionStats
	for _, def := range globalTableInfo.Partition.Definitions {
		partitionID := def.ID
		partitionTable, ok := getTableByPhysicalIDFn(is, partitionID)
		if !ok {
			err = errors.Errorf("unknown physical ID %d in stats meta table, maybe it has been dropped", partitionID)
			return
		}
		tableInfo := partitionTable.Meta()
		var partitionStats *statistics.Table
		var okLoad bool
		if allPartitionStats != nil {
			partitionStats, okLoad = allPartitionStats[partitionID]
		} else {
			okLoad = false
		}
		// If pre-load partition stats isn't provided, then we load partition stats directly and set it into allPartitionStats
		if !okLoad {
			var err1 error
			partitionStats, err1 = loadTablePartitionStatsFn(tableInfo, &def)
			if err1 != nil {
				if skipMissingPartitionStats && types.ErrPartitionStatsMissing.Equal(err1) {
					globalStats.MissingPartitionStats = append(globalStats.MissingPartitionStats, fmt.Sprintf("partition `%s`", def.Name.L))
					continue
				}
				err = err1
				return
			}
			if externalCache {
				allPartitionStats[partitionID] = partitionStats
			}
		}

		for i := 0; i < globalStats.Num; i++ {
			// GetStatsInfo will return the copy of the statsInfo, so we don't need to worry about the data race.
			// partitionStats will be released after the for loop.
			hg, cms, topN, fms, analyzed := partitionStats.GetStatsInfo(histIDs[i], isIndex, externalCache)
			skipPartition := false
			if !analyzed {
				var missingPart string
				if !isIndex {
					missingPart = fmt.Sprintf("partition `%s` column `%s`", def.Name.L, tableInfo.FindColumnNameByID(histIDs[i]))
				} else {
					missingPart = fmt.Sprintf("partition `%s` index `%s`", def.Name.L, tableInfo.FindIndexNameByID(histIDs[i]))
				}
				if !skipMissingPartitionStats {
					err = types.ErrPartitionStatsMissing.GenWithStackByArgs(fmt.Sprintf("table `%s` %s", tableInfo.Name.L, missingPart))
					return
				}
				globalStats.MissingPartitionStats = append(globalStats.MissingPartitionStats, missingPart)
				skipPartition = true
			}

			// Partition stats is not empty but column stats(hist, topN) is missing.
			if partitionStats.RealtimeCount > 0 && (hg == nil || hg.TotalRowCount() <= 0) && (topN == nil || topN.TotalCount() <= 0) {
				var missingPart string
				if !isIndex {
					missingPart = fmt.Sprintf("partition `%s` column `%s`", def.Name.L, tableInfo.FindColumnNameByID(histIDs[i]))
				} else {
					missingPart = fmt.Sprintf("partition `%s` index `%s`", def.Name.L, tableInfo.FindIndexNameByID(histIDs[i]))
				}
				if !skipMissingPartitionStats {
					err = types.ErrPartitionColumnStatsMissing.GenWithStackByArgs(fmt.Sprintf("table `%s` %s", tableInfo.Name.L, missingPart))
					return
				}
				globalStats.MissingPartitionStats = append(globalStats.MissingPartitionStats, missingPart+" hist and topN")
				skipPartition = true
			}

			if i == 0 {
				// In a partition, we will only update globalStats.Count once.
				globalStats.Count += partitionStats.RealtimeCount
				globalStats.ModifyCount += partitionStats.ModifyCount
			}

			if !skipPartition {
				allHg[i] = append(allHg[i], hg)
				allCms[i] = append(allCms[i], cms)
				allTopN[i] = append(allTopN[i], topN)
				allFms[i] = append(allFms[i], fms)
			}
		}
	}

	// After collect all the statistics from the partition-level stats,
	// we should merge them together.
	for i := 0; i < globalStats.Num; i++ {
		if len(allHg[i]) == 0 {
			// If all partitions have no stats, we skip merging global stats because it may not handle the case `len(allHg[i]) == 0`
			// correctly. It can avoid unexpected behaviors such as nil pointer panic.
			continue
		}
		// FMSketch use many memory, so we first deal with it and then destroy it.
		// Merge FMSketch.
		globalStats.Fms[i] = allFms[i][0]
		for j := 1; j < len(allFms[i]); j++ {
			globalStats.Fms[i].MergeFMSketch(allFms[i][j])
			allFms[i][j].DestroyAndPutToPool()
		}

		// Update the global NDV.
		globalStatsNDV := globalStats.Fms[i].NDV()
		if globalStatsNDV > globalStats.Count {
			globalStatsNDV = globalStats.Count
		}
		globalStats.Fms[i].DestroyAndPutToPool()

		// Merge CMSketch.
		globalStats.Cms[i] = allCms[i][0]
		for j := 1; j < len(allCms[i]); j++ {
			err = globalStats.Cms[i].MergeCMSketch(allCms[i][j])
			if err != nil {
				return
			}
		}

		// Merge topN.
		// Note: We need to merge TopN before merging the histogram.
		// Because after merging TopN, some numbers will be left.
		// These remaining topN numbers will be used as a separate bucket for later histogram merging.
		var poppedTopN []statistics.TopNMeta
		wrapper := NewStatsWrapper(allHg[i], allTopN[i])
		globalStats.TopN[i], poppedTopN, allHg[i], err = mergeGlobalStatsTopN(gpool, sc, wrapper,
			sc.GetSessionVars().StmtCtx.TimeZone, sc.GetSessionVars().AnalyzeVersion, uint32(opts[ast.AnalyzeOptNumTopN]), isIndex)
		if err != nil {
			return
		}

		// Merge histogram.
		globalStats.Hg[i], err = statistics.MergePartitionHist2GlobalHist(sc.GetSessionVars().StmtCtx, allHg[i], poppedTopN,
			int64(opts[ast.AnalyzeOptNumBuckets]), isIndex)
		if err != nil {
			return
		}

		// NOTICE: after merging bucket NDVs have the trend to be underestimated, so for safe we don't use them.
		for j := range globalStats.Hg[i].Buckets {
			globalStats.Hg[i].Buckets[j].NDV = 0
		}

		globalStats.Hg[i].NDV = globalStatsNDV
	}
	return
}

// MergePartitionStats2GlobalStatsByTableID merge the partition-level stats to global-level stats based on the tableID.
func MergePartitionStats2GlobalStatsByTableID(
	sc sessionctx.Context,
	gpool *gp.Pool,
	opts map[ast.AnalyzeOptionType]uint64,
	is infoschema.InfoSchema,
	physicalID int64,
	isIndex bool,
	histIDs []int64,
	allPartitionStats map[int64]*statistics.Table,
	getTableByPhysicalIDFn getTableByPhysicalIDFunc,
	loadTablePartitionStatsFn loadTablePartitionStatsFunc,
) (globalStats *GlobalStats, err error) {
	// Get the partition table IDs.
	globalTable, ok := getTableByPhysicalIDFn(is, physicalID)
	if !ok {
		err = errors.Errorf("unknown physical ID %d in stats meta table, maybe it has been dropped", physicalID)
		return
	}

	globalTableInfo := globalTable.Meta()
	globalStats, err = MergePartitionStats2GlobalStats(sc, gpool, opts, is, globalTableInfo, isIndex, histIDs, allPartitionStats, getTableByPhysicalIDFn, loadTablePartitionStatsFn)
	if err != nil {
		return
	}

	if len(globalStats.MissingPartitionStats) > 0 {
		var item string
		if !isIndex {
			item = "columns"
		} else {
			item = "index"
			if len(histIDs) > 0 {
				item += " " + globalTableInfo.FindIndexNameByID(histIDs[0])
			}
		}

		logutil.BgLogger().Warn("missing partition stats when merging global stats", zap.String("table", globalTableInfo.Name.L),
			zap.String("item", item), zap.Strings("missing", globalStats.MissingPartitionStats))
	}

	return
}
