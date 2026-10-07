package dto

import "time"

type User struct {
	ID        int       `json:"-" db:"id"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	FirstName string    `json:"first_name" db:"first_name" binding:"required"`
	LastName  string    `json:"last_name" db:"last_name" binding:"required"`
	CreatedAt time.Time `json:"registered_date" db:"registered_date"`
}

type UserSignInRequest struct {
	Email    string `json:"email" default:"marisasha228@bk.ru"`
	Password string `json:"password" default:"123"`
}
