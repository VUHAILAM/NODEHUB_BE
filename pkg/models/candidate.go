package models

import (
	"encoding/json"
	"time"
)

type CandidateRequest struct {
	CandidateID       int64         `json:"candidate_id"`
	FirstName         string        `json:"first_name,omitempty"`
	LastName          string        `json:"last_name,omitempty"`
	BirthDay          string        `json:"birth_day,omitempty"`
	Address           string        `json:"address,omitempty"`
	Avatar            string        `json:"avatar,omitempty"`
	Banner            string        `json:"banner,omitempty"`
	Phone             string        `json:"phone,omitempty"`
	FindJob           bool          `json:"find_job,omitempty"`
	NodehubReview     string        `json:"nodehub_review,omitempty"`
	CVManage          []CV          `json:"cv_manage,omitempty"`
	ExperienceManage  []Experience  `json:"experience_manage,omitempty"`
	SocialManage      []Social      `json:"social_manage,omitempty"`
	ProjectManage     []Project     `json:"project_manage,omitempty"`
	CertificateManage []Certificate `json:"certificate_manage,omitempty"`
	PrizeManage       []Prize       `json:"prize_manage,omitempty"`
	CreatedAt         time.Time     `json:"created_at,omitempty"`
	UpdatedAt         time.Time     `json:"updated_at,omitempty"`
}

