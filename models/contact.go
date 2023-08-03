package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email     string
}

func (c *Contact) ToMap() map[string]any {
	return map[string]any{
		"ID":        c.ID,
		"FirstName": c.FirstName,
		"LastName":  c.LastName,
		"Email":     c.Email,
	}
}

func GetContactById(contacts []Contact, ID int) Contact {
	var result Contact
	for _, v := range contacts {
		if v.ID == ID {
			result = v
		}
	}
	return result
}
