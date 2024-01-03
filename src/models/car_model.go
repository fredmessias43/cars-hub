package models

import (
	"encoding/json"

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

func (m CarModel) GetID() int {
	return m.ID
}

func (m CarModel) GetName() string {
	return m.Name
}

func (m CarModel) GetRelationshipValue() int {
	return m.BrandID
}

func (m CarModel) ToMap() map[string]any {
	return map[string]any{
		"ID":              m.ID,
		"BrandID":         m.BrandID,
		"Name":            m.Name,
		"ManufactureYear": m.ManufactureYear,
	}
}

func (m CarModel) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

// func (m CarModel) AfterSave(tx *gorm.DB) error {
// return websocket.Emit("carModel-created", m)
// }
