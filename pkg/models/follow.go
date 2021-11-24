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
	CandidateID int64 `json:"candidate_id,omitempty"`
	RecruiterID int64 `json:"recruiter_id,omitempty"`
}

type RequestUnfollow struct {
	CandidateID int64 `json:"candidate_id"`
	RecruiterID int64 `json:"recruiter_id"`
}

type RequestGetCandidateFollow struct {
	RecruiterID int64 `json:"recruiter_id,omitempty"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type RequestGetRecruiterFollow struct {
	CandidateID int64 `json:"candidate_id,omitempty"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type ResponseCount struct {
	Count int64 `json:"count"`
}

type ResponseGetCandidate struct {
	Total      int64
	Candidates []*Candidate
}

type ResponseGetRecruiter struct {
	Total      int64
	Recruiters []*Recruiter
}
