package models

import (
	"encoding/json"
	"time"
)

type CandidateRequest struct {
	CandidateID       int64         `json:"candidate_id"`
	FirstName         string        `json:"first_name,omitempty"`
	LastName          string        `json:"last_name,omitempty"`
	Gender            string        `json:"gender,omitempty"`
	BirthDay          string        `json:"birth_day,omitempty"`
	Address           string        `json:"address,omitempty"`
	Avatar            string        `json:"avatar,omitempty"`
	Banner            string        `json:"banner,omitempty"`
	Phone             string        `json:"phone,omitempty"`
	FindJob           *bool         `json:"find_job,omitempty"`
	NodehubReview     string        `json:"nodehub_review,omitempty"`
	NodehubScore      int           `json:"nodehub_score,omitempty"`
	CVManage          []CV          `json:"cv_manage,omitempty"`
	ExperienceManage  []Experience  `json:"experience_manage,omitempty"`
	EducationManage   []Education   `json:"education_manage,omitempty"`
	SocialManage      []Social      `json:"social_manage,omitempty"`
	ProjectManage     []Project     `json:"project_manage,omitempty"`
	CertificateManage []Certificate `json:"certificate_manage,omitempty"`
	PrizeManage       []Prize       `json:"prize_manage,omitempty"`
	CreatedAt         time.Time     `json:"created_at,omitempty"`
	UpdatedAt         time.Time     `json:"updated_at,omitempty"`
}

type CandidateResponse struct {
	CandidateID       int64         `json:"candidate_id"`
	Email             string        `json:"email"`
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	Gender            string        `json:"gender"`
	BirthDay          string        `json:"birth_day"`
	Address           string        `json:"address"`
	Avatar            string        `json:"avatar"`
	Banner            string        `json:"banner"`
	Phone             string        `json:"phone"`
	FindJob           *bool         `json:"find_job"`
	NodehubReview     string        `json:"nodehub_review"`
	NodehubScore      int           `json:"nodehub_score"`
	CVManage          []CV          `json:"cv_manage"`
	ExperienceManage  []Experience  `json:"experience_manage"`
	EducationManage   []Education   `json:"education_manage"`
	SocialManage      []Social      `json:"social_manage"`
	ProjectManage     []Project     `json:"project_manage"`
	CertificateManage []Certificate `json:"certificate_manage"`
	PrizeManage       []Prize       `json:"prize_manage"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type CandidateRequestAdmin struct {
	CandidateID       int64         `json:"candidate_id"`
	Email             string        `json:"email"`
	Fullname          string        `json:"full_name"`
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	Gender            string        `json:"gender"`
	BirthDay          string        `json:"birth_day"`
	Address           string        `json:"address"`
	Avatar            string        `json:"avatar"`
	Banner            string        `json:"banner"`
	Phone             string        `json:"phone"`
	FindJob           *bool         `json:"find_job"`
	NodehubReview     string        `json:"nodehub_review"`
	NodehubScore      int           `json:"nodehub_score"`
	CVManage          []CV          `json:"cv_manage"`
	ExperienceManage  []Experience  `json:"experience_manage"`
	EducationManage   []Education   `json:"education_manage"`
	SocialManage      []Social      `json:"social_manage"`
	ProjectManage     []Project     `json:"project_manage"`
	CertificateManage []Certificate `json:"certificate_manage"`
	PrizeManage       []Prize       `json:"prize_manage"`
	Status            bool          `json:"status"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
}

type CandidateAdmin struct {
	CandidateID       int64     `json:"candidate_id,omitempty" gorm:"primaryKey"`
	Email             string    `json:"email"`
	Fullname          string    `json:"full_name"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Gender            string    `json:"gender"`
	BirthDay          string    `json:"birth_day"`
	Address           string    `json:"address"`
	Avatar            string    `json:"avatar"`
	Banner            string    `json:"banner"`
	Phone             string    `json:"phone"`
	FindJob           *bool     `json:"find_job"`
	NodehubReview     string    `json:"nodehub_review"`
	NodehubScore      int       `json:"nodehub_score"`
	CvManage          string    `json:"cv_manage"`
	ExperienceManage  string    `json:"experience_manage"`
	EducationManage   string    `json:"education_manage"`
	SocialManage      string    `json:"social_manage"`
	ProjectManage     string    `json:"project_manage"`
	CertificateManage string    `json:"certificate_manage"`
	PrizeManage       string    `json:"prize_manage"`
	Status            bool      `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
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
	NodehubScore   int64  `json:"nodehub_score,omitempty" mapstructure:"nodehub_score,omitempty"`
}
type RequestUpdateStatusCandidate struct {
	ID     int64 `json:"id,omitempty" mapstructure:"id,omitempty"`
	Status bool  `json:"status,omitempty" mapstructure:"status,omitempty"`
}

