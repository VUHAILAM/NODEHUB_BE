package candidate

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type ICandidateService interface {
	CreateCandidateProfile(ctx context.Context, req models.CandidateRequest) (int64, error)
	UpdateCandidateProfile(ctx context.Context, req models.CandidateRequest) error
	GetCandidateProfile(ctx context.Context, candidateID int64) (*models.CandidateRequest, error)
}

type CandidateService struct {
	CanGorm ICandidateDatabase
	Logger  *zap.Logger
}

func NewCandidateService(gorm *CandidateGorm, logger *zap.Logger) *CandidateService {
	return &CandidateService{
		CanGorm: gorm,
		Logger:  logger,
	}
}

func (s *CandidateService) CreateCandidateProfile(ctx context.Context, req models.CandidateRequest) (int64, error) {
	candidate, err := req.ToCandidate()
	if err != nil {
		s.Logger.Error(err.Error())
		return 0, err
	}
	s.Logger.Info("Create: candidate models", zap.Reflect("candidate", candidate))
	canID, err := s.CanGorm.Create(ctx, &candidate)
	if err != nil {
		s.Logger.Error(err.Error())
		return 0, err
	}
	return canID, nil
}

func (s *CandidateService) UpdateCandidateProfile(ctx context.Context, req models.CandidateRequest) error {
	candidate, err := req.ToCandidate()
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	s.Logger.Info("Update: candidate models", zap.Reflect("candidate", candidate))
	err = s.CanGorm.Update(ctx, req.CandidateID, &candidate)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *CandidateService) GetCandidateProfile(ctx context.Context, candidateID int64) (*models.CandidateRequest, error) {
	candidate, err := s.CanGorm.GetByCandidateID(ctx, candidateID)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	req, err := candidate.ToCandidateRequest()
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return &req, nil
}
