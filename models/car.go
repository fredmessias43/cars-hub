package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID             int     `json:"id" gorm:"primaryKey"`
	ModelVersionID int     `json:"model_version_id"`
	Color          string  `json:"color"`
	Mileage        int     `json:"mileage"`
	Price          float64 `json:"price"`
	FuelType       string  `json:"fuel_type"`
	Transmission   string  `json:"transmission"`
}

func (m *Car) ToMap() map[string]any {
	return map[string]any{
		"ID":             m.ID,
		"ModelVersionID": m.ModelVersionID,
		"Color":          m.Color,
		"Mileage":        m.Mileage,
		"Price":          m.Price,
		"FuelType":       m.FuelType,
		"Transmission":   m.Transmission,
	}
}
