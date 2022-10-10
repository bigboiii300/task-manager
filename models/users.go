package models

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	Username  string     `json:"username" gorm:"unique"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
