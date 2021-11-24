package models

import "time"

type Follow struct {
	ID          int64     `json:"id"`
	CandidateID int64     `json:"candidate_id"`
	RecruiterID int64     `json:"recruiter_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestFollow struct {
	CandidateID int64 `json:"candidate_id"`
	RecruiterID int64 `json:"recruiter_id"`
}

type RequestUnfollow struct {
	CandidateID int64 `json:"candidate_id"`
	RecruiterID int64 `json:"recruiter_id"`
}
