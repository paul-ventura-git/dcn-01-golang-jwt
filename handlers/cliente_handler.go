package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/database"
	"microservice/models"
)

// CrearCliente godoc
// @Summary Crea un cliente
// @Description Crea un nuevo cliente (requiere JWT)
// @Tags Clientes
// @Security Bearer
// @Accept  json
// @Produce  json
// @Param cliente body models.Cliente true "Datos del cliente"
// @Success 200 {object} models.Cliente
// @Router /clientes [post]
func CrearCliente(w http.ResponseWriter, r *http.Request) {
	var c models.Cliente
	json.NewDecoder(r.Body).Decode(&c)

	database.DB.Create(&c)
	json.NewEncoder(w).Encode(c)
}

// ListarClientes godoc
// @Summary Lista clientes
// @Description Obtiene todos los clientes (requiere JWT)
// @Tags Clientes
// @Security Bearer
// @Produce  json
// @Success 200 {array} models.Cliente
// @Router /clientes [get]
func ListarClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []models.Cliente
	database.DB.Find(&clientes)
	json.NewEncoder(w).Encode(clientes)
}
