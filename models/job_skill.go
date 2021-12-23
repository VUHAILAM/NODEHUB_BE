package models

import "time"

type JobSkill struct {
	ID        int64     `json:"id"`
	SkillID   int64     `json:"skill_id"`
	JobID     int64     `json:"job_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestGetJobsBySkill struct {
	SkillID int64 `json:"skill_id"`
	Page    int64 `json:"page"`
	Size    int64 `json:"size"`
}

type ResponseGetJobsBySkill struct {
	Total  int64  `json:"total"`
	Result []*Job `json:"result"`
}

type RequestGetSkillsByJob struct {
	JobID int64 `json:"job_id"`
}
