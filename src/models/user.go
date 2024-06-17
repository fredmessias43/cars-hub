package models

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/fredmessias43/car-hub/src/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

func (m User) GetID() int {
	return m.ID
}

func (m User) GetName() string {
	return m.FirstName
}

func (m User) GetRelationshipValue() int {
	return 0
}

func (m User) ToMap() map[string]any {
	return map[string]any{
		"ID":        m.ID,
		"FirstName": m.FirstName,
		"LastName":  m.LastName,
		"Email":     m.Email,
	}
}

func (m User) ToJson() []byte {
	bytes, _ := json.Marshal(m.ToMap())
	return bytes
}

// func (m User) AfterSave(tx *gorm.DB) error {
// return websocket.Emit("user-created", m)
// }

func GetUser(email, password string) *User {
	var user User
	result := config.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return nil
	}
	return &user
}

func CreateUser(email, password string) (*User, error) {
	user := User{Email: email, Password: password}
	result := config.DB.Create(&user)
	return &user, result.Error
}

func GetAuthUser() *User {
	session := config.GetSession()
	if session.Values["authenticated"] != true {
		return nil
	}
	value := session.Values["user"].(string)
	userID, _ := strconv.Atoi(strings.Split(value, "|")[0])
	user := User{ID: userID}
	config.DB.Find(&user)

	return &user
}
