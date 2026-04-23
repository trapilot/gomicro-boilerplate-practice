package handler

import (
    "context"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHealthHandler_Ping(t *testing.T) {
    h := NewHealthHandler()
    resp, err := h.Ping(context.Background(), &PingRequest{})
    assert.NoError(t, err)
    assert.Equal(t, "pong", resp.Message)
}
