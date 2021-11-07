package models

import "time"

type JobSkill struct {
	ID        int64     `json:"id"`
	SkillID   int64     `json:"skill_id"`
	JobID     int64     `json:"job_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
