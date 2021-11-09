package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func renderMetricsResponse() (string, error) {
	var res string

	startTime := time.Now().UnixNano()

	status, err := getMetrics()
	if err != nil {
		log.Println("Unable to get metrics!", err)
		res = getFetchSuccessString("UNKNOWN", "UNKNOWN", 0) +
			getFetchDurationString("UNKNOWN", "UNKNOWN", time.Now().UnixNano()-startTime)

		return res, nil
	} else {
		res = getMetricsString(status) +
			getFetchSuccessString(status.UPSType, status.UPSSerial, 0) +
			getFetchDurationString(status.UPSType, status.UPSSerial, time.Now().UnixNano()-startTime)

		return res, nil
	}
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/metrics" {
		response, err := renderMetricsResponse()
		if err != nil {
			log.Println("Error fetching metrics!", err)
			w.WriteHeader(500)
			_, _ = fmt.Fprint(w, err)
			return
		}
		_, _ = fmt.Fprint(w, response)
	} else {
		log.Println("Not found: '" + r.RequestURI)
		w.WriteHeader(404)
	}
}

func main() {
	var listenAddress = flag.String("listenAddress", ":9122", "address the exporter will listen on")
	flag.Parse()

	server := &http.Server{
		Addr:         *listenAddress,
		Handler:      http.HandlerFunc(handleMetrics),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server on " + *listenAddress)
	err2 := server.ListenAndServe()
	if err2 != nil {
		log.Fatalln("Unable to start server!", err2)
	}
}
