package models

import (
	"html"
	"strings"
	"time"

	"github.com/ReCodink/task-5-pbi-btpns-Muhammad-Raihan/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"not null" json:"username" binding:"required"`
	Email     string  `gorm:"unique; not null" json:"email" binding:"required"` // Pastikan validasi format email
	Password  string  `gorm:"not null" json:"password" binding:"required"`
	Photos    []Photo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photos"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		hashedPassword, err := helpers.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return
}

func (user *User) ComparePassword(password string) error {
	return helpers.CheckedPasswordHash(password, user.Password)
}
