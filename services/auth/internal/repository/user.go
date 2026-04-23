package repository

import "fmt"

type User struct {
    Username string
    Password string
}

type UserRepository interface {
    // GetPassword returns the password for a given username, or error if not found
    GetPassword(username string) (string, error)
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
