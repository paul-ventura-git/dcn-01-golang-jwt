package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/database"
	"microservice/models"
)

func CrearProducto(w http.ResponseWriter, r *http.Request) {
	var p models.Producto
	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func ListarProductos(w http.ResponseWriter, r *http.Request) {
	var productos []models.Producto
	database.DB.Find(&productos)
	json.NewEncoder(w).Encode(productos)
}
