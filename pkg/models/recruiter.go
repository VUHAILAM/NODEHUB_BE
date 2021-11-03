package models

import "time"

type Recruiter struct {
	RecruiterID      int64     `json:"recruiter_id" gorm:"primaryKey"`
	CompanyName      string    `json:"company_name"`
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
	CompanyName      string `json:"company_name"`
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
