## Sensu Prometheus Collector

The Sensu Prometheus Collector is a Sensu Check Plugin that collects
metrics from a Prometheus exporter or the Prometheus query API. The
collected metrics are outputted to STDOUT in one of three formats,
Influx (the default), Graphite, or JSON.

The Sensu Prometheus Collector turns Sensu into a *SUPER POWERED*
Prometheus metric poller, leveraging Sensu's pubsub design and client
auto-registration (discovery).

### Examples

Help:

```
$ sensu-prometheus-collector -h
Usage of ./sensu-prometheus-collector:
  -exporter-url string
        Prometheus exporter URL to pull metrics from.
  -metric-prefix string
        Metric name prefix, only supported by line protocol output formats.
  -output-format string
        The check output format to use for metrics {influx|graphite|json}. (default "influx")
  -prom-query string
        Prometheus API query string. (default "up")
  -prom-url string
        Prometheus API URL. (default "http://localhost:9090")
```

Application instrumentation:

```
$ sensu-prometheus-collector -exporter-url http://localhost:8080/metrics
go_memstats_last_gc_time_seconds value=0 1506991233
go_memstats_mspan_sys_bytes value=32768 1506991233
...
```

```
$ sensu-prometheus-collector -exporter-url http://localhost:8080/metrics -output-format graphite -metric-prefix foo.bar.
foo.bar.go_memstats_stack_inuse_bytes 294912 1506991405
foo.bar.go_memstats_mallocs_total 6375 1506991405
...
```

Prometheus query API:

```
$ sensu-prometheus-collector -prom-url http://localhost:9090 -prom-query up
up,instance=localhost:9090,job=prometheus value=1 1506991495
```
