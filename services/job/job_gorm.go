package job

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	jobTable = "job"
)

type IJobDatabase interface {
	Create(ctx context.Context, job *models.Job) (int64, error)
}

type JobGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewJobGorm(db *gorm.DB, logger *zap.Logger) *JobGorm {
	return &JobGorm{
		DB:     db,
		Logger: logger,
	}
}

func (g *JobGorm) Create(ctx context.Context, job *models.Job) (int64, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobTable).Create(job).Error
	if err != nil {
		g.Logger.Error("JobGorm: Create job error", zap.Error(err))
		return 0, err
	}
	return job.JobID, nil
}
