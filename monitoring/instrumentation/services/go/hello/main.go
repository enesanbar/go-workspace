package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)


const Port = 8000

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", handler).Methods("GET")
	router.Path("/metrics").Handler(promhttp.Handler())
	router.Use(metricsMiddleware)

	//prometheus.MustRegister(RequestCount, RequestInProgress, RequestLatencySummary)
	prometheus.MustRegister(RequestCount, RequestInProgress, RequestLatencyHistogram)

	log.Println("Starting application on", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), router))
}
