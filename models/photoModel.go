package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    string `json:"user_id"` // Gunakan tipe data yang sesuai, misalnya uint
	User      User   `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
