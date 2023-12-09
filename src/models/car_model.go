package models

import "gorm.io/gorm"

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
