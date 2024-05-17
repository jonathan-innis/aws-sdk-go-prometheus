# AWS SDK Go Prometheus Metrics

### v1

```go
// Register Prometheus metrics with the sigs.k8s.io/controller-runtime/pkg/metrics registry
sess := v1.WithPrometheusMetrics(session.Must(session.NewSession()), metrics.Registry)
```

### v2

```go
ctx := context.Background()

// Register Prometheus metrics with the sigs.k8s.io/controller-runtime/pkg/metrics registry
cfg = v2.WithPrometheusMetrics(lo.Must(config.LoadDefaultConfig(ctx)), metrics.Registry)
```