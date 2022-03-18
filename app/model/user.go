package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model  `json:"model"`
	PhoneNumber string
	Password    string
	CreatedBy   int
	UpdatedBy   int
}