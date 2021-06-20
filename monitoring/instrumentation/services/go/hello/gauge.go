package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RequestInProgress exposes total number of request
// PromQL
// http_requests_in_progress
var RequestInProgress = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "http_requests_in_progress",
		Help: "Numbe of HTTP request in progress, partitioned by status code, path and HTTP method.",
	},
	[]string{"code", "path", "method"},
)
