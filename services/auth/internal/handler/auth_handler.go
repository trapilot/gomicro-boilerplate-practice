package handler

import (
    "context"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/domain"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/repository"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/usecase"
)

// AuthHandler implements the AuthService RPC methods for go-micro.
// It internally uses the domain.AuthService implementation.
type AuthHandler struct {
    service domain.AuthService
}

// NewAuthHandler creates a new AuthHandler with the provided repository.
func NewAuthHandler(repo repository.UserRepository) *AuthHandler {
    svc := usecase.NewAuthService(repo)
    return &AuthHandler{service: svc}
}

// LoginRequest represents the RPC request payload.
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginResponse represents the RPC response payload.
type LoginResponse struct {
    Token string `json:"token,omitempty"`
    Err   string `json:"error,omitempty"`
}

// Login handles authentication requests.
func (h *AuthHandler) Login(ctx context.Context, req *LoginRequest, rsp *LoginResponse) error {
    token, err := h.service.Login(ctx, req.Username, req.Password)
    if err != nil {
        rsp.Err = err.Error()
        return nil // return nil to indicate RPC handled; error conveyed in response
    }
    rsp.Token = token
    return nil
}
