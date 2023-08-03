package models

import "gorm.io/gorm"

type CarModelVersion struct {
	gorm.Model
	ID             int    `json:"ID" gorm:"primaryKey"`
	Name           string `json:"Name"`
	ModelID        int    `json:"Model_ID"`
	Specifications string `json:"Specifications"`
	Features       string `json:"Features"`
}
