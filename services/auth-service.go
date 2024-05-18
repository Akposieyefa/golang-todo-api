package services

import (
	"akposieyefa/golang-todo-api/models"
	"akposieyefa/golang-todo-api/pkg"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// find user details for login
func FindOne(email, password string) map[string]interface{} {
	user := &models.User{}

	if err := pkg.DB.Where("Email = ?", email).First(user).Error; err != nil {
		return map[string]interface{}{
			"status":  false,
			"message": "Email address not found",
		}
	}
	expirationTime := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return map[string]interface{}{
			"status":  false,
			"message": "Invalid login credentials. Please try again",
		}
	}

	tk := &models.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(os.Getenv("JWT_ALGO")), tk)

	tokenString, error := token.SignedString(jwtKey)
	if error != nil {
		return map[string]interface{}{
			"status":  false,
			"message": error,
		}
	}

	var resp = map[string]interface{}{"status": true, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp
}

// get authenticated user
func AuthUser(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	tokenString := r.Header.Get("Authorization")

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return map[string]interface{}{
				"status":  false,
				"message": "Signature validation error",
			}
		}
		w.WriteHeader(http.StatusBadRequest)
		return map[string]interface{}{
			"status":  false,
			"message": "Signature validation error",
		}
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return map[string]interface{}{
			"status":  false,
			"message": "Invalid token",
		}
	}
	loggedIn, err := getUserByEmail(claims.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return map[string]interface{}{
			"status":  false,
			"message": "Error retrieving user data",
		}
	}
	w.WriteHeader(http.StatusOK)
	return map[string]interface{}{
		"message": "Logged in email pulled successfully",
		"user":    loggedIn,
		"success": true,
	}
}

// get user by email
func getUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := pkg.DB.Where("Email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
