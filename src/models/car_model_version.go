package models

import (
	"encoding/json"

	"github.com/fredmessias43/car-hub/src/config"
	"gorm.io/gorm"
)

type CarModelVersion struct {
	gorm.Model
	ID             int      `json:"id" gorm:"primaryKey"`
	Name           string   `json:"name"`
	ModelID        int      `json:"model_id"`
	Specifications string   `json:"specifications"`
	Features       string   `json:"features"`
	CarModel       CarModel `gorm:"foreignKey:ModelID" form:"-"`
}

func (m *CarModelVersion) ToMap() map[string]any {
	return map[string]any{
		"ID":             m.ID,
		"Name":           m.Name,
		"ModelID":        m.ModelID,
		"Specifications": m.Specifications,
		"Features":       m.Features,
	}
}

func (m *CarModelVersion) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

func (m *CarModelVersion) AfterCreate(tx *gorm.DB) error {
	room := config.WS.FindRoomByName("")
	room.BroadcastToClientsInRoom(m.ToJson())
	return nil
}
