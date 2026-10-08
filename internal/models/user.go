package models

import "time"

type User struct {
	ID           uint      `gorm:"column:id;primaryKey"`
	Email        string    `gorm:"column:email;uniqueIndex;not null"`
	PasswordHash string    `gorm:"column:password_hash;not null"`
	FirstName    string    `gorm:"column:first_name"`
	LastName     string    `gorm:"column:last_name"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}
