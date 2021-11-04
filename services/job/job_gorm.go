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
	Create(ctx context.Context, job *models.Job) (*models.Job, error)
	Get(ctx context.Context, jobID int64) (*models.Job, error)
	Update(ctx context.Context, jobID int64, data map[string]interface{}) error
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

func (g *JobGorm) Create(ctx context.Context, job *models.Job) (*models.Job, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobTable).Create(job).Error
	if err != nil {
		g.Logger.Error("JobGorm: Create job error", zap.Error(err))
		return nil, err
	}
	return job, nil
}

func (g *JobGorm) Get(ctx context.Context, jobID int64) (*models.Job, error) {
	db := g.DB.WithContext(ctx)
	job := models.Job{}
	err := db.Table(jobTable).Where("job_id=?", jobID).First(&job).Error
	if err != nil {
		g.Logger.Error("JobGorm: Get job error", zap.Error(err), zap.Int64("job_id", jobID))
		return nil, err
	}
	return &job, nil
}

func (g *JobGorm) Update(ctx context.Context, jobID int64, data map[string]interface{}) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobTable).Where("job_id=?", jobID).Updates(data).Error
	if err != nil {
		g.Logger.Error("JobGorm: Update job error", zap.Error(err), zap.Int64("job_id", jobID))
		return err
	}
	return nil
}
