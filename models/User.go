package models

import (
	"github.com/jinzhu/gorm"
)

// User : user model
type User struct {
	gorm.Model
	Login    string
	Password string
	Enable   bool
	Info     string
}

//UserRequest : user request
type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Enable   bool   `json:"enable"`
	Info     string `json:"info"`
}

//UserResponse : user Response
type UserResponse struct {
	Error       bool   `json:"error"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}
