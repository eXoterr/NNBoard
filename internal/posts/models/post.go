package models

type Post struct {
	ID     int    `db:"id"`
	Title  string `form:"title" validate:"required"`
	Body   string `form:"body" validate:"required"`
	Author uint   `db:"author"`
}
