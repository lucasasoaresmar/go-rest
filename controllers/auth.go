package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/auth/adapters/token"
	"example.com/auth/models"
)

// Auth http handlers
type Auth struct {
	UserRepository models.UserRepository
	Token          token.Token
}

// Login a user
func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		Email    string
		Password string
	}

	var body Body

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := a.UserRepository.FindByEmailAndPassword(body.Email, body.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	marshaledRoles, err := json.Marshal(user.Roles)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	token := a.Token.Create(map[string]string{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"roles": string(marshaledRoles),
	})

	json.NewEncoder(w).Encode(&struct {
		Token string `json:"token"`
	}{token})
}
