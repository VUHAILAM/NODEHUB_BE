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

type RequestGetJobApplyByJobID struct {
	JobID int64 `json:"job_id"`
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
}

type RequestGetJobApplyByCandidateID struct {
	CandidateID int64 `json:"candidate_id"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type RequestUpdateStatusJobApplied struct {
	JobID       int64  `json:"job_id"`
	CandidateID int64  `json:"candidate_id"`
	Status      string `json:"status"`
}

type ResponseGetJobApply struct {
	Total  int64  `json:"total"`
	Result []*Job `json:"result"`
}

type ResponseGetCandidateApply struct {
	Total  int64 `json:"total"`
	Result []*Candidate
}
