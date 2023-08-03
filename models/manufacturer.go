package models

import "gorm.io/gorm"

type Manufacturer struct {
	gorm.Model
	ID            int    `json:"ID" gorm:"primaryKey"`
	Name          string `json:"Name"`
	CountryOrigin string `json:"Country_of_Origin"`
}
