package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_updateMetricHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        200,
				contentType: "plain/text",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/update/counter/someMetric/527", nil)
			request.SetPathValue("metricType", "counter")
			request.SetPathValue("metricName", "someMetric")
			request.SetPathValue("metricValue", "527")

			// создаём новый Recorder
			w := httptest.NewRecorder()
			updateMetricHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()

			//assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
