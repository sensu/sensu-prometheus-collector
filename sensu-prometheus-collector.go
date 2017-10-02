package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/api/prometheus"
	"github.com/prometheus/common/model"
)

type Tag struct {
	Name  string
	Value string
}

type Metric struct {
	Tags  []Tag
	Value float64
}

func CreateJSONMetrics(samples model.Vector) string {
	metrics := []Metric{}

	for _, sample := range samples {
		metric := Metric{}

		for name, value := range sample.Metric {
			tag := Tag{
				Name:  string(name),
				Value: string(value),
			}

			metric.Tags = append(metric.Tags, tag)
		}

		metric.Value = float64(sample.Value)

		metrics = append(metrics, metric)
	}

	jsonMetrics, _ := json.Marshal(metrics)

	return string(jsonMetrics)
}

func CreateGraphiteMetrics(samples model.Vector) string {
	metrics := ""

	for _, sample := range samples {
		name := sample.Metric["__name__"]

		value := sample.Value

		now := time.Now()
		timestamp := now.Unix()

		metric := fmt.Sprintf("%s %f %v\n", name, value, timestamp)

		metrics += metric
	}

	return metrics
}

func CreateInfluxMetrics(samples model.Vector) string {
	metrics := ""

	for _, sample := range samples {
		metric := string(sample.Metric["__name__"])

		for name, value := range sample.Metric {
			if name != "__name__" {
				metric += fmt.Sprintf(",%s=%s", name, value)
			}
		}

		value := sample.Value

		now := time.Now()
		timestamp := now.Unix()

		metric += fmt.Sprintf(" value=%f %v\n", value, timestamp)

		metrics += metric
	}

	return metrics
}

func OutputMetrics(samples model.Vector, outputFormat string) error {
	output := ""

	switch outputFormat {
	case "influx":
		output = CreateInfluxMetrics(samples)
	case "graphite":
		output = CreateGraphiteMetrics(samples)
	case "json":
		output = CreateJSONMetrics(samples)
	}

	fmt.Println(output)

	return nil
}

func QueryPrometheus(promURL string, queryString string) (model.Vector, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	promConfig := prometheus.Config{Address: promURL}
	promClient, err := prometheus.New(promConfig)

	if err != nil {
		fmt.Errorf("%v", err)
		return nil, err
	}

	promQueryClient := prometheus.NewQueryAPI(promClient)

	promResponse, err := promQueryClient.Query(ctx, queryString, time.Now())

	if err != nil {
		fmt.Errorf("%v", err)
		return nil, err
	}

	if promResponse.Type() == model.ValVector {
		return promResponse.(model.Vector), nil
	}

	return nil, errors.New("unexpected response type")
}

func main() {
	promURL := flag.String("url", "http://localhost:9090", "Prometheus API URL")
	queryString := flag.String("query", "up", "Prometheus API query string")
	outputFormat := flag.String("output-format", "influx", "The check output format to use for metrics {influx|graphite|json}")
	flag.Parse()

	samples, err := QueryPrometheus(*promURL, *queryString)

	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(2)
	}

	err = OutputMetrics(samples, *outputFormat)

	if err != nil {
		fmt.Errorf("%v", err)
		os.Exit(2)
	}
}
