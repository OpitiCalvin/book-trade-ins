package datamodels

import (
	"github.com/jinzhu/gorm"
)

// Books type capturing schema and fields relevant to books model
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}
