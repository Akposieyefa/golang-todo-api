package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
	UserID uint   `json:"userId"`
}
