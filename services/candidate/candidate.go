package candidate

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type ICandidateService interface {
	CreateCandidateProfile(ctx context.Context, req models.CandidateRequest) (int64, error)
	UpdateCandidateProfile(ctx context.Context, req models.CandidateRequest) error
	GetCandidateProfile(ctx context.Context, candidateID int64) (*models.CandidateResponse, error)
	GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListCandidateAdmin, error)
	UpdateReviewCandidateByAdmin(ctx context.Context, updateRequest *models.RequestUpdateReviewCandidateAdmin) error
	UpdateStatusCandidate(ctx context.Context, candidate *models.RequestUpdateStatusCandidate, candidate_id int64) error
	AddCandidateSkill(ctx context.Context, candidateSkill *models.CandidateSkill) error
	DeleteCandidateSkill(ctx context.Context, candidate_skill_id int64) error
	UpdateCandidateSkill(ctx context.Context, updateRequest *models.RequestUpdateCandidateSkill) error
	GetCandidateSkill(ctx context.Context, candidate_id int64) ([]models.ResponseCandidateSkill, error)
	SearchCandidate(ctx context.Context, req models.RequestSearchCandidate) (*models.ResponseSearchCandidate, error)
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

func (s *CandidateService) GetCandidateProfile(ctx context.Context, candidateID int64) (*models.CandidateResponse, error) {
	fmt.Println("lõi cai deo gì day: ", candidateID)
	candidate, err := s.CanGorm.GetByCandidateID(ctx, candidateID)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	resp, err := candidate.ToCandidateResponse()
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}
	return &resp, nil
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

// candidateSkill
func (s *CandidateService) AddCandidateSkill(ctx context.Context, candidateSkill *models.CandidateSkill) error {
	CandidateSkillModels := &models.CandidateSkill{
		Id:          candidateSkill.Id,
		CandidateId: candidateSkill.CandidateId,
		SkillId:     candidateSkill.SkillId,
		Media:       candidateSkill.Media}
	err := s.CanGorm.AddCandidateSkill(ctx, CandidateSkillModels)
	if err != nil {
		return err
	}
	return nil
}

func (s *CandidateService) DeleteCandidateSkill(ctx context.Context, candidate_skill_id int64) error {

	err := s.CanGorm.DeleteCandidateSkill(ctx, candidate_skill_id)
	if err != nil {
		s.Logger.Error("Can not delete to MySQL", zap.Error(err))
		return err
	}
	return nil
}

func (s *CandidateService) UpdateCandidateSkill(ctx context.Context, updateRequest *models.RequestUpdateCandidateSkill) error {
	updateData := map[string]interface{}{}
	err1 := mapStructureDecodeWithTextUnmarshaler(updateRequest, &updateData)
	if err1 != nil {
		s.Logger.Error("Can not convert to map", zap.Error(err1))
		return err1
	}

	err := s.CanGorm.UpdateCandidateSkill(ctx, updateRequest.ID, updateData)
	if err != nil {
		s.Logger.Error("Can not Update to MySQL", zap.Error(err))
		return err
	}
	return nil
}

func (s *CandidateService) GetCandidateSkill(ctx context.Context, candidate_id int64) ([]models.ResponseCandidateSkill, error) {
	acc, err := s.CanGorm.GetCandidateSkill(ctx, candidate_id)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *CandidateService) SearchCandidate(ctx context.Context, req models.RequestSearchCandidate) (*models.ResponseSearchCandidate, error) {
	offset := (req.Page - 1) * req.Size
	candidates, total, err := s.CanGorm.SearchCandidate(ctx, req.Text, 0, offset, req.Size)
	if err != nil {
		s.Logger.Error("Search candidate error", zap.Error(err))
		return nil, err
	}

	resp := models.ResponseSearchCandidate{
		Total:      total,
		Candidates: candidates,
	}

	return &resp, nil
}
