package models

type Factura struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	ClienteID uint       `json:"cliente_id"`
	Cliente   Cliente    `json:"cliente" gorm:"foreignKey:ClienteID"`
	Productos []Producto `json:"productos" gorm:"many2many:factura_productos"`
	Total     float64    `json:"total"`
}
