package models

import "gorm.io/gorm"

type CarModelVersion struct {
	gorm.Model
	ID             int      `json:"id" gorm:"primaryKey"`
	Name           string   `json:"name"`
	ModelID        int      `json:"model_id"`
	Specifications string   `json:"specifications"`
	Features       string   `json:"features"`
	CarModel       CarModel `gorm:"foreignKey:ModelID"`
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
