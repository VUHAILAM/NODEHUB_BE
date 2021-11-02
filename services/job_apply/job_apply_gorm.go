package job_apply

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	jobApplyTable = "job_candidate"
)

type IJobApplyDatabase interface {
	Create(ctx context.Context, jobApply *models.JobApply) (int64, error)
}

type JobApplyGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewJobApplyGorm(db *gorm.DB, logger *zap.Logger) *JobApplyGorm {
	return &JobApplyGorm{
		DB:     db,
		Logger: logger,
	}
}

func (g *JobApplyGorm) Create(ctx context.Context, jobApply *models.JobApply) (int64, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobApplyTable).Create(jobApply).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: Create job apply error", zap.Error(err))
		return 0, err
	}
	return jobApply.ID, nil
}
