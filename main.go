package main

import (
	"log"
	"net/http"

	"microservice/database"
	"microservice/handlers"
	"microservice/middleware"
	"microservice/models"
)

func main() {

	database.Connect()

	database.DB.AutoMigrate(
		&models.Usuario{},
		&models.Cliente{},
		&models.Producto{},
		&models.Factura{},
	)

	// Auth
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	// Rutas protegidas
	http.HandleFunc("/clientes", middleware.Protected(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.CrearCliente(w, r)
		} else {
			handlers.ListarClientes(w, r)
		}
	}))

	http.HandleFunc("/productos", middleware.Protected(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.CrearProducto(w, r)
		} else {
			handlers.ListarProductos(w, r)
		}
	}))

	http.HandleFunc("/facturas", middleware.Protected(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handlers.CrearFactura(w, r)
		} else {
			handlers.ListarFacturas(w, r)
		}
	}))

	log.Println("🚀 Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
