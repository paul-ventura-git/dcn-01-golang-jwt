package handlers

import (
	"encoding/json"
	"net/http"

	"microservice/database"
	"microservice/models"
)

type FacturaInput struct {
	ClienteID   uint   `json:"cliente_id"`
	ProductoIDs []uint `json:"producto_ids"`
}

// CrearFactura godoc
// @Summary Crear una factura
// @Description Crea una factura a partir de un cliente y productos (requiere JWT)
// @Tags Facturas
// @Security Bearer
// @Accept json
// @Produce json
// @Param factura body FacturaInput true "Datos para crear factura"
// @Success 200 {object} models.Factura
// @Failure 400 {object} map[string]string
// @Router /facturas [post]
func CrearFactura(w http.ResponseWriter, r *http.Request) {
	var input FacturaInput
	json.NewDecoder(r.Body).Decode(&input)

	var cliente models.Cliente
	if err := database.DB.First(&cliente, input.ClienteID).Error; err != nil {
		http.Error(w, "Cliente no encontrado", http.StatusBadRequest)
		return
	}

	var productos []models.Producto
	database.DB.Find(&productos, input.ProductoIDs)

	var total float64
	for _, p := range productos {
		total += p.Precio
	}

	factura := models.Factura{
		ClienteID: input.ClienteID,
		Productos: productos,
		Total:     total,
	}

	database.DB.Create(&factura)
	json.NewEncoder(w).Encode(factura)
}

// ListarFacturas godoc
// @Summary Listar facturas
// @Description Lista todas las facturas y sus relaciones (requiere JWT)
// @Tags Facturas
// @Security Bearer
// @Produce json
// @Success 200 {array} models.Factura
// @Router /facturas [get]
func ListarFacturas(w http.ResponseWriter, r *http.Request) {
	var facturas []models.Factura
	database.DB.Preload("Cliente").Preload("Productos").Find(&facturas)
	json.NewEncoder(w).Encode(facturas)
}
