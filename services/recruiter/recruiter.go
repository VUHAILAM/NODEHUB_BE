package recruiter

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IRecruiterService interface {
	AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error
	UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error
	GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error)
	GetProfileRecruiter(ctx context.Context, id int64) (*models.Recruiter, error)
}

type IRecruiterDatabase interface {
	AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error
	UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error
	GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error)
	GetProfile(ctx context.Context, id int64) (*models.Recruiter, error)
}

type Recruiter struct {
	RecruiterGorm *RecruiterGorm
	Logger        *zap.Logger
}

func NewRecruiterCategory(recruiterGorm *RecruiterGorm, logger *zap.Logger) *Recruiter {
	return &Recruiter{
		RecruiterGorm: recruiterGorm,
		Logger:        logger,
	}
}

func (r *Recruiter) GetProfileRecruiter(ctx context.Context, id int64) (*models.Recruiter, error) {
	acc, err := r.RecruiterGorm.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (r *Recruiter) UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error {
	recruiterModels := &models.RequestUpdateRecruiter{
		Name:             recruiter.Name,
		Address:          recruiter.Address,
		Avartar:          recruiter.Avartar,
		Banner:           recruiter.Banner,
		Phone:            recruiter.Phone,
		Website:          recruiter.Website,
		Description:      recruiter.Description,
		EmployeeQuantity: recruiter.EmployeeQuantity,
		ContacterName:    recruiter.ContacterName,
		ContacterPhone:   recruiter.ContacterPhone,
		Media:            recruiter.Media}
	err := r.RecruiterGorm.UpdateProfile(ctx, recruiterModels, recruiter_id)
	if err != nil {
		return err
	}
	return nil
}

// recruiterSkill
func (r *Recruiter) AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error {
	RecruiterSkillModels := &models.RecruiterSkill{
		Id:          recruiterSkill.Id,
		RecruiterId: recruiterSkill.RecruiterId,
		SkillId:     recruiterSkill.SkillId}
	err := r.RecruiterGorm.AddRecruiterSkill(ctx, RecruiterSkillModels)
	if err != nil {
		return err
	}
	return nil
}

func (r *Recruiter) GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error) {
	acc, err := r.RecruiterGorm.GetRecruiterSkill(ctx, recruiter_id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