type Candidate struct {
	CandidateID       int64     `json:"candidate_id,omitempty" gorm:"primaryKey"`
	Email             string    `json:"email" gorm:"-"`
	FirstName         string    `json:"first_name,omitempty"`
	LastName          string    `json:"last_name,omitempty"`
	Gender            string    `json:"gender,omitempty"`
	BirthDay          string    `json:"birth_day,omitempty"`
	Address           string    `json:"address,omitempty"`
	Avatar            string    `json:"avatar,omitempty"`
	Banner            string    `json:"banner,omitempty"`
	Phone             string    `json:"phone,omitempty"`
	FindJob           *bool     `json:"find_job,omitempty"`
	NodehubReview     string    `json:"nodehub_review,omitempty"`
	NodehubScore      int       `json:"nodehub_score,omitempty"`
	JobStatus         string    `json:"job_status" gorm:"->"`
	CvManage          string    `json:"cv_manage,omitempty"`
	ExperienceManage  string    `json:"experience_manage,omitempty"`
	EducationManage   string    `json:"education_manage,omitempty"`
	SocialManage      string    `json:"social_manage,omitempty"`
	ProjectManage     string    `json:"project_manage,omitempty"`
	CertificateManage string    `json:"certificate_manage,omitempty"`
	PrizeManage       string    `json:"prize_manage,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

type CandidateSkill struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	CandidateId int64     `json:"candidate_id"`
	SkillId     int64     `json:"skill_id"`
	Media       string    `json:"media,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestUpdateCandidateSkill struct {
	ID    int64  `json:"id"`
	Media string `json:"media,omitempty" mapstructure:"media,omitempty"`
}

type ResponseCandidateSkill struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	CandidateId int64     `json:"candidateId"`
	SkillId     int64     `json:"skill_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Questions   string    `json:"questions"`
	Icon        string    `json:"icon"`
	Status      string    `json:"status"`
	Media       string    `json:"media"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

type Education struct {
	School      string `json:"school,omitempty"`
	Degree      string `json:"degree,omitempty"`
	Field       string `json:"field,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	Description string `json:"description,omitempty"`
	Media       string `json:"media,omitempty"`
	Link        string `json:"link,omitempty"`
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
	eduJson, err := json.Marshal(req.EducationManage)
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
		Gender:        req.Gender,
		BirthDay:      req.BirthDay,
		Address:       req.Address,
		Avatar:        req.Avatar,
		Banner:        req.Banner,
		Phone:         req.Phone,
		FindJob:       req.FindJob,
		NodehubReview: req.NodehubReview,
		NodehubScore:  req.NodehubScore,
	}
	if req.CVManage != nil {
		candidate.CvManage = string(cvJson)
	}
	if req.ExperienceManage != nil {
		candidate.ExperienceManage = string(exJson)
	}
	if req.EducationManage != nil {
		candidate.EducationManage = string(eduJson)
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
	var eduManage []Education
	err = json.Unmarshal([]byte(c.EducationManage), &eduManage)
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
		Gender:            c.Gender,
		BirthDay:          c.BirthDay,
		Address:           c.Address,
		Avatar:            c.Avatar,
		Banner:            c.Banner,
		Phone:             c.Phone,
		FindJob:           c.FindJob,
		NodehubReview:     c.NodehubReview,
		NodehubScore:      c.NodehubScore,
		CVManage:          cvManage,
		ExperienceManage:  exManage,
		EducationManage:   eduManage,
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
	var eduManage []Education
	err = json.Unmarshal([]byte(c.EducationManage), &eduManage)
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
		Email:             c.Email,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		Gender:            c.Gender,
		BirthDay:          c.BirthDay,
		Address:           c.Address,
		Avatar:            c.Avatar,
		Banner:            c.Banner,
		Phone:             c.Phone,
		FindJob:           c.FindJob,
		NodehubReview:     c.NodehubReview,
		NodehubScore:      c.NodehubScore,
		CVManage:          cvManage,
		ExperienceManage:  exManage,
		EducationManage:   eduManage,
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
	var eduManage []Education
	err = json.Unmarshal([]byte(c.EducationManage), &eduManage)
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
		Gender:            c.Gender,
		BirthDay:          c.BirthDay,
		Address:           c.Address,
		Avatar:            c.Avatar,
		Banner:            c.Banner,
		Phone:             c.Phone,
		FindJob:           c.FindJob,
		NodehubReview:     c.NodehubReview,
		NodehubScore:      c.NodehubScore,
		CVManage:          cvManage,
		ExperienceManage:  exManage,
		EducationManage:   eduManage,
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

type RequestSearchCandidate struct {
	Text string `json:"text,omitempty"`
	Page int64  `json:"page"`
	Size int64  `json:"size"`
}

type ResponseSearchCandidate struct {
	Total      int64
	Candidates []CandidateWithSkill
}

type CandidateWithSkill struct {
	Candidate *Candidate `json:"candidate"`
	Skills    []*Skill   `json:"skills"`
}
