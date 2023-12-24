package models

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/fredmessias43/car-hub/src/config"
	"github.com/fredmessias43/car-hub/src/websocket"
	"gorm.io/gorm"
)

type Manufacturer struct {
	gorm.Model
	ID            int    `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	CountryOrigin string `json:"country_of_origin"`
}

func (m *Manufacturer) ToMap() map[string]any {
	return map[string]any{
		"ID":            m.ID,
		"Name":          m.Name,
		"CountryOrigin": m.CountryOrigin,
	}
}

func (m *Manufacturer) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

func (m *Manufacturer) AfterSave(tx *gorm.DB) error {
	room := config.WS.FindRoomByName("manufacturer-created")
	log.Println(strconv.Itoa(m.ID) + " => " + room.Name)

	message := websocket.Message{
		Action:  websocket.SendMessageAction,
		Message: string(m.ToJson()),
		Target:  room,
	}

	room.BroadcastToClientsInRoom(message.Encode())
	return nil
}
