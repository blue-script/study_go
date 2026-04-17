package postgres

import "time"

type BookModel struct {
	ID          int        `db:"id"`
	Title       string     `db:"title"`
	Author      string     `db:"author"`
	Review      *string    `db:"review"`
	ReleaseYear int        `db:"release_year"`
	IsRead      bool       `db:"is_read"`
	AddedAt     time.Time  `db:"added_at"`
	CompletedAt *time.Time `db:"completed_at"`
}

type UserModel struct {
	ID          int    `db:"id"`
	FullName    string `db:"full_name"`
	PhoneNumber *string `db:"phone_number"`
}
