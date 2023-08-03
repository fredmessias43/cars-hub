package models

import "gorm.io/gorm"

type CarModel struct {
	gorm.Model
	ID              int    `json:"ID" gorm:"primaryKey"`
	BrandID         int    `json:"Brand_ID"`
	Name            string `json:"Name"`
	ManufactureYear int    `json:"Manufacture_Year"`
}
