package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RequestLatencySummary exposes total number of request
// PromQL
// rate(http_response_latency_seconds_sum[1m]) / rate(http_response_latency_seconds_count[1m])
var RequestLatencySummary = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_response_latency_seconds",
		Help: "Response latency in seconds",
	},
	[]string{"code", "path", "method"},
)
