package models

import "time"

type Job struct {
	JobID       int64     `json:"job_id" gorm:"primaryKey"`
	RecruiterID int64     `json:"recruiter_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SalaryRange string    `json:"salary_range"`
	Quantity    int64     `json:"quantity"`
	Role        string    `json:"role"`
	Experience  string    `json:"experience"`
	Location    string    `json:"location"`
	HireDate    time.Time `json:"hire_date"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ESJobCreate struct {
	JobID       int64     `json:"job_id" mapstructure:"job_id"`
	RecruiterID int64     `json:"recruiter_id" mapstructure:"recruiter_id"`
	Title       string    `json:"title" mapstructure:"title"`
	Description string    `json:"description" mapstructure:"description"`
	SalaryRange string    `json:"salary_range" mapstructure:"salary_range"`
	Quantity    int64     `json:"quantity" mapstructure:"quantity"`
	Role        string    `json:"role" mapstructure:"role"`
	Experience  string    `json:"experience" mapstructure:"experience"`
	Location    string    `json:"location" mapstructure:"location"`
	HireDate    time.Time `json:"hire_date" mapstructure:"hire_date"`
	Status      int       `json:"status" mapstructure:"status"`
	Skills      []ESSkill `json:"skills" mapstructure:"skills"`
}

func ToESJobCreate(job *Job) *ESJobCreate {
	return &ESJobCreate{
		JobID:       job.JobID,
		RecruiterID: job.RecruiterID,
		Title:       job.Title,
		Description: job.Description,
		SalaryRange: job.SalaryRange,
		Quantity:    job.Quantity,
		Role:        job.Role,
		Experience:  job.Experience,
		Location:    job.Location,
		HireDate:    job.HireDate,
		Status:      job.Status,
	}
}

type CreateJobRequest struct {
	RecruiterID int64     `json:"recruiter_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SalaryRange string    `json:"salary_range"`
	Quantity    int64     `json:"quantity"`
	Role        string    `json:"role"`
	Experience  string    `json:"experience"`
	Location    string    `json:"location"`
	HireDate    time.Time `json:"hire_date"`
	Status      int       `json:"status"`
	SkillIDs    []int64   `json:"skill_ids"`
}

type RequestGetJobDetail struct {
	JobID int64 `json:"job_id"`
}

type RequestUpdateJob struct {
	JobID       int64     `json:"job_id,omitempty" mapstructure:"job_id,omitempty"`
	RecruiterID int64     `json:"recruiter_id,omitempty" mapstructure:"recruiter_id,omitempty"`
	Title       string    `json:"title,omitempty" mapstructure:"title,omitempty"`
	Description string    `json:"description,omitempty" mapstructure:"description,omitempty"`
	SalaryRange string    `json:"salary_range,omitempty" mapstructure:"salary_range,omitempty"`
	Quantity    int64     `json:"quantity,omitempty" mapstructure:"quantity,omitempty"`
	Role        string    `json:"role,omitempty" mapstructure:"role,omitempty"`
	Experience  string    `json:"experience,omitempty" mapstructure:"experience,omitempty"`
	Location    string    `json:"location,omitempty" mapstructure:"location,omitempty"`
	Status      int       `json:"status,omitempty" mapstructure:"status,omitempty"`
	HireDate    time.Time `json:"hire_date,omitempty" mapstructure:"hire_date,omitempty"`
}

type RequestGetAllJob struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type ResponseGetAllJob struct {
	Total  int64 `json:"total"`
	Result []Job `json:"result"`
}

type JobForAdmin struct {
	JobID         int64     `json:"job_id" gorm:"primaryKey"`
	RecruiterID   int64     `json:"recruiter_id"`
	RecruiterName string    `json:"recruiter_name"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	SalaryRange   string    `json:"salary_range"`
	Quantity      int64     `json:"quantity"`
	Role          string    `json:"role"`
	Experience    string    `json:"experience"`
	Location      string    `json:"location"`
	HireDate      int64     `json:"hire_date"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RequestGetListJobAdmin struct {
	Name string `json:"name"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type ResponsetListJobAdmin struct {
	Total       int64         `json:"total"`
	TotalPage   float64       `json:"totalPage"`
	CurrentPage int64         `json:"currentPage"`
	Data        []JobForAdmin `json:"data"`
}

type RequestUpdateStatusJob struct {
	JobID  int64 `json:"job_id"`
	Status bool  `json:"status"`
}
