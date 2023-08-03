package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID             int     `json:"ID" gorm:"primaryKey"`
	ModelVersionID int     `json:"Model_Version_ID"`
	Color          string  `json:"Color"`
	Mileage        int     `json:"Mileage"`
	Price          float64 `json:"Price"`
	FuelType       string  `json:"Fuel_Type"`
	Transmission   string  `json:"Transmission"`
}
