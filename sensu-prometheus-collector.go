package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
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

func CreateMetrics(samples model.Vector) ([]Metric, error) {
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

	return metrics, nil
}

func CreateGraphiteMetrics(samples model.Vector) (string, error) {
	metrics := ""

	for _, sample := range samples {
		name := sample.Metric["__name__"]

		value := sample.Value

		now := time.Now()
		timestamp := now.Unix()

		metric := fmt.Sprintf("%s %f %v\n", name, value, timestamp)

		metrics += metric
	}

	return metrics, nil
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
	flag.Parse()

	samples, err := QueryPrometheus(*promURL, *queryString)

	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	metrics, _ := CreateMetrics(samples)
	for _, metric := range metrics {
		fmt.Printf("%+v\n", metric)
	}

	graphiteMetrics, _ := CreateGraphiteMetrics(samples)
	fmt.Printf("%s\n", graphiteMetrics)
}
