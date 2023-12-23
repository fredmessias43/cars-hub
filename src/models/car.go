package models

import (
	"encoding/json"

	"github.com/fredmessias43/car-hub/src/config"
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

func (m *Car) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

func (m *Car) AfterCreate(tx *gorm.DB) error {
	room := config.WS.FindRoomByName("")
	room.BroadcastToClientsInRoom(m.ToJson())
	return nil
}
