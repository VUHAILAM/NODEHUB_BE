package candidate

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type ICandidateService interface {
	CreateCandidateProfile(ctx context.Context, req models.CandidateRequest) (int64, error)
	UpdateCandidateProfile(ctx context.Context, req models.CandidateRequest) error
	GetCandidateProfile(ctx context.Context, candidateID int64) (*models.CandidateRequest, error)
	GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListCandidateAdmin, error)
	UpdateReviewCandidateByAdmin(ctx context.Context, updateRequest *models.RequestUpdateReviewCandidateAdmin) error
	UpdateStatusCandidate(ctx context.Context, candidate *models.RequestUpdateStatusCandidate, candidate_id int64) error
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

//candidate admin
func (s *CandidateService) GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListCandidateAdmin, error) {
	acc, err := s.CanGorm.GetAllCandidateForAdmin(ctx, name, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *CandidateService) UpdateReviewCandidateByAdmin(ctx context.Context, updateRequest *models.RequestUpdateReviewCandidateAdmin) error {
	updateData := map[string]interface{}{}
	err1 := mapStructureDecodeWithTextUnmarshaler(updateRequest, &updateData)
	if err1 != nil {
		s.Logger.Error("Can not convert to map", zap.Error(err1))
		return err1
	}

	err := s.CanGorm.UpdateReviewCandidateByAdmin(ctx, updateRequest.CandidateID, updateData)
	if err != nil {
		s.Logger.Error("Can not Update to MySQL", zap.Error(err))
		return err
	}
	return nil
}

func mapStructureDecodeWithTextUnmarshaler(input, output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     output,
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(),
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

func (s *CandidateService) UpdateStatusCandidate(ctx context.Context, candidate *models.RequestUpdateStatusCandidate, candidate_id int64) error {
	candidateModels := &models.RequestUpdateStatusCandidate{
		Status: candidate.Status}
	err := s.CanGorm.UpdateStatusCandidate(ctx, candidateModels, candidate_id)
	if err != nil {
		return err
	}
	return nil
}
