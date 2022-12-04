package models

type User struct {
	ID        int    `db:"id"`
	Login     string `form:"login" validate:"required"`
	Password  string `form:"password" validate:"required"`
	Password2 string `form:"password2"`
	Role      uint   `db:"role"`
}
