package models

import (
	"time"
)

type Skill struct {
	Skill_id    int64     `json:"skill_id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Questions   string    `json:"questions"`
	Icon        string    `json:"icon"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestCreateSkill struct {
	Skill_id    int64  `json:"skill_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Questions   string `json:"questions"`
	Icon        string `json:"icon"`
	Status      bool   `json:"status"`
}

type RequestUpdateSkill struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Questions   string `json:"questions"`
	Icon        string `json:"icon"`
	Status      bool   `json:"status"`
}

type ResponsetListSkill struct {
	TotalSkill  int64   `json:"totalSkill"`
	TotalPage   float64 `json:"totalPage"`
	CurrentPage int64   `json:"currentPage"`
	Data        []Skill `json:"data"`
}

type RequestGetListSkill struct {
	Name string `json:"name"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type ESSkill struct {
	SkillID     int64  `json:"skill_id" mapstructure:"skill_id"`
	Name        string `json:"name" mapstructure:"name"`
	Description string `json:"description" mapstructure:"description"`
	Questions   string `json:"questions" mapstructure:"questions"`
	Icon        string `json:"icon" mapstructure:"icon"`
	Status      bool   `json:"status" mapstructure:"status"`
}

func ToESSkill(skill *Skill) ESSkill {
	return ESSkill{
		SkillID:     skill.Skill_id,
		Name:        skill.Name,
		Description: skill.Description,
		Questions:   skill.Questions,
		Icon:        skill.Icon,
		Status:      skill.Status,
	}
}

// type RequestGetAllSkill struct {
// 	Name string `json:"name"`
// }
