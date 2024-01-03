package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID             int          `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name"`
	ManufacturerID int          `json:"manufacturer_id"`
	Manufacturer   Manufacturer `form:"-"`
}

func (m Brand) GetID() int {
	return m.ID
}

func (m Brand) GetName() string {
	return m.Name
}

func (m Brand) GetRelationshipValue() int {
	return m.ManufacturerID
}

func (m Brand) ToMap() map[string]any {
	return map[string]any{
		"ID":             m.ID,
		"Name":           m.Name,
		"ManufacturerID": m.ManufacturerID,
	}
}

func (m Brand) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

// func (m Brand) AfterSave(tx *gorm.DB) error {
// return websocket.Emit("brand-created", m)
// }
