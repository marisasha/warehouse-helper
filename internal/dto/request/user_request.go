package request

type User struct {
	Email     string `json:"email"  binding:"required"`
	Password  string `json:"password" gorm:"column:password_hash" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type UserSignIn struct {
	Email    string `json:"email" default:"marisasha228@bk.ru"`
	Password string `json:"password" default:"123"`
}
