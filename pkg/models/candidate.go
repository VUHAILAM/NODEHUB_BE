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
	OriginCV string `json:"origin_cv"`
	MediaCV  string `json:"media_cv"`
}

type Experience struct {
	CompanyName string `json:"company_name"`
	WorkingRole string `json:"working_role"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
	Media       string `json:"media"`
	Link        string `json:"link"`
}

type Social struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Media       string `json:"media"`
	Link        string `json:"link"`
}

type Project struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
	Media       string `json:"media"`
	Link        string `json:"link"`
}

type Certificate struct {
	Name            string `json:"name"`
	Host            string `json:"host"`
	CertificateTime string `json:"certificate_time"`
	Media           string `json:"media"`
	Link            string `json:"link"`
}

type Prize struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	PrizeTime string `json:"prize_time"`
	Media     string `json:"media"`
	Link      string `json:"link"`
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
		CandidateID:       req.CandidateID,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		BirthDay:          req.BirthDay,
		Address:           req.Address,
		Avatar:            req.Avatar,
		Banner:            req.Banner,
		Phone:             req.Phone,
		FindJob:           req.FindJob,
		NodehubReview:     req.NodehubReview,
		CvManage:          string(cvJson),
		ExperienceManage:  string(exJson),
		SocialManage:      string(socialJson),
		ProjectManage:     string(prjJson),
		CertificateManage: string(certJson),
		PrizeManage:       string(prizeJson),
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
