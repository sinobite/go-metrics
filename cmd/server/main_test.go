package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestRouter(t *testing.T) {
	ts := httptest.NewServer(NewRouter())
	defer ts.Close()

	var testTable = []struct {
		url    string
		want   string
		status int
		method string
	}{
		{"/update/gauge/testGaugeMetricName/328479.927", "", http.StatusOK, "POST"},
		{"/value/gauge/testGaugeMetricName", "328479.927", http.StatusOK, "GET"},
		{"/update/counter/testMetricName/123", "", http.StatusOK, "POST"},
		{"/update/wrongType/testMetricName/123", "", http.StatusBadRequest, "POST"},
		{"/value/counter/testMetricName", "123", http.StatusOK, "GET"},
		{"/value/counter/wrongTestMetricName", "", http.StatusNotFound, "GET"},
		{"/", "testGaugeMetricName: 328479.927 \ntestMetricName: 123 \n", http.StatusOK, "GET"},
	}
	for _, v := range testTable {
		resp, get := testRequest(t, ts, v.method, v.url)
		assert.Equal(t, v.status, resp.StatusCode)
		assert.Equal(t, v.want, get)
		resp.Body.Close()
	}
}
