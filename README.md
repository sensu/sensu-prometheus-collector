## Sensu Prometheus Collector

The Sensu Prometheus Collector is a Sensu Check Plugin that collects
metrics from a Prometheus exporter or the Prometheus query API. The
metrics are outputted to STDOUT in one of three formats, Influx (the
default), Graphite, or JSON.

### Examples

```shell
$sensu-prometheus-collector -exporter-url http://localhost:8080/metrics
go_memstats_last_gc_time_seconds value=0.000000 1506991233
go_memstats_mspan_sys_bytes value=32768.000000 1506991233
...
```

```shell
$sensu-prometheus-collector -exporter-url http://localhost:8080/metrics -output-format graphite -metric-prefix foo.bar.
foo.bar.go_memstats_stack_inuse_bytes 294912.000000 1506991405
foo.bar.go_memstats_mallocs_total 6375.000000 1506991405
...
```

```shell
$sensu-prometheus-collector
up,instance=localhost:9090,job=prometheus value=1.000000 1506991495
```
