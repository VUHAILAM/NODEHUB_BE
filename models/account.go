package models

import "time"

type Account struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique,index"`
	Password  string    `json:"password,omitempty"`
	TokenHash string    `json:"token_hash,omitempty"`
	Phone     string    `json:"phone"`
	Type      int64     `json:"type"`
	RoleName  string    `json:"role_name,omitempty" gorm:"-"`
	FullName  string    `json:"full_name,omitempty" gorm:"-"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestRegisterAccount struct {
	Email          string           `json:"email"`
	Phone          string           `json:"phone"`
	Password       string           `json:"password"`
	Type           int64            `json:"type"`
	RecruiterInfor RequestRecruiter `json:"recruiter_infor,omitempty"`
	CandidateInfor CandidateRequest `json:"candidate_infor,omitempty"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     int64  `json:"type"`
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

type RequestPublicProfile struct {
	ID int64 `json:"id"`
}
