package models

import "time"

type Account struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique,index"`
	Password  string    `json:"password"`
	TokenHash string    `json:"token_hash"`
	Phone     string    `json:"phone"`
	Type      int64     `json:"type"`
	IsVerify  bool      `json:"is_verify"`
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

type RequestForgotPassword struct {
	Email string `json:"email"`
}

type RequestChangePassword struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type RequestResetPassword struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

type RequestVerifyEmail struct {
	Email string `json:"email"`
}
