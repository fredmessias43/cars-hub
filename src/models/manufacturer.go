package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Manufacturer struct {
	gorm.Model
	ID            int    `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	CountryOrigin string `json:"country_of_origin"`
}

func (m Manufacturer) GetID() int {
	return m.ID
}

func (m Manufacturer) GetName() string {
	return m.Name
}

func (m Manufacturer) GetRelationshipValue() int {
	return 0
}

func (m Manufacturer) ToMap() map[string]any {
	return map[string]any{
		"ID":            m.ID,
		"Name":          m.Name,
		"CountryOrigin": m.CountryOrigin,
	}
}

func (m Manufacturer) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

// func (m Manufacturer) AfterSave(tx *gorm.DB) error {
// return websocket.Emit("manufacturer-created", m)
// }
