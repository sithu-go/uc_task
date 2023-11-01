package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	RequestCounter  *prometheus.CounterVec
	ErrorCounter    *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec

	CronErrorCounter *prometheus.CounterVec
}

var (
	Metrics *Metric
)

func NewMetric() {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	errorCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total number of HTTP request errors.",
		},
		[]string{"method", "endpoint", "status_code"},
	)

	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.ExponentialBuckets(0.005, 1.5, 10), // Adjust buckets as needed
		},
		[]string{"method", "endpoint"},
	)

	cronErrorCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cron_errors_total",
			Help: "Total number of cron job errors and messages",
		},
		[]string{"job_name", "error_message", "data"},
	)

	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(errorCounter)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(cronErrorCounter)

	Metrics = &Metric{
		RequestCounter:   requestCounter,
		ErrorCounter:     errorCounter,
		RequestDuration:  requestDuration,
		CronErrorCounter: cronErrorCounter,
	}
}
