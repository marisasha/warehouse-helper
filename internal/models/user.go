package models

import "time"

type User struct {
	ID           uint      `gorm:"column:id;primaryKey" json:"id"`
	Email        string    `gorm:"column:password_hash;uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	FirstName    string    `gorm:"column:first_name" json:"first_name"`
	LastName     string    `gorm:"column:last_name" json:"last_name"`
	CreatedAt    time.Time `gorm:"column:last_name" json:"created_at"`
}
