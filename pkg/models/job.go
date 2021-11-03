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
	Expereience string    `json:"expereience"`
	Location    string    `json:"location"`
	HireDate    time.Time `json:"hire_date"`
	Status      bool      `json:"status"`
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
	Expereience string    `json:"expereience" mapstructure:"expereience"`
	Location    string    `json:"location" mapstructure:"location"`
	HireDate    time.Time `json:"hire_date" mapstructure:"hire_date"`
	Status      bool      `json:"status" mapstructure:"status"`
	CreateAt    time.Time `json:"create_at" mapstructure:"create_at"`
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
		Expereience: job.Expereience,
		Location:    job.Location,
		HireDate:    job.HireDate,
		Status:      job.Status,
		CreateAt:    job.CreatedAt,
	}
}

type CreateJobRequest struct {
	RecruiterID int64     `json:"recruiter_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SalaryRange string    `json:"salary_range"`
	Quantity    int64     `json:"quantity"`
	Role        string    `json:"role"`
	Expereience string    `json:"expereience"`
	Location    string    `json:"location"`
	HireDate    time.Time `json:"hire_date"`
	Status      bool      `json:"status"`
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
	Expereience string    `json:"expereience,omitempty" mapstructure:"expereience,omitempty"`
	Location    string    `json:"location,omitempty" mapstructure:"location,omitempty"`
	HireDate    time.Time `json:"hire_date,omitempty" mapstructure:"hire_date,omitempty"`
	Status      bool      `json:"status,omitempty" mapstructure:"status,omitempty"`
}

type RequestGetAllJob struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type ResponseGetAllJob struct {
	Total  int64 `json:"total"`
	Result []Job `json:"result"`
}
