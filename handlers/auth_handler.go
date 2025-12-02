package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/auth"
	"microservice/database"
	"microservice/models"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	json.NewDecoder(r.Body).Decode(&input)

	hashed, _ := auth.HashPassword(input.Password)

	user := models.Usuario{
		Email:    input.Email,
		Password: hashed,
	}

	database.DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario creado"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	json.NewDecoder(r.Body).Decode(&input)

	var user models.Usuario
	database.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 || !auth.CheckPassword(user.Password, input.Password) {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	token, _ := auth.GenerateToken(user.ID)

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
