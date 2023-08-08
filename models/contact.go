package models

import "gorm.io/gorm"

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
