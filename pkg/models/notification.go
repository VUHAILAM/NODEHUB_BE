package models

import "time"

type Notification struct {
	NotificationID int64     `json:"notification_id" gorm:"primaryKey"`
	CandidateID    int64     `json:"candidate_id"`
	RecruiterID    int64     `json:"recruiter_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Key            string    `json:"key"`
	CheckRead      bool      `json:"check_read"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RequestCreateNotification struct {
	RecruiterID int64  `json:"recruiter_id"`
	CandidateID int64  `json:"candidate_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Key         string `json:"key"`
	CheckRead   bool   `json:"check_read"`
}

type RequestGetListNotification struct {
	RecruiterID int64 `json:"recruiter_id,omitempty"`
	CandidateID int64 `json:"candidate_id,omitempty"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type ResponseListNotification struct {
	Total         int64
	Notifications []*Notification
}

type RequestMarkRead struct {
	NotificationID int64 `json:"notification_id"`
}

type RequestMarkReadAll struct {
	AccountID int64 `json:"account_id"`
	Role      int64 `json:"role"`
}
