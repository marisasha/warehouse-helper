package dto

import "time"

type User struct {
	ID        int       `json:"-"`
	Email     string    `json:"email" gorm:"column:email" binding:"required"`
	Password  string    `json:"password" gorm:"column:password_hash" binding:"required"`
	FirstName string    `json:"first_name" gorm:"column:first_name" binding:"required"`
	LastName  string    `json:"last_name" gorm:"column:last_name" binding:"required"`
	CreatedAt time.Time `json:"-"`
}

type UserSignInRequest struct {
	Email    string `json:"email" default:"marisasha228@bk.ru"`
	Password string `json:"password" default:"123"`
}
