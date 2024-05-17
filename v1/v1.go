package v1

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jonathan-innis/aws-sdk-go-prometheus/common"
	"github.com/prometheus/client_golang/prometheus"
)

// WithPrometheusMetrics wraps an session.Session, injecting prometheus metric firing
// into the handlers to track request count totals, latencies, and retry counts
func WithPrometheusMetrics(sess *session.Session, r prometheus.Registerer) *session.Session {
	common.MustRegisterMetrics(r)
	sess.Handlers.Complete.PushBackNamed(PrometheusHandler)
	sess.Handlers.CompleteAttempt.PushBackNamed(PrometheusRetryHandler)
	return sess
}

// PrometheusHandler is a request handler to fire prometheus metrics on requests.
var PrometheusHandler = request.NamedHandler{Name: "PrometheusHandler", Fn: func(r *request.Request) {
	common.TotalRequests.With(common.RequestLabels(r.ClientInfo.ServiceID, r.Operation.Name, r.HTTPResponse.StatusCode)).Inc()
	common.RequestLatency.With(common.RequestLabels(r.ClientInfo.ServiceID, r.Operation.Name, r.HTTPResponse.StatusCode)).Observe(float64(time.Since(r.Time).Milliseconds()))
	common.RetryCount.With(common.RequestLabels(r.ClientInfo.ServiceID, r.Operation.Name, r.HTTPResponse.StatusCode)).Observe(float64(r.RetryCount))
}}

var PrometheusRetryHandler = request.NamedHandler{Name: "PrometheusRetryHandler", Fn: func(r *request.Request) {
	common.TotalRequestAttempts.With(common.RequestLabels(r.ClientInfo.ServiceID, r.Operation.Name, r.HTTPResponse.StatusCode)).Inc()
	common.RequestAttemptLatency.With(common.RequestLabels(r.ClientInfo.ServiceID, r.Operation.Name, r.HTTPResponse.StatusCode)).Observe(float64(time.Since(r.AttemptTime).Milliseconds()))
}}
