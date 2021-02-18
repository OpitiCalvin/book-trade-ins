package datamodels

import "github.com/jinzhu/gorm"

// User type to capture schema and fields for user model
type User struct {
	gorm.Model
	Email    string `json:"email"`
	UserName string `json:"username"`
	password []byte
}
