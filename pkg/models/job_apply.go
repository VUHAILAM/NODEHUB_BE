package models

import "time"

type JobApply struct {
	ID          int64     `json:"id"`
	CandidateID int64     `json:"candidate_id"`
	JobID       int64     `json:"job_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestApply struct {
	CandidateID int64  `json:"candidate_id"`
	JobID       int64  `json:"job_id"`
	Status      string `json:"status"`
}
