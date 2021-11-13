package candidate

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	candidateTable = "candidate"
)

type ICandidateDatabase interface {
	Create(ctx context.Context, createData *models.Candidate) (int64, error)
	Update(ctx context.Context, candidateID int64, updateData *models.Candidate) error
	GetByCandidateID(ctx context.Context, candidateID int64) (*models.Candidate, error)
}

type CandidateGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewCandidateGorm(db *gorm.DB, logger *zap.Logger) *CandidateGorm {
	return &CandidateGorm{
		DB:     db,
		Logger: logger,
	}
}

func (g *CandidateGorm) Create(ctx context.Context, createData *models.Candidate) (int64, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(candidateTable).Create(createData).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return 0, err
	}
	return createData.CandidateID, nil
}

func (g *CandidateGorm) Update(ctx context.Context, candidateID int64, updateData *models.Candidate) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(candidateTable).Where("candidate_id=?", candidateID).Updates(updateData).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (g *CandidateGorm) GetByCandidateID(ctx context.Context, candidateID int64) (*models.Candidate, error) {
	db := g.DB.WithContext(ctx)
	candidate := models.Candidate{}
	err := db.Table(candidateTable).Where("candidate_id=?", candidateID).Take(&candidate).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, err
	}
	return &candidate, nil
}
