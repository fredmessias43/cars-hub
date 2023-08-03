package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	ID   int    `json:"ID" gorm:"primaryKey"`
	Name string `json:"Name"`
}
