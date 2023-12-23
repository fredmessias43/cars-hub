package models

import (
	"encoding/json"

	"github.com/fredmessias43/car-hub/src/config"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (m *Contact) ToMap() map[string]any {
	return map[string]any{
		"ID":        m.ID,
		"FirstName": m.FirstName,
		"LastName":  m.LastName,
		"Email":     m.Email,
	}
}

func (m *Contact) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

func (m *Contact) AfterCreate(tx *gorm.DB) error {
	room := config.WS.FindRoomByName("")
	room.BroadcastToClientsInRoom(m.ToJson())
	return nil
}
