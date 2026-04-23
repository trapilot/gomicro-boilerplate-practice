package usecase

import (
    "context"
    "testing"

    "github.com/trapilot/gomicro-boilerplate-practice/internal/repository"
    "github.com/stretchr/testify/assert"
)

func TestAuthService_Login_HappyPath(t *testing.T) {
    repo := repository.NewInMemoryUserRepo()
    service := NewAuthService(repo)
    token, err := service.Login(context.Background(), "alice", "secret")
    assert.NoError(t, err)
    assert.Equal(t, "token-alice", token)
}

func TestAuthService_Login_InvalidUser(t *testing.T) {
    repo := repository.NewInMemoryUserRepo()
    service := NewAuthService(repo)
    _, err := service.Login(context.Background(), "bob", "whatever")
    assert.Error(t, err)
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
    repo := repository.NewInMemoryUserRepo()
    service := NewAuthService(repo)
    _, err := service.Login(context.Background(), "alice", "wrong")
    assert.Error(t, err)
}
