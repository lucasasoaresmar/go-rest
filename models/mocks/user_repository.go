package mocks

import (
	"github.com/lucasasoaresmar/go-rest/models"
)

// MockedUserRepository expects an error and an user to be returned
type MockedUserRepository struct {
	ExpectedError error
	ExpectedUser  *models.User
}

// FindByEmailAndPassword ...
func (mur *MockedUserRepository) FindByEmailAndPassword(email string, password string) (user *models.User, err error) {
	if mur.ExpectedError != nil {
		return nil, err
	}

	return mur.ExpectedUser, nil
}
