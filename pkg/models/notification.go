package models

import "time"

type Notification struct {
	Account_id int64     `json:"account_id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Key        string    `json:"key"`
	Check_read bool      `json:"check_read"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RequestCreateNotification struct {
	Account_id int64  `json:"account_id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Key        string `json:"key"`
	Check_read bool   `json:"check_read"`
}

type ResponsetListNotification struct {
	Total       int64          `json:"total`
	TotalPage   float64        `json:"totalPage"`
	CurrentPage int64          `json:"currentPage"`
	Data        []Notification `json:"data"`
}

type RequestGetListNotification struct {
	Account_id int64 `json:"account_id"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
}
