package usecase

import (
    "context"
    "fmt"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/domain"
    "github.com/trapilot/gomicro-boilerplate-practice/internal/repository"
)

type AuthServiceImpl struct {
    repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) domain.AuthService {
    return &AuthServiceImpl{repo: repo}
}

func (a *AuthServiceImpl) Login(ctx context.Context, username, password string) (string, error) {
    storedPwd, err := a.repo.GetPassword(username)
    if err != nil {
        return "", fmt.Errorf("invalid credentials: %w", err)
    }
    if storedPwd != password {
        return "", fmt.Errorf("invalid credentials")
    }
    return fmt.Sprintf("token-%s", username), nil
}

// Register creates a new user with the given password.
func (a *AuthServiceImpl) Register(ctx context.Context, username, password string) error {
    if err := a.repo.CreateUser(username, password); err != nil {
        return fmt.Errorf("registration failed: %w", err)
    }
    return nil
}
