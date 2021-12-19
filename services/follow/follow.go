package follow

import (
	"context"
	"strconv"

	"gitlab.com/hieuxeko19991/job4e_be/services/candidate"
	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"

	"gitlab.com/hieuxeko19991/job4e_be/services/notification"

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
	FollowGorm    IFollowDatabase
	NotiGorm      notification.INotificationDatabase
	CandidateGorm candidate.ICandidateDatabase
	RecruiterGorm recruiter.IRecruiterDatabase

	Logger *zap.Logger
}

func NewFollowService(gorm *FollowGorm, notiGorm *notification.NotificationGorm, canGorm *candidate.CandidateGorm, recruiterGorm *recruiter.RecruiterGorm, logger *zap.Logger) *FollowService {
	return &FollowService{
		FollowGorm:    gorm,
		NotiGorm:      notiGorm,
		CandidateGorm: canGorm,
		RecruiterGorm: recruiterGorm,
		Logger:        logger,
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
	candidateInfor, err := s.CandidateGorm.GetByCandidateID(ctx, req.CandidateID)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}

	recruiterInfor, err := s.RecruiterGorm.GetProfile(ctx, req.RecruiterID)
	if err != nil {
		s.Logger.Error(err.Error())
		return err
	}

	notiCandidate := &models.Notification{
		CandidateID: req.CandidateID,
		Title:       "You followed " + recruiterInfor.Name,
		Content:     "You followed " + recruiterInfor.Name,
		Key:         "/public/recruiter/getProfile?recruiterID=" + strconv.FormatInt(recruiterInfor.RecruiterID, 10),
		CheckRead:   false,
	}

	notiRecruiter := &models.Notification{
		RecruiterID: recruiterInfor.RecruiterID,
		Title:       candidateInfor.FirstName + " is following you",
		Content:     candidateInfor.FirstName + " is following you",
		Key:         "/candidate/profile?candidateID=" + strconv.FormatInt(req.CandidateID, 10),
		CheckRead:   false,
	}
	notis := make([]*models.Notification, 0)
	notis = append(notis, notiCandidate)
	notis = append(notis, notiRecruiter)

	err = s.NotiGorm.Create(ctx, notis)
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
