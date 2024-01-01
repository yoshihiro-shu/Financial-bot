package kafka

import (
	"io"
	"net/http"
	"testing"

	"github.com/IBM/sarama"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestHandleRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	type args struct {
		msg    *sarama.ConsumerMessage
		method string
		path   string
		res    httpmock.Responder
	}
	type expected struct {
		statusCode int
		body       string
	}
	tests := []struct {
		key   string
		input args
		expected
	}{
		{
			key: "GET",
			input: args{
				msg: &sarama.ConsumerMessage{
					Topic:     "request",
					Partition: 0,
					Offset:    0,
					Key:       []byte(http.MethodGet),
					Value:     []byte("http://example.com/v1/users"),
				},
				method: http.MethodGet,
				path:   "http://example.com/v1/users",
				res:    httpmock.NewStringResponder(200, http.StatusText(http.StatusOK)),
			},
			expected: expected{
				statusCode: http.StatusOK,
				body:       http.StatusText(http.StatusOK),
			},
		},
	}
	for _, tt := range tests {
		httpmock.RegisterResponder(tt.key, tt.input.path, tt.input.res)
		res, err := handleRequest(tt.input.msg)
		assert.Nil(t, err)
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, tt.expected.statusCode, res.StatusCode)
		assert.Equal(t, tt.expected.body, string(body))
	}
}
