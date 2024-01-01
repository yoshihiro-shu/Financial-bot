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
	httpmock.RegisterResponder("GET", "http://example.com/v1/users",
		httpmock.NewStringResponder(200, "mocked"),
	)
	msg := &sarama.ConsumerMessage{
		Topic:     "request",
		Partition: 0,
		Offset:    0,
		Key:       []byte(http.MethodGet),
		Value:     []byte("http://example.com/v1/users"),
	}
	res, err := handleRequest(msg)
	assert.Nil(t, err)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "mocked", string(body))
}
