package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RequestLatencyHistogram exposes latency metrics
// PromQL
// rate(http_response_latency_seconds_sum[1m]) / rate(http_response_latency_seconds_count[1m])
var RequestLatencyHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_latency_seconds",
		Help: "Response latency in seconds",
	},
	[]string{"code", "path", "method"},
)
