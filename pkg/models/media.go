package models

import "time"

type Media struct {
	Media_id  int64     `json:"media_id" gorm:"primaryKey"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestCreateMedia struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// type RequestUpdateSetting struct {
// 	Type string `json:"type"`
// 	Name string `json:"name"`
// }

// type ResponsetListSetting struct {
// 	Total       int64     `json:"total"`
// 	TotalPage   float64   `json:"totalPage"`
// 	CurrentPage int64     `json:"currentPage"`
// 	Data        []Setting `json:"data"`
// }

// type RequestGetListSetting struct {
// 	Name string `json:"name"`
// 	Page int64  `json:"page"`
// 	Size int64  `json:"size"`
// }
