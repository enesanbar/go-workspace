package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		route := mux.CurrentRoute(request)
		template, _ := route.GetPathTemplate()

		RequestInProgress.WithLabelValues("200", template, request.Method).Inc()

		next.ServeHTTP(writer, request)

		RequestCount.WithLabelValues("200", template, request.Method).Inc()
		RequestInProgress.WithLabelValues("200", template, request.Method).Dec()
		//RequestLatencySummary.WithLabelValues("200", template, request.Method).Observe(time.Since(startTime).Seconds())
		RequestLatencyHistogram.WithLabelValues("200", template, request.Method).Observe(time.Since(startTime).Seconds())
	})
}

