[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Prometheus%20Collector-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sensu/sensu-prometheus-collector) [![Build Status](https://travis-ci.org/sensu/sensu-prometheus-collector.svg?branch=master)](https://travis-ci.org/sensu/sensu-prometheus-collector)
## Sensu Prometheus Collector

- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Sensu Go](#sensu-go)
    - [Asset definition](#asset-definition)
    - [Check definition(s)](#check-definitions)
  - [Sensu Core](#sensu-core)
    - [Check definition](#check-definition)
    - [Handler definition](#handler-definition)
- [Installation from source and contributing](#installation-from-source-and-contributing)

### Overview

The Sensu Prometheus Collector is a Sensu Check Plugin that collects
metrics from a Prometheus exporter or the Prometheus query API. The
collected metrics are outputted to STDOUT in one of three formats,
Influx (the default), Graphite, or JSON.

The Sensu Prometheus Collector turns Sensu into a *SUPER POWERED
Prometheus metric poller*, leveraging Sensu's pubsub design and client
auto-registration (discovery). Sensu can deliver collected metrics to
one or more time-series databases, for example InfluxDB and/or
Graphite! Instrument your applications with the Prometheus libraries
and immediately begin collecting your metrics with Sensu!

### Usage examples


Help:

```
$ sensu-prometheus-collector -h
Usage of ./sensu-prometheus-collector:
  -exporter-authorization string
        Prometheus exporter Authorization header.
  -exporter-password string
        Prometheus exporter basic auth password.
  -exporter-url string
        Prometheus exporter URL to pull metrics from.
  -exporter-user string
        Prometheus exporter basic auth user.
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
$ sensu-prometheus-collector -exporter-url http://localhost:8080/metrics -exporter-user admin -exporter-password secretpassword -output-format graphite -metric-prefix foo.bar.
foo.bar.go_memstats_stack_inuse_bytes 294912 1506991405
foo.bar.go_memstats_mallocs_total 6375 1506991405
...
```

Exporter basic auth credentials can also be set via environment vars `EXPORTER_USER` and `EXPORTER_PASSWORD`.

Prometheus query API:

```
$ sensu-prometheus-collector -prom-url http://localhost:9090 -prom-query up
up,instance=localhost:9090,job=prometheus value=1 1506991495
```

### Asset registration

Assets are the best way to make use of this handler. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset: 

`sensuctl asset add sensu/sensu-prometheus-collector`

If you're using an earlier version of sensuctl, you can download the asset definition from [this project's Bonsai Asset Index page](https://bonsai.sensu.io/assets/sensu/sensu-prometheus-collector).

### Configuration

#### Sensu Go

##### Asset definition
```yaml
---
type: Asset
api_version: core/v2
metadata:
  name: sensu-prometheus-collector_linux_amd64
  labels: 
  annotations:
    io.sensu.bonsai.url: https://bonsai.sensu.io/assets/sensu/sensu-prometheus-collector
    io.sensu.bonsai.api_url: https://bonsai.sensu.io/api/v1/assets/sensu/sensu-prometheus-collector
    io.sensu.bonsai.tier: Supported
    io.sensu.bonsai.version: 1.1.6
    io.sensu.bonsai.namespace: sensu
    io.sensu.bonsai.name: sensu-prometheus-collector
    io.sensu.bonsai.tags: check
spec:
  url: https://assets.bonsai.sensu.io/ef812286f59de36a40e51178024b81c69666e1b7/sensu-prometheus-collector_1.1.6_linux_amd64.tar.gz
  sha512: a70056ca02662fbf2999460f6be93f174c7e09c5a8b12efc7cc42ce1ccb5570ee0f328a2dd8223f506df3b5972f7f521728f7bdd6abf9f6ca2234d690aeb3808
  filters:
  - entity.system.os == 'linux'
  - entity.system.arch == 'amd64'
```

##### Check definition

```yaml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-prom-metrics
  namespace: default
spec:
  check_hooks: null
  command: sensu-prometheus-collector -exporter-url http://localhost:9182/metrics
  output_metric_format: influxdb_line
  output_metric_handlers:
  - influxdb
  runtime_assets:
  - sensu/sensu-prometheus-collector
  stdin: false
  subdue: null
  subscriptions:
  - linux
  timeout: 5
  ttl: 0
```

#### Sensu Core

##### Check definition

```
{
  "checks": {
    "prometheus_metrics": {
      "type": "metric",
      "command": "sensu-prometheus-collector -exporter-url http://localhost:8080/metrics",
      "subscribers": ["app_tier"],
      "interval": 10,
      "handler": "influx"
    }
  }
}
```

##### Handler definition

```
{
  "handlers": {
    "influx": {
      "type": "udp",
      "mutator": "only_check_output",
      "socket": {
        "host": "influx.example.com",
        "port": 8189
      }
    }
  }
}
```


## Installation from source and contributing

The preferred way of installing and deploying this plugin is to use it as an [asset]. If you would like to compile and install the plugin from source or contribute to it, download the latest version of the sensu-CHANGEME from [releases][1]
or create an executable script from this source.

From the local path of the sensu-prometheus-collector repository:

```
go build -o /usr/local/bin/sensu-prometheus-collector main.go
```

If you'd like to contribute to this plugin, see https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/sensu/sensu-prometheus-collector/releases
[2]: #asset-registration
