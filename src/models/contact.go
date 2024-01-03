package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (m Contact) GetID() int {
	return m.ID
}

func (m Contact) GetName() string {
	return m.FirstName
}

func (m Contact) GetRelationshipValue() int {
	return 0
}

func (m Contact) ToMap() map[string]any {
	return map[string]any{
		"ID":        m.ID,
		"FirstName": m.FirstName,
		"LastName":  m.LastName,
		"Email":     m.Email,
	}
}

func (m Contact) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

// func (m Contact) AfterSave(tx *gorm.DB) error {
// return websocket.Emit("contact-created", m)
// }
