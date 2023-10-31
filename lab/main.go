package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Metrics
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.ExponentialBuckets(0.005, 1.5, 10), // Adjust buckets as needed
		},
		[]string{"method", "endpoint"},
	)

	// Custom metrics (add more as needed)
	customMetric = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "custom_metric",
			Help: "Custom metric example.",
		},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(customMetric)
	prometheus.MustRegister(requestDuration)
}

func main() {
	// Create a Gin router
	r := gin.Default()

	// Define a route for fetching car park data
	r.GET("/carparks", func(c *gin.Context) {
		// Simulate some processing time
		start := time.Now()
		time.Sleep(1000 * time.Millisecond)

		// Increment the request counter
		requestCounter.WithLabelValues("GET", "/carparks", "200").Inc()

		// You can also record custom metrics
		customMetric.Set(42.0)

		elapsed := time.Since(start).Seconds()

		requestDuration.WithLabelValues("GET", "/carparks").Observe(elapsed)

		// Serialize and return car park data as JSON
		c.JSON(http.StatusOK, []string{})
	})

	r.GET("/server-error", func(c *gin.Context) {
		// Simulate an intentional error (e.g., 500 Internal Server Error)
		c.JSON(http.StatusInternalServerError, "Internal Server Error")

		// Increment the errorCounter metric
		requestCounter.WithLabelValues("GET", "/server-error", "500").Inc()
	})

	// Define a route for handling errors (e.g., 404 Not Found)
	r.NoRoute(func(c *gin.Context) {
		// Increment the error counter
		requestCounter.WithLabelValues("GET", c.Request.URL.Path, "404").Inc()

		// Return an error response
		c.JSON(http.StatusNotFound, "Not Found")
	})

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start your Gin server
	r.Run(":8080")
}
