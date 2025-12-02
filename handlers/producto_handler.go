package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/database"
	"microservice/models"
)

// CrearProducto godoc
// @Summary Crear un producto
// @Description Crea un nuevo producto en la base de datos
// @Tags Productos
// @Accept json
// @Produce json
// @Param producto body models.Producto true "Datos del producto"
// @Success 200 {object} models.Producto
// @Router /productos [post]
func CrearProducto(w http.ResponseWriter, r *http.Request) {
	var p models.Producto
	json.NewDecoder(r.Body).Decode(&p)

	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

// ListarProductos godoc
// @Summary Listar productos
// @Description Retorna todos los productos guardados
// @Tags Productos
// @Produce json
// @Success 200 {array} models.Producto
// @Router /productos [get]
func ListarProductos(w http.ResponseWriter, r *http.Request) {
	var productos []models.Producto
	database.DB.Find(&productos)
	json.NewEncoder(w).Encode(productos)
}
