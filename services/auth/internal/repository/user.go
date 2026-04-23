package repository

import "fmt"

type User struct {
    Username string
    Password string
}

type UserRepository interface {
    // GetPassword returns the password for a given username, or error if not found
    GetPassword(username string) (string, error)
    // CreateUser adds a new user. Returns error if username already exists.
    CreateUser(username, password string) error
}

type InMemoryUserRepo struct {
    users map[string]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    // Prepopulate with a sample user
    defaultUsers := map[string]User{
        "alice": {Username: "alice", Password: "secret"},
    }
    return &InMemoryUserRepo{users: defaultUsers}
}

func (r *InMemoryUserRepo) GetPassword(username string) (string, error) {
    if user, ok := r.users[username]; ok {
        return user.Password, nil
    }
    return "", fmt.Errorf("user %s not found", username)
}

// CreateUser adds a new user to the repo. Returns error if username exists.
func (r *InMemoryUserRepo) CreateUser(username, password string) error {
    if _, exists := r.users[username]; exists {
        return fmt.Errorf("user %s already exists", username)
    }
    r.users[username] = User{Username: username, Password: password}
    return nil
}
