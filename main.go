package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	exampleGaugeMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "test",
		Name:      "example_gauge",
		Help:      "Example Gauge metric",
	}, []string{})
	exampleCounterMetric = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "test",
		Name:      "example_counter",
		Help:      "Example Counter metric",
	}, []string{})
	exampleCounterMetric2 = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "test",
		Name:      "example_counter2",
		Help:      "Example Counter2 metric",
	}, []string{})
)

func main() {
	var (
		addr = flag.String("listen-address", "0.0.0.0:8080", "The address to listen on for HTTP requests.")
	)
	flag.Parse()

	counter := float64(0)

	go func() {
		for {
			gauge := float64(rand.Intn(100))
			exampleGaugeMetric.WithLabelValues().Set(gauge)

			counterInc := float64(rand.Intn(100))
			counter += counterInc

			exampleCounterMetric.WithLabelValues().Add(counterInc)
			exampleCounterMetric2.WithLabelValues().Add(counterInc)

			fmt.Printf("gauge: %f, counterInc: %f, counter: %f\n", gauge, counterInc, counter)
			time.Sleep(10 * time.Second)
		}
	}()

	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.Handler())
	fmt.Printf("Listen on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
