package token

import "fmt"

// Token library contract
type Token interface {
	Create(claims map[string]string) (token string)
}

// MockedToken ...
type MockedToken struct {
	ExpectedToken string
}

// Create a mocked token
func (mt *MockedToken) Create(claims map[string]string) string {
	fmt.Println(claims)
	return mt.ExpectedToken
}
