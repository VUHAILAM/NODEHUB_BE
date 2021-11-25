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
	Premium          bool      `json:"premium"`
	Nodehub_review   string    `json:"nodehub_review"`
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
	Premium          bool   `json:"premium"`
	Nodehub_review   string `json:"nodehub_review"`
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

type RequestGetListRecruiter struct {
	Name string `json:"name"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type RecruiterForAdmin struct {
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
	Premium          bool      `json:"premium"`
	Nodehub_review   string    `json:"nodehub_review"`
	Status           bool      `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ResponsetListRecruiter struct {
	Total       int64               `json:"total"`
	TotalPage   float64             `json:"totalPage"`
	CurrentPage int64               `json:"currentPage"`
	Data        []RecruiterForAdmin `json:"data"`
}

type RequestUpdateRecruiterAdmin struct {
	RecruiterID    int64  `json:"recruiter_id,omitempty" mapstructure:"recruiter_id,omitempty"`
	Premium        bool   `json:"premium" mapstructure:"premium"`
	Nodehub_review string `json:"nodehub_review,omitempty" mapstructure:"nodehub_review,omitempty"`
}

type RequestUpdateStatusRecruiter struct {
	ID     int64 `json:"id,omitempty" mapstructure:"id,omitempty"`
	Status bool  `json:"status,omitempty" mapstructure:"status,omitempty"`
}

type RequestGetListRecruiterForCandidate struct {
	RecruiterName string `json:"recruiterName"`
	SkillName     string `json:"skillName"`
	Address       string `json:"address"`
	Page          int64  `json:"page"`
	Size          int64  `json:"size"`
}

type SkilList struct {
	Skill_name string `json:"skill_name"`
	Skill_icon string `json:"skill_icon"`
}

type RecruiterForCandidate struct {
	RecruiterID int64      `json:"recruiter_id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Skills      []SkilList `json:"skills"`
	Address     string     `json:"address"`
	Avartar     string     `json:"avartar"`
	Banner      string     `json:"banner"`
	Description string     `json:"description"`
}

type ResponsetListRecruiterForCandidate struct {
	Total       int64                        `json:"total"`
	TotalPage   float64                      `json:"totalPage"`
	CurrentPage int64                        `json:"currentPage"`
	Data        []RecruiterForCandidateCheck `json:"data"`
}

type RecruiterForCandidateCheck struct {
	RecruiterID int64  `json:"recruiter_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Skill_name  string `json:"skill_name"`
	Skill_icon  string `json:"skill_icon"`
	Address     string `json:"address"`
	Avartar     string `json:"avartar"`
	Banner      string `json:"banner"`
	Description string `json:"description"`
}

type RequestSearchRecruiter struct {
	Text string `json:"text,omitempty"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type ResponseSearchRecruiter struct {
	Total      int64
	Recruiters []RecruiterWithSkill
}

type RecruiterWithSkill struct {
	Recruiter *Recruiter `json:"recruiter"`
	Skills    []*Skill   `json:"skills"`
}
