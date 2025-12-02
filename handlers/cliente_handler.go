package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/database"
	"microservice/models"
)

func CrearCliente(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	json.NewDecoder(r.Body).Decode(&c)

	database.DB.Create(&c)
	json.NewEncoder(w).Encode(c)
}

func ListarClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []models.Cliente
	database.DB.Find(&clientes)
	json.NewEncoder(w).Encode(clientes)
}
