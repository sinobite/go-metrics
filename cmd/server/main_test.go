package main

import (
	"github.com/sinobite/go-metrics/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (int, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)
}

func TestRouter(t *testing.T) {
	s := storage.New()
	ts := httptest.NewServer(NewRouter(s))
	defer ts.Close()

	var testTable = []struct {
		url    string
		want   string
		status int
		method string
		name   string
	}{
		{"/update/gauge/testGaugeMetricName/328479.927", "", http.StatusOK, "POST", "Update gauge metric"},
		{"/value/gauge/testGaugeMetricName", "328479.927", http.StatusOK, "GET", "Find gauge metric"},
		{"/update/counter/testMetricName/123", "", http.StatusOK, "POST", "Update counter metric"},
		{"/update/wrongType/testMetricName/123", "Metric type is incorrect\n", http.StatusBadRequest, "POST", "Update metric with wrong type"},
		{"/value/counter/testMetricName", "123", http.StatusOK, "GET", "Find counter metric"},
		{"/value/counter/wrongTestMetricName", "Failed to find counter metrics\n", http.StatusNotFound, "GET", "Find missing metric"},
		{"/", "testGaugeMetricName: 328479.927 \ntestMetricName: 123 \n", http.StatusOK, "GET", "Find all metrics"},
	}
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			statusCode, body := testRequest(t, ts, v.method, v.url)
			assert.Equal(t, v.status, statusCode, "Status code mismatch")
			assert.Equal(t, v.want, body, "Response body mismatch")
		})
	}
}
