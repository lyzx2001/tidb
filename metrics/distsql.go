// Copyright 2018 PingCAP, Inc.
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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// distsql metrics.
var (
	DistSQLQueryHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "handle_query_duration_seconds",
			Help:      "Bucketed histogram of processing time (s) of handled queries.",
			Buckets:   prometheus.ExponentialBuckets(0.0005, 2, 29), // 0.5ms ~ 1.5days
		}, []string{LblType, LblSQLType, LblCoprType})

	DistSQLScanKeysPartialHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "scan_keys_partial_num",
			Help:      "number of scanned keys for each partial result.",
		},
	)
	DistSQLScanKeysHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "scan_keys_num",
			Help:      "number of scanned keys for each query.",
		},
	)
	DistSQLPartialCountHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "partial_num",
			Help:      "number of partial results for each query.",
		},
	)
	DistSQLCoprCacheCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "copr_cache",
			Help:      "coprocessor cache hit, evict and miss number",
		}, []string{LblType})
	DistSQLCoprClosestReadCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "copr_closest_read",
			Help:      "counter of total copr read local read hit.",
		}, []string{LblType})
	DistSQLCoprRespBodySize = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "tidb",
			Subsystem: "distsql",
			Name:      "copr_resp_size",
			Help:      "copr task response data size in bytes.",
			Buckets:   prometheus.ExponentialBuckets(1024, 2, 20),
		}, []string{LblStore})
)
