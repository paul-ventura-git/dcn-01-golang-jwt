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

// Register godoc
// @Summary Registro de usuario
// @Description Crea un nuevo usuario con email y password
// @Accept  json
// @Produce  json
// @Param user body RegisterInput true "Datos de registro"
// @Success 201 {object} map[string]interface{}
// @Router /register [post]
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

// Login godoc
// @Summary Login
// @Description Retorna un JWT válido
// @Accept  json
// @Produce  json
// @Param user body LoginInput true "Credenciales"
// @Success 200 {object} map[string]string
// @Router /login [post]
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
