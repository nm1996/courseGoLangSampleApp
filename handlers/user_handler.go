package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"simpleGoLang/model"
	"time"
)

type LoginDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ResponseDto struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Token   string `json:"token"`
}

var jwtSecretKey = []byte("your_secret_key")

func generateJwt(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserByNameAndPasswordHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var loginDto LoginDto
	if err := json.NewDecoder(r.Body).Decode(&loginDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := model.GetUserByNameAndPassword(loginDto.Name, loginDto.Password, db)
	if err != nil {
		http.Error(w, "Credentials are missed", http.StatusInternalServerError)
		return
	}

	jwt, err := generateJwt(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	responseDto := ResponseDto{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Surname: user.Surname,
		Token:   jwt,
	}

	//map user to responseDto

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDto)
}

func InsertUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := model.InsertUser(user, db)
	if err != nil {
		http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
