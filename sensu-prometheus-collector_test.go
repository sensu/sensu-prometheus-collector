package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestQueryExporter(t *testing.T) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":7777", nil)
	}()

	time.Sleep(2 * time.Second)

	samples, err := QueryExporter("http://localhost:7777/metrics", exporterAuth{User: "", Password: ""})

	assert.NoError(t, err)
	assert.NotNil(t, samples)
}
