package models

import "time"

type Recruiter struct {
	RecruiterID      int64     `json:"recruiter_id" gorm:"primaryKey"`
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	Avartar          string    `json:"avartar"`
	Banner           string    `json:"banner"`
	Phone            string    `json:"phone"`
	Website          string    `json:"website"`
	Description      string    `json:"description"`
	EmployeeQuantity string    `json:"employee_quantity"`
	ContacterName    string    `json:"contacter_name"`
	ContacterPhone   string    `json:"contacter_phone"`
	Media            string    `json:"media"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type RequestRecruiter struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	Avartar          string `json:"avartar"`
	Banner           string `json:"banner"`
	Phone            string `json:"phone"`
	Website          string `json:"website"`
	Description      string `json:"description"`
	EmployeeQuantity string `json:"employee_quantity"`
	ContacterName    string `json:"contacter_name"`
	ContacterPhone   string `json:"contacter_phone"`
	Media            string `json:"media"`
}

type RequestUpdateRecruiter struct {
	RecruiterID      int64  `json:"recruiter_id" gorm:"primaryKey"`
	Name             string `json:"name"`
	Address          string `json:"address"`
	Avartar          string `json:"avartar"`
	Banner           string `json:"banner"`
	Phone            string `json:"phone"`
	Website          string `json:"website"`
	Description      string `json:"description"`
	EmployeeQuantity string `json:"employee_quantity"`
	ContacterName    string `json:"contacter_name"`
	ContacterPhone   string `json:"contacter_phone"`
	Media            string `json:"media"`
}

type RequestGetRecruiter struct {
	RecruiterID int64 `json:"recruiter_id" gorm:"primaryKey"`
}

type RecruiterSkill struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	RecruiterId int64     `json:"recruiter_id"`
	SkillId     int64     `json:"skill_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseRecruiterSkill struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	RecruiterId int64     `json:"recruiter_id"`
	SkillId     int64     `json:"skill_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Questions   string    `json:"questions"`
	Icon        string    `json:"icon"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
