package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID              int             `json:"id" gorm:"primaryKey"`
	Color           string          `json:"color"`
	Mileage         int             `json:"mileage"`
	Price           float64         `json:"price"`
	FuelType        string          `json:"fuel_type"`
	Transmission    string          `json:"transmission"`
	ModelVersionID  int             `json:"model_version_id"`
	CarModelVersion CarModelVersion `gorm:"foreignKey:ModelVersionID" form:"-"`
}

func (m Car) GetID() int {
	return m.ID
}

func (m Car) GetName() string {
	return m.Color
}

func (m Car) GetRelationshipValue() int {
	return m.ModelVersionID
}

func (m Car) ToMap() map[string]any {
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

func (m Car) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

// func (m Car) AfterSave(tx *gorm.DB) error {
// return websocket.Emit("car-created", m)
// }
