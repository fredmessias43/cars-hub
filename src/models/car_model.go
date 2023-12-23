package models

import (
	"encoding/json"

	"github.com/fredmessias43/car-hub/src/config"
	"gorm.io/gorm"
)

type CarModel struct {
	gorm.Model
	ID              int    `json:"id" gorm:"primaryKey"`
	BrandID         int    `json:"brand_id"`
	Name            string `json:"name"`
	ManufactureYear int    `json:"manufacture_year"`
	Brand           Brand  `form:"-"`
}

func (m *CarModel) ToMap() map[string]any {
	return map[string]any{
		"ID":              m.ID,
		"BrandID":         m.BrandID,
		"Name":            m.Name,
		"ManufactureYear": m.ManufactureYear,
	}
}

func (m *CarModel) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

func (m *CarModel) AfterCreate(tx *gorm.DB) error {
	room := config.WS.FindRoomByName("")
	room.BroadcastToClientsInRoom(m.ToJson())
	return nil
}
