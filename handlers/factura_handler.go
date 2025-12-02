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

func CrearFactura(w http.ResponseWriter, r *http.Request) {
	var input FacturaInput
	json.NewDecoder(r.Body).Decode(&input)

	var cliente models.Cliente
	database.DB.First(&cliente, input.ClienteID)

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

func ListarFacturas(w http.ResponseWriter, r *http.Request) {
	var facturas []models.Factura
	database.DB.Preload("Cliente").Preload("Productos").Find(&facturas)
	json.NewEncoder(w).Encode(facturas)
}
