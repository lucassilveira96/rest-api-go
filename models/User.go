package models

import (
	"time"
)

type User struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	password  string
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

var Users []User
