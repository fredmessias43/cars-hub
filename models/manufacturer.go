package models

import "gorm.io/gorm"

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
