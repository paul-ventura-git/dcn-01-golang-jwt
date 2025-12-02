package models

type Cliente struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}
