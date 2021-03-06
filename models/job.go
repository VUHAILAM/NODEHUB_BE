package models

import (
	"encoding/json"
	"strings"
	"time"
)

type HDate time.Time

type Job struct {
	JobID           int64     `json:"job_id" gorm:"primaryKey"`
	RecruiterID     int64     `json:"recruiter_id"`
	CompanyName     string    `json:"company_name" gorm:"->"`
	Avatar          string    `json:"avatar" gorm:"->"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	SalaryRange     string    `json:"salary_range"`
	Quantity        int64     `json:"quantity"`
	Role            string    `json:"role"`
	Experience      string    `json:"experience"`
	Location        string    `json:"location"`
	HireDate        time.Time `json:"hire_date"`
	Status          int       `json:"status"`
	Questions       string    `json:"questions"`
	CandidateStatus string    `json:"candidate_status" gorm:"->"`
	Skills          []ESSkill `json:"skills" gorm:"-"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ESJobCreate struct {
	JobID       int64     `json:"job_id" mapstructure:"job_id"`
	RecruiterID int64     `json:"recruiter_id" mapstructure:"recruiter_id"`
	CompanyName string    `json:"company_name" mapstructure:"company_name"`
	Avatar      string    `json:"avatar" mapstructure:"avatar"`
	Title       string    `json:"title" mapstructure:"title"`
	Description string    `json:"description" mapstructure:"description"`
	SalaryRange string    `json:"salary_range" mapstructure:"salary_range"`
	Quantity    int64     `json:"quantity" mapstructure:"quantity"`
	Role        string    `json:"role" mapstructure:"role"`
	Experience  string    `json:"experience" mapstructure:"experience"`
	Location    string    `json:"location" mapstructure:"location"`
	HireDate    string    `json:"hire_date" mapstructure:"hire_date"`
	CreateAt    string    `json:"created_at" mapstructure:"created_at"`
	Status      int       `json:"status" mapstructure:"status"`
	Questions   string    `json:"questions" mapstructure:"questions"`
	Skills      []ESSkill `json:"skills" mapstructure:"skills"`
}

func ToESJobCreate(job *Job) *ESJobCreate {
	return &ESJobCreate{
		JobID:       job.JobID,
		RecruiterID: job.RecruiterID,
		CompanyName: job.CompanyName,
		Avatar:      job.Avatar,
		Title:       job.Title,
		Description: job.Description,
		SalaryRange: job.SalaryRange,
		Quantity:    job.Quantity,
		Role:        job.Role,
		Experience:  job.Experience,
		Location:    job.Location,
		HireDate:    job.HireDate.Format("2006-01-02"),
		CreateAt:    job.CreatedAt.Format("2006-01-02"),
		Questions:   job.Questions,
		Status:      job.Status,
	}
}

type ESJobUpdate struct {
	JobID       int64     `json:"job_id,omitempty" mapstructure:"job_id,omitempty"`
	RecruiterID int64     `json:"recruiter_id,omitempty" mapstructure:"recruiter_id,omitempty"`
	CompanyName string    `json:"company_name,omitempty" mapstructure:"company_name,omitempty"`
	Avatar      string    `json:"avatar,omitempty" mapstructure:"avatar,omitempty"`
	Title       string    `json:"title,omitempty" mapstructure:"title,omitempty"`
	Description string    `json:"description,omitempty" mapstructure:"description,omitempty"`
	SalaryRange string    `json:"salary_range,omitempty" mapstructure:"salary_range,omitempty"`
	Quantity    int64     `json:"quantity,omitempty" mapstructure:"quantity,omitempty"`
	Role        string    `json:"role,omitempty" mapstructure:"role,omitempty"`
	Experience  string    `json:"experience,omitempty" mapstructure:"experience,omitempty"`
	Location    string    `json:"location,omitempty" mapstructure:"location,omitempty"`
	HireDate    string    `json:"hire_date,omitempty" mapstructure:"hire_date,omitempty"`
	Status      int       `json:"status,omitempty" mapstructure:"status,omitempty"`
	Questions   string    `json:"questions,omitempty" mapstructure:"questions,omitempty"`
	Skills      []ESSkill `json:"skills,omitempty" mapstructure:"skills,omitempty"`
}

type ESJob struct {
	JobID       int64     `json:"job_id"`
	RecruiterID int64     `json:"recruiter_id"`
	CompanyName string    `json:"company_name"`
	Avatar      string    `json:"avatar"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SalaryRange string    `json:"salary_range"`
	Quantity    int64     `json:"quantity"`
	Role        string    `json:"role"`
	Experience  string    `json:"experience"`
	Location    string    `json:"location"`
	HireDate    HDate     `json:"hire_date"`
	Status      int       `json:"status"`
	Questions   string    `json:"questions"`
	Skills      []ESSkill `json:"skills"`
	CreatedAt   HDate     `json:"created_at"`
}

type CreateJobRequest struct {
	RecruiterID int64   `json:"recruiter_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	SalaryRange string  `json:"salary_range"`
	Quantity    int64   `json:"quantity"`
	Role        string  `json:"role"`
	Experience  string  `json:"experience"`
	Location    string  `json:"location"`
	HireDate    HDate   `json:"hire_date"`
	Status      int     `json:"status"`
	Questions   string  `json:"questions"`
	SkillIDs    []int64 `json:"skill_ids"`
}

func (j *HDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = HDate(t)
	return nil
}

func (j HDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j HDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

type RequestGetJobDetail struct {
	JobID int64 `json:"job_id"`
}

type RequestUpdateJob struct {
	JobID       int64   `json:"job_id,omitempty"`
	RecruiterID int64   `json:"recruiter_id,omitempty"`
	CompanyName string  `json:"company_name,omitempty"`
	Avatar      string  `json:"avatar,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	SalaryRange string  `json:"salary_range,omitempty"`
	Quantity    int64   `json:"quantity,omitempty"`
	Role        string  `json:"role,omitempty"`
	Experience  string  `json:"experience,omitempty"`
	Location    string  `json:"location,omitempty"`
	Status      int     `json:"status,omitempty"`
	Questions   string  `json:"questions,omitempty"`
	HireDate    HDate   `json:"hire_date,omitempty"`
	SkillIDs    []int64 `json:"skill_ids,omitempty"`
}

type RequestGetAllJob struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type RequestGetJobsByRecruiter struct {
	RecruiterID int64 `json:"recruiter_id"`
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
}

type ResponseGetJob struct {
	Total  int64   `json:"total"`
	Result []ESJob `json:"result"`
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
	HireDate      time.Time `json:"hire_date"`
	Status        int64     `json:"status"`
	Questions     string    `json:"questions"`
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
	Status int64 `json:"status"`
}

type RequestSearchJob struct {
	Text     string `json:"text"`
	Location string `json:"location"`
	Page     int64  `json:"page"`
	Size     int64  `json:"size"`
}
