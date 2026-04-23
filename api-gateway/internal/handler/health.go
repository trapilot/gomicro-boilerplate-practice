package handler

import "context"

type PingRequest struct{}

type PingResponse struct {
    Message string `json:"message"`
}

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
    return &PingResponse{Message: "pong"}, nil
}
