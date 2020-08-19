package models

// User entity
type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

// UserRepository ...
type UserRepository interface {
	FindByEmailAndPassword(email string, password string) (user *User, err error)
}

// IsEmailValid is not implemented yet
func IsEmailValid(email string) bool {
	return true
}
