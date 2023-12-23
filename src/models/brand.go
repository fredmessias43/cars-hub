package models

import (
	"encoding/json"

	"github.com/fredmessias43/car-hub/src/config"
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID             int          `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name"`
	ManufacturerID int          `json:"manufacturer_id"`
	Manufacturer   Manufacturer `form:"-"`
}

func (m *Brand) ToMap() map[string]any {
	return map[string]any{
		"ID":             m.ID,
		"Name":           m.Name,
		"ManufacturerID": m.ManufacturerID,
	}
}

func (m *Brand) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

func (m *Brand) AfterCreate(tx *gorm.DB) error {
	room := config.WS.FindRoomByName("")
	room.BroadcastToClientsInRoom(m.ToJson())
	return nil
}
