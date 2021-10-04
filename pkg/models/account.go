package models

import "time"

type Account struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique,index"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Type      int64     `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestRegisterAccount struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Type     int64  `json:"type"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
