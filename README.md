[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Prometheus%20Collector-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/sensu/sensu-prometheus-collector)

# Sensu Prometheus Collector

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Resource definition](#resource-definition)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

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

## Usage examples

Help:

```
Usage of sensu-prometheus-collector:
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

## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add sensu/sensu-prometheus-collector
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/sensu/sensu-prometheus-collector).

### Resource definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-prom-metrics
spec:
  command: sensu-prometheus-collector -exporter-url http://localhost:9182/metrics
  output_metric_format: influxdb_line
  output_metric_handlers:
  - influxdb
  runtime_assets:
  - sensu/sensu-prometheus-collector
  subscriptions:
  - linux
  timeout: 5
  ttl: 0
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the sensu-prometheus-collector repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/check-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/check-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu-community/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
