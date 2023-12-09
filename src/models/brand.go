package models

import "gorm.io/gorm"

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
