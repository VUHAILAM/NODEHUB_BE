package job_apply

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	jobApplyTable = "job_candidate"
	jobTable      = "job"
)

type IJobApplyDatabase interface {
	Create(ctx context.Context, jobApply *models.JobApply) (int64, error)
	GetAppliedJobByJobID(ctx context.Context, jobID int64, offset, size int64) ([]*models.Job, int64, error)
	GetAppliedJobByCandidateID(ctx context.Context, candidateID int64, offset, size int64) ([]*models.Job, int64, error)
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

func (g *JobApplyGorm) GetAppliedJobByJobID(ctx context.Context, jobID int64, offset, size int64) ([]*models.Job, int64, error) {
	db := g.DB.WithContext(ctx)
	var jobs []*models.Job
	err := db.Table(jobTable).
		Joins("INNER JOIN "+jobApplyTable+" ON "+jobTable+".job_id = "+jobApplyTable+".job_id").
		Where(jobTable+".job_id=?", jobID).
		Offset(int(offset)).Limit(int(size)).Order(jobApplyTable + ".updated_at desc").
		Find(&jobs).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: GetAppliedJobByJobID error", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = db.Table(jobApplyTable).Count(&total).Where("job_id=?", jobID).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: Count total job error")
		return jobs, total, err
	}
	return jobs, total, nil
}

func (g *JobApplyGorm) GetAppliedJobByCandidateID(ctx context.Context, candidateID int64, offset, size int64) ([]*models.Job, int64, error) {
	db := g.DB.WithContext(ctx)
	var jobs []*models.Job
	err := db.Table(jobTable).
		Joins("INNER JOIN "+jobApplyTable+" ON "+jobTable+".job_id = "+jobApplyTable+".job_id").
		Where(jobApplyTable+".candidate_id=?", candidateID).
		Offset(int(offset)).Limit(int(size)).Order(jobApplyTable + ".updated_at desc").
		Find(&jobs).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: GetAppliedJobByJobID error", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = db.Table(jobApplyTable).Count(&total).Where("candidate_id=?", candidateID).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: Count total job error")
		return jobs, total, err
	}
	return jobs, total, nil
}
