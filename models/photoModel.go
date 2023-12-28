package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"photo_id" binding:"required"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserID    uint   `json:"user_id"` // Gunakan tipe data yang sesuai, misalnya uint
	User      User   `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
