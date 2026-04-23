package domain

import "context"

// AuthService defines authentication operations.
type AuthService interface {
    // Login validates credentials and returns an auth token.
    // Returns an error if the credentials are invalid.
    Login(ctx context.Context, username, password string) (string, error)
}
