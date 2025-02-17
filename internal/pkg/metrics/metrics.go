// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ashwinyue/onex.
//

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	DefaultMetricsAddress           = ":8081"
	DefaultMinerSetMetricsAddress   = ":8082"
	DefaultMinerMetricsAddress      = ":8083"
	DefaultNodeServerMetricsAddress = ":8084"
)

var (
	// MinerCountDesc is a metric about miner object count in the cluster.
	MinerCountDesc = prometheus.NewDesc("napi_miner_items", "Count of miner objects currently at the apiserver", nil, nil)
	// MinerSetCountDesc Count of minerset object count at the apiserver.
	MinerSetCountDesc = prometheus.NewDesc("napi_minerset_items", "Count of minersets at the apiserver", nil, nil)
	// MinerInfoDesc is a metric about miner object info in the cluster.
	MinerInfoDesc = prometheus.NewDesc("napi_miner_created_timestamp_seconds", "Timestamp of the napi managed Miner creation time", []string{"name", "namespace", "spec_provider_id", "node", "api_version", "phase"}, nil)
	// MinerSetInfoDesc is a metric about miner object info in the cluster.
	MinerSetInfoDesc = prometheus.NewDesc("napi_minerset_created_timestamp_seconds", "Timestamp of the napi managed Minerset creation time", []string{"name", "namespace", "api_version"}, nil)

	// MinerSetStatusAvailableReplicasDesc is the information of the Minerset's status for available replicas.
	MinerSetStatusAvailableReplicasDesc = prometheus.NewDesc("napi_miner_set_status_replicas_available", "Information of the napi managed Minerset's status for available replicas", []string{"name", "namespace"}, nil)

	// MinerSetStatusReadyReplicasDesc is the information of the Minerset's status for ready replicas.
	MinerSetStatusReadyReplicasDesc = prometheus.NewDesc("napi_miner_set_status_replicas_ready", "Information of the napi managed Minerset's status for ready replicas", []string{"name", "namespace"}, nil)

	// MinerSetStatusReplicasDesc is the information of the Minerset's status for replicas.
	MinerSetStatusReplicasDesc = prometheus.NewDesc("napi_miner_set_status_replicas", "Information of the napi managed Minerset's status for replicas", []string{"name", "namespace"}, nil)

	// MinerCollectorUp is a Prometheus metric, which reports reflects successful collection and reporting of all the metrics.
	MinerCollectorUp = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "napi_mao_collector_up",
		Help: "Node API Controller metrics are being collected and reported successfully",
	}, []string{"kind"})

	failedInstanceCreateCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "napi_instance_create_failed",
			Help: "Number of times provider instance create has failed.",
		}, []string{"name", "namespace", "reason"},
	)

	failedInstanceUpdateCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "napi_instance_update_failed",
			Help: "Number of times provider instance update has failed.",
		}, []string{"name", "namespace", "reason"},
	)

	failedInstanceDeleteCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "napi_instance_delete_failed",
			Help: "Number of times provider instance delete has failed.",
		}, []string{"name", "namespace", "reason"},
	)
)

func init() {
	prometheus.MustRegister(KratosMetricSeconds, KratosServerMetricRequests, MinerCollectorUp)
	metrics.Registry.MustRegister(
		failedInstanceCreateCount,
		failedInstanceUpdateCount,
		failedInstanceDeleteCount,
	)
}
