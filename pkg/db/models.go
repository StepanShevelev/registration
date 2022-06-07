package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           int `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `json:"name" db:"name"`
	Email        string         `json:"email" db:"email"`
	Password     string         `json:"-" db:"password"`
	PasswordHash string         `json:"-" db:"password_hash"`
	JwtToken     string         `json:"-" db:"token"`

	Posts []Post `json:"posts" db:"posts" gorm:"many2many:user_posts;"`
}

type Post struct {
	ID          int `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Users       []User         `json:"users" db:"users" gorm:"many2many:user_posts;"`
	Title       string         `json:"title" db:"title"`
	Description string         `json:"description" db:"description"`
	Image       string         `json:"image" db:"image"`
}

// ErrLogs storage some error logs
type ErrLogs struct {
	gorm.Model
	Error string `json:"error" db:"error"`
	Place string `json:"place" db:"place"`
	Count int    `json:"count" db:"count"`
}
