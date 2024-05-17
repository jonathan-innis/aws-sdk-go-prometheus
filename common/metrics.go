package common

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/samber/lo"
)

var (
	labels        = []string{"service", "action", "status_code"}
	TotalRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "aws_sdk_go_requests_total",
		Help: "The total number of AWS SDK Go requests",
	}, labels)

	TotalRequestAttempts = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "aws_sdk_go_request_attempts_total",
		Help: "The total number of AWS SDK Go request attempts",
	}, labels)

	RequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "aws_sdk_go_request_latency",
		Help: "Latency of AWS SDK Go requests",
		Buckets: []float64{
			10, 20, 30, 40, 50, 60, 70, 80, 90, 100,
			125, 150, 175, 200, 225, 250, 275, 300,
			400, 500, 600, 700, 800, 900,
			1_000, 1_500, 2_000, 2_500, 3_000, 3_500, 4_000, 4_500, 5_000,
			6_000, 7_000, 8_000, 9_000, 10_000,
		},
	}, labels)

	RequestAttemptLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "aws_sdk_go_request_attempt_latency",
		Help: "Latency of AWS SDK Go request attempts",
		Buckets: []float64{
			10, 20, 30, 40, 50, 60, 70, 80, 90, 100,
			125, 150, 175, 200, 225, 250, 275, 300,
			400, 500, 600, 700, 800, 900,
			1_000, 1_500, 2_000, 2_500, 3_000, 3_500, 4_000, 4_500, 5_000,
			6_000, 7_000, 8_000, 9_000, 10_000,
		},
	}, labels)

	RetryCount = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "aws_sdk_go_request_retry_count",
		Help: "The total number of AWS SDK Go retry attempts per request",
		Buckets: []float64{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		},
	}, labels)
)

func RequestLabels(service string, action string, statusCode int) prometheus.Labels {
	return prometheus.Labels{
		"service":     service,
		"action":      action,
		"status_code": fmt.Sprint(statusCode),
	}
}

func MustRegisterMetrics(registry prometheus.Registerer) {
	for _, c := range []prometheus.Collector{TotalRequests, TotalRequestAttempts, RequestLatency, RequestAttemptLatency, RetryCount} {
		lo.Must0(registry.Register(c))
	}
}