type CandidateResponse struct {
	CandidateID       int64         `json:"candidate_id"`
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	BirthDay          string        `json:"birth_day"`
	Address           string        `json:"address"`
	Avatar            string        `json:"avatar"`
	Banner            string        `json:"banner"`
	Phone             string        `json:"phone"`
	FindJob           bool          `json:"find_job"`
	NodehubReview     string        `json:"nodehub_review"`
	CVManage          []CV          `json:"cv_manage"`
	ExperienceManage  []Experience  `json:"experience_manage"`
	SocialManage      []Social      `json:"social_manage"`
	ProjectManage     []Project     `json:"project_manage"`
	CertificateManage []Certificate `json:"certificate_manage"`
	PrizeManage       []Prize       `json:"prize_manage"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type CandidateRequestAdmin struct {
	CandidateID       int64         `json:"candidate_id"`
	Email             string        `json:"email,omitempty"`
	Fullname          string        `json:"full_name,omitempty"`
	FirstName         string        `json:"first_name,omitempty"`
	LastName          string        `json:"last_name,omitempty"`
	BirthDay          string        `json:"birth_day,omitempty"`
	Address           string        `json:"address,omitempty"`
	Avatar            string        `json:"avatar,omitempty"`
	Banner            string        `json:"banner,omitempty"`
	Phone             string        `json:"phone,omitempty"`
	FindJob           bool          `json:"find_job,omitempty"`
	NodehubReview     string        `json:"nodehub_review,omitempty"`
	CVManage          []CV          `json:"cv_manage,omitempty"`
	ExperienceManage  []Experience  `json:"experience_manage,omitempty"`
	SocialManage      []Social      `json:"social_manage,omitempty"`
	ProjectManage     []Project     `json:"project_manage,omitempty"`
	CertificateManage []Certificate `json:"certificate_manage,omitempty"`
	PrizeManage       []Prize       `json:"prize_manage,omitempty"`
	Status            bool          `json:"status"`
	CreatedAt         time.Time     `json:"created_at,omitempty"`
	UpdatedAt         time.Time     `json:"updated_at,omitempty"`
}

type CandidateAdmin struct {
	CandidateID       int64     `json:"candidate_id,omitempty" gorm:"primaryKey"`
	Email             string    `json:"email,omitempty"`
	Fullname          string    `json:"full_name,omitempty"`
	FirstName         string    `json:"first_name,omitempty"`
	LastName          string    `json:"last_name,omitempty"`
	BirthDay          string    `json:"birth_day,omitempty"`
	Address           string    `json:"address,omitempty"`
	Avatar            string    `json:"avatar,omitempty"`
	Banner            string    `json:"banner,omitempty"`
	Phone             string    `json:"phone,omitempty"`
	FindJob           bool      `json:"find_job,omitempty"`
	NodehubReview     string    `json:"nodehub_review,omitempty"`
	CvManage          string    `json:"cv_manage,omitempty"`
	ExperienceManage  string    `json:"experience_manage,omitempty"`
	SocialManage      string    `json:"social_manage,omitempty"`
	ProjectManage     string    `json:"project_manage,omitempty"`
	CertificateManage string    `json:"certificate_manage,omitempty"`
	PrizeManage       string    `json:"prize_manage,omitempty"`
	Status            bool      `json:"status"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type ResponsetListCandidateAdmin struct {
	Total       int64                   `json:"total"`
	TotalPage   float64                 `json:"totalPage"`
	CurrentPage int64                   `json:"currentPage"`
	Data        []CandidateRequestAdmin `json:"data"`
}

type RequestGetListCandidateAdmin struct {
	Name string `json:"name"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type RequestUpdateReviewCandidateAdmin struct {
	CandidateID    int64  `json:"candidate_id,omitempty" mapstructure:"candidate_id,omitempty"`
	Nodehub_review string `json:"nodehub_review,omitempty" mapstructure:"nodehub_review,omitempty"`
}
type RequestUpdateStatusCandidate struct {
	ID     int64 `json:"id,omitempty" mapstructure:"id,omitempty"`
	Status bool  `json:"status,omitempty" mapstructure:"status,omitempty"`
}

type Candidate struct {
	CandidateID       int64     `json:"candidate_id,omitempty" gorm:"primaryKey"`
	FirstName         string    `json:"first_name,omitempty"`
	LastName          string    `json:"last_name,omitempty"`
	BirthDay          string    `json:"birth_day,omitempty"`
	Address           string    `json:"address,omitempty"`
	Avatar            string    `json:"avatar,omitempty"`
	Banner            string    `json:"banner,omitempty"`
	Phone             string    `json:"phone,omitempty"`
	FindJob           bool      `json:"find_job,omitempty"`
	NodehubReview     string    `json:"nodehub_review,omitempty"`
	CvManage          string    `json:"cv_manage,omitempty"`
	ExperienceManage  string    `json:"experience_manage,omitempty"`
	SocialManage      string    `json:"social_manage,omitempty"`
	ProjectManage     string    `json:"project_manage,omitempty"`
	CertificateManage string    `json:"certificate_manage,omitempty"`
	PrizeManage       string    `json:"prize_manage,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type CV struct {
	OriginCV string `json:"origin_cv,omitempty"`
	MediaCV  string `json:"media_cv,omitempty"`
}

type Experience struct {
	CompanyName string `json:"company_name,omitempty"`
	WorkingRole string `json:"working_role,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Description string `json:"description,omitempty"`
	Media       string `json:"media,omitempty"`
	Link        string `json:"link,omitempty"`
}

type Social struct {
	Name        string `json:"name,omitempty"`
	Role        string `json:"role,omitempty"`
	Description string `json:"description,omitempty"`
	Media       string `json:"media,omitempty"`
	Link        string `json:"link,omitempty"`
}

type Project struct {
	Name        string `json:"name,omitempty"`
	Role        string `json:"role,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	Description string `json:"description,omitempty"`
	Media       string `json:"media,omitempty"`
	Link        string `json:"link,omitempty"`
}

type Certificate struct {
	Name            string `json:"name,omitempty"`
	Host            string `json:"host,omitempty"`
	CertificateTime string `json:"certificate_time,omitempty"`
	Media           string `json:"media,omitempty"`
	Link            string `json:"link,omitempty"`
}

type Prize struct {
	Name      string `json:"name,omitempty"`
	Host      string `json:"host,omitempty"`
	PrizeTime string `json:"prize_time,omitempty"`
	Media     string `json:"media,omitempty"`
	Link      string `json:"link,omitempty"`
}

func (req *CandidateRequest) ToCandidate() (Candidate, error) {
	cvJson, err := json.Marshal(req.CVManage)
	if err != nil {
		return Candidate{}, err
	}
	exJson, err := json.Marshal(req.ExperienceManage)
	if err != nil {
		return Candidate{}, err
	}
	socialJson, err := json.Marshal(req.SocialManage)
	if err != nil {
		return Candidate{}, err
	}
	prjJson, err := json.Marshal(req.ProjectManage)
	if err != nil {
		return Candidate{}, err
	}
	certJson, err := json.Marshal(req.CertificateManage)
	if err != nil {
		return Candidate{}, err
	}
	prizeJson, err := json.Marshal(req.PrizeManage)
	if err != nil {
		return Candidate{}, err
	}
	candidate := Candidate{
		CandidateID:   req.CandidateID,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		BirthDay:      req.BirthDay,
		Address:       req.Address,
		Avatar:        req.Avatar,
		Banner:        req.Banner,
		Phone:         req.Phone,
		FindJob:       req.FindJob,
		NodehubReview: req.NodehubReview,
	}
	if req.CVManage != nil {
		candidate.CvManage = string(cvJson)
	}
	if req.ExperienceManage != nil {
		candidate.ExperienceManage = string(exJson)
	}
	if req.SocialManage != nil {
		candidate.SocialManage = string(socialJson)
	}
	if req.ProjectManage != nil {
		candidate.ProjectManage = string(prjJson)
	}
	if req.CertificateManage != nil {
		candidate.CertificateManage = string(certJson)
	}
	if req.PrizeManage != nil {
		candidate.PrizeManage = string(prizeJson)
	}
	return candidate, nil
}

func (c *Candidate) ToCandidateRequest() (CandidateRequest, error) {
	var cvManage []CV
	err := json.Unmarshal([]byte(c.CvManage), &cvManage)
	if err != nil {
		return CandidateRequest{}, err
	}
	var exManage []Experience
	err = json.Unmarshal([]byte(c.ExperienceManage), &exManage)
	if err != nil {
		return CandidateRequest{}, err
	}
	var socialManage []Social
	err = json.Unmarshal([]byte(c.SocialManage), &socialManage)
	if err != nil {
		return CandidateRequest{}, err
	}
	var prjManage []Project
	err = json.Unmarshal([]byte(c.ProjectManage), &prjManage)
	if err != nil {
		return CandidateRequest{}, err
	}
	var certManage []Certificate
	err = json.Unmarshal([]byte(c.CertificateManage), &certManage)
	if err != nil {
		return CandidateRequest{}, err
	}
	var prizeManage []Prize
	err = json.Unmarshal([]byte(c.PrizeManage), &prizeManage)
	if err != nil {
		return CandidateRequest{}, err
	}

	req := CandidateRequest{
		CandidateID:       c.CandidateID,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		BirthDay:          c.BirthDay,
		Address:           c.Address,
		Avatar:            c.Avatar,
		Banner:            c.Banner,
		Phone:             c.Phone,
		FindJob:           c.FindJob,
		NodehubReview:     c.NodehubReview,
		CVManage:          cvManage,
		ExperienceManage:  exManage,
		SocialManage:      socialManage,
		ProjectManage:     prjManage,
		CertificateManage: certManage,
		PrizeManage:       prizeManage,
		CreatedAt:         c.CreatedAt,
		UpdatedAt:         c.UpdatedAt,
	}
	return req, nil
}

func (c *Candidate) ToCandidateResponse() (CandidateResponse, error) {
	var cvManage []CV
	err := json.Unmarshal([]byte(c.CvManage), &cvManage)
	if err != nil {
		return CandidateResponse{}, err
	}
	var exManage []Experience
	err = json.Unmarshal([]byte(c.ExperienceManage), &exManage)
	if err != nil {
		return CandidateResponse{}, err
	}
	var socialManage []Social
	err = json.Unmarshal([]byte(c.SocialManage), &socialManage)
	if err != nil {
		return CandidateResponse{}, err
	}
	var prjManage []Project
	err = json.Unmarshal([]byte(c.ProjectManage), &prjManage)
	if err != nil {
		return CandidateResponse{}, err
	}
	var certManage []Certificate
	err = json.Unmarshal([]byte(c.CertificateManage), &certManage)
	if err != nil {
		return CandidateResponse{}, err
	}
	var prizeManage []Prize
	err = json.Unmarshal([]byte(c.PrizeManage), &prizeManage)
	if err != nil {
		return CandidateResponse{}, err
	}

	req := CandidateResponse{
		CandidateID:       c.CandidateID,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		BirthDay:          c.BirthDay,
		Address:           c.Address,
		Avatar:            c.Avatar,
		Banner:            c.Banner,
		Phone:             c.Phone,
		FindJob:           c.FindJob,
		NodehubReview:     c.NodehubReview,
		CVManage:          cvManage,
		ExperienceManage:  exManage,
		SocialManage:      socialManage,
		ProjectManage:     prjManage,
		CertificateManage: certManage,
		PrizeManage:       prizeManage,
		CreatedAt:         c.CreatedAt,
		UpdatedAt:         c.UpdatedAt,
	}
	return req, nil
}

func (c *CandidateAdmin) ToCandidateRequestAdmin() (CandidateRequestAdmin, error) {
	var cvManage []CV
	err := json.Unmarshal([]byte(c.CvManage), &cvManage)
	if err != nil {
		return CandidateRequestAdmin{}, err
	}
	var exManage []Experience
	err = json.Unmarshal([]byte(c.ExperienceManage), &exManage)
	if err != nil {
		return CandidateRequestAdmin{}, err
	}
	var socialManage []Social
	err = json.Unmarshal([]byte(c.SocialManage), &socialManage)
	if err != nil {
		return CandidateRequestAdmin{}, err
	}
	var prjManage []Project
	err = json.Unmarshal([]byte(c.ProjectManage), &prjManage)
	if err != nil {
		return CandidateRequestAdmin{}, err
	}
	var certManage []Certificate
	err = json.Unmarshal([]byte(c.CertificateManage), &certManage)
	if err != nil {
		return CandidateRequestAdmin{}, err
	}
	var prizeManage []Prize
	err = json.Unmarshal([]byte(c.PrizeManage), &prizeManage)
	if err != nil {
		return CandidateRequestAdmin{}, err
	}

	req := CandidateRequestAdmin{
		Email:             c.Email,
		Fullname:          c.Fullname,
		CandidateID:       c.CandidateID,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		BirthDay:          c.BirthDay,
		Address:           c.Address,
		Avatar:            c.Avatar,
		Banner:            c.Banner,
		Phone:             c.Phone,
		FindJob:           c.FindJob,
		NodehubReview:     c.NodehubReview,
		CVManage:          cvManage,
		ExperienceManage:  exManage,
		SocialManage:      socialManage,
		ProjectManage:     prjManage,
		CertificateManage: certManage,
		PrizeManage:       prizeManage,
		Status:            c.Status,
		CreatedAt:         c.CreatedAt,
		UpdatedAt:         c.UpdatedAt,
	}
	return req, nil
}
