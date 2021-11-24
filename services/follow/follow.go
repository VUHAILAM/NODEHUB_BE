package follow

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type IFollowService interface {
	Follow(ctx context.Context, req models.RequestFollow) error
	UnFollow(ctx context.Context, req models.RequestUnfollow) error
	CountOfRecruiter(ctx context.Context, recruiterID int64) (*models.ResponseCount, error)
	CountOfCandidate(ctx context.Context, candidateID int64) (*models.ResponseCount, error)
	FollowExist(ctx context.Context, req models.RequestFollow) (*models.Follow, error)
	GetCandidate(ctx context.Context, req models.RequestGetCandidateFollow) (*models.ResponseGetCandidate, error)
	GetRecruiter(ctx context.Context, req models.RequestGetRecruiterFollow) (*models.ResponseGetRecruiter, error)
}

type FollowService struct {
	FollowGorm IFollowDatabase
	Logger     *zap.Logger
}

func NewFollowService(gorm *FollowGorm, logger *zap.Logger) *FollowService {
	return &FollowService{
		FollowGorm: gorm,
		Logger:     logger,
	}
}

func (s *FollowService) Follow(ctx context.Context, req models.RequestFollow) error {
	follow := models.Follow{
		CandidateID: req.CandidateID,
		RecruiterID: req.RecruiterID,
	}
	err := s.FollowGorm.Create(ctx, &follow)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *FollowService) UnFollow(ctx context.Context, req models.RequestUnfollow) error {
	follow := models.Follow{
		CandidateID: req.CandidateID,
		RecruiterID: req.RecruiterID,
	}
	err := s.FollowGorm.Delete(ctx, &follow)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *FollowService) CountOfRecruiter(ctx context.Context, recruiterID int64) (*models.ResponseCount, error) {
	count, err := s.FollowGorm.CountFollowOfRecruiter(ctx, recruiterID)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	return &models.ResponseCount{
		Count: count,
	}, nil
}

func (s *FollowService) CountOfCandidate(ctx context.Context, candidateID int64) (*models.ResponseCount, error) {
	count, err := s.FollowGorm.CountFollowOfCandidate(ctx, candidateID)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	return &models.ResponseCount{
		Count: count,
	}, nil
}

func (s *FollowService) FollowExist(ctx context.Context, req models.RequestFollow) (*models.Follow, error) {
	follow, err := s.FollowGorm.GetFollow(ctx, req.CandidateID, req.RecruiterID)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return follow, nil
}

func (s *FollowService) GetCandidate(ctx context.Context, req models.RequestGetCandidateFollow) (*models.ResponseGetCandidate, error) {
	offset := (req.Page - 1) * req.Size
	candidates, total, err := s.FollowGorm.GetFollowedRecruiter(ctx, req.RecruiterID, offset, req.Size)
	if err != nil {
		s.Logger.Error(err.Error(), zap.Int64("recruiter id", req.RecruiterID))
		return nil, err
	}

	resp := &models.ResponseGetCandidate{
		Total:      total,
		Candidates: candidates,
	}
	return resp, nil
}

func (s *FollowService) GetRecruiter(ctx context.Context, req models.RequestGetRecruiterFollow) (*models.ResponseGetRecruiter, error) {
	offset := (req.Page - 1) * req.Size
	recruiters, total, err := s.FollowGorm.GetFollowingRecruiter(ctx, req.CandidateID, offset, req.Size)
	if err != nil {
		s.Logger.Error(err.Error(), zap.Int64("recruiter id", req.CandidateID))
		return nil, err
	}

	resp := &models.ResponseGetRecruiter{
		Total:      total,
		Recruiters: recruiters,
	}
	return resp, nil
}
