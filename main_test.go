package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestQueryExporter(t *testing.T) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":7777", nil)
		if err != nil {
			fmt.Printf("failed to create a test webserver: %s", err)
			os.Exit(3)
		}
	}()

	time.Sleep(2 * time.Second)

	samples, err := QueryExporter("http://localhost:7777/metrics", ExporterAuth{User: "", Password: "", Header: ""}, false, "", "", "")

	assert.NoError(t, err)
	assert.NotNil(t, samples)
}
