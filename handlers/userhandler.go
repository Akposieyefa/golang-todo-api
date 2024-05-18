package handlers

import (
	"akposieyefa/golang-todo-api/helpers"
	"akposieyefa/golang-todo-api/models"
	"akposieyefa/golang-todo-api/pkg"
	"akposieyefa/golang-todo-api/services"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// login user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Invalid request",
			"body":    r.Body,
			"success": false,
		})
		return
	}
	resp := services.FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

// create user account
func CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	err := helpers.Validate(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Password Encryption  failed",
			"success": false,
		})
		return
	}

	user.Password = string(pass)

	createdUser := pkg.DB.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": errMessage,
			"success": false,
		})
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":   user,
		"message": "Success in creating users",
		"success": true,
	})
}

// get logged in users profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	resp := services.AuthUser(w, r)
	json.NewEncoder(w).Encode(resp)
}
