package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RequestCount exposes total number of request
// PromQL
// rate(http_requests_total[1m])
var RequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "How many HTTP requests processed, partitioned by status code, path and HTTP method.",
	},
	[]string{"code", "path", "method"},
)
