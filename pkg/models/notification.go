package models

import "time"

type Notification struct {
	NotificationID int64     `json:"notification_id" gorm:"primaryKey"`
	CandidateID    int64     `json:"candidate_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Key            string    `json:"key"`
	CheckRead      bool      `json:"check_read"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RequestCreateNotification struct {
	CandidateID int64  `json:"candidate_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Key         string `json:"key"`
	CheckRead   bool   `json:"check_read"`
}

type RequestGetListNotification struct {
	CandidateID int64 `json:"candidate_id"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type ResponseListNotification struct {
	Total         int64
	Notifications []*Notification
}
