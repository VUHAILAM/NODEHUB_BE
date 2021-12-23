package job_apply

import (
	"context"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	jobApplyTable  = "job_candidate"
	jobTable       = "job"
	candidateTable = "candidate"
	recruiterTable = "recruiter"
)

type IJobApplyDatabase interface {
	Create(ctx context.Context, jobApply *models2.JobApply) (int64, error)
	GetAppliedJobByJobID(ctx context.Context, jobID int64, offset, size int64) ([]*models2.Job, int64, error)
	GetAppliedJobByCandidateID(ctx context.Context, candidateID int64, offset, size int64) ([]*models2.Job, int64, error)
	UpdateStatus(ctx context.Context, status string, jobID, candidateID int64) error
	GetCandidateApplyJob(ctx context.Context, jobID int64, offset, size int64) ([]*models2.Candidate, int64, error)
	CountByStatus(ctx context.Context, status string) (int64, error)
	CheckApplied(ctx context.Context, jobID, candidateID int64) (*models2.JobApply, error)
	GetApply(ctx context.Context, jobID, candidateID int64) (*models2.JobApply, error)
	CountApplyOnMonth(ctx context.Context, date string) (int64, error)
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

func (g *JobApplyGorm) Create(ctx context.Context, jobApply *models2.JobApply) (int64, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobApplyTable).Create(jobApply).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: Create job apply error", zap.Error(err))
		return 0, err
	}
	return jobApply.ID, nil
}

func (g *JobApplyGorm) GetAppliedJobByJobID(ctx context.Context, jobID int64, offset, size int64) ([]*models2.Job, int64, error) {
	db := g.DB.WithContext(ctx)
	var jobs []*models2.Job
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

func (g *JobApplyGorm) GetAppliedJobByCandidateID(ctx context.Context, candidateID int64, offset, size int64) ([]*models2.Job, int64, error) {
	var jobs []*models2.Job
	db := g.DB.WithContext(ctx).Table(jobTable).Select("job.*, recruiter.name as company_name, recruiter.avartar as avartar, job_candidate.status as candidate_status").
		Joins("JOIN "+recruiterTable+" ON job.recruiter_id=recruiter.recruiter_id").Joins("JOIN "+jobApplyTable+" ON job.job_id=job_candidate.job_id").
		Where("job_candidate.candidate_id=?", candidateID).Find(&jobs)
	total := db.RowsAffected
	jobs = make([]*models2.Job, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Order(jobApplyTable + ".updated_at desc").Find(&jobs).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: GetAppliedJobByJobID error", zap.Error(err))
		return nil, 0, err
	}

	return jobs, total, nil
}

func (g *JobApplyGorm) UpdateStatus(ctx context.Context, status string, jobID, candidateID int64) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobApplyTable).Where("job_id=? and candidate_id=?", jobID, candidateID).Update("status", status).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (g *JobApplyGorm) GetCandidateApplyJob(ctx context.Context, jobID int64, offset, size int64) ([]*models2.Candidate, int64, error) {
	var candidates []*models2.Candidate
	db := g.DB.WithContext(ctx).Table(candidateTable).Select("candidate.*, job_candidate.status as job_status").
		Joins("JOIN "+jobApplyTable+" ON candidate.candidate_id=job_candidate.candidate_id").
		Where("job_candidate.job_id=?", jobID).Find(&candidates)
	total := db.RowsAffected
	candidates = make([]*models2.Candidate, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Order(jobApplyTable + ".updated_at desc").Find(&candidates).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: GetAppliedJobByJobID error", zap.Error(err))
		return nil, 0, err
	}

	return candidates, total, nil
}

func (g *JobApplyGorm) CountByStatus(ctx context.Context, status string) (int64, error) {
	var count int64
	db := g.DB.WithContext(ctx).Table(jobApplyTable).Select("COUNT(candidate_id) as count").
		Where("status=?", status).Count(&count)
	err := db.Error
	if err != nil {
		g.Logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}

func (g *JobApplyGorm) CheckApplied(ctx context.Context, jobID, candidateID int64) (*models2.JobApply, error) {
	db := g.DB.WithContext(ctx)
	jobapply := models2.JobApply{}
	err := db.Table(jobApplyTable).Where("job_id=? and candidate_id=?", jobID, candidateID).Take(&jobapply).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, err
	}
	return &jobapply, nil
}

func (g *JobApplyGorm) GetApply(ctx context.Context, jobID, candidateID int64) (*models2.JobApply, error) {
	db := g.DB.WithContext(ctx)
	jobapply := models2.JobApply{}
	err := db.Table(jobApplyTable).Select("job_candidate.*, job.questions as questions").
		Joins("JOIN "+jobTable+" ON job.job_id=job_candidate.job_id").
		Where("job_candidate.job_id=? and job_candidate.candidate_id=?", jobID, candidateID).
		Take(&jobapply).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, err
	}
	return &jobapply, nil
}

func (g *JobApplyGorm) CountApplyOnMonth(ctx context.Context, date string) (int64, error) {
	db := g.DB.WithContext(ctx)
	var count int64
	err := db.Table(jobApplyTable).Where("created_at>=(LAST_DAY(?) + INTERVAL 1 DAY - INTERVAL 1 MONTH) and created_at<(LAST_DAY(?) + INTERVAL 1 DAY)", date, date).Count(&count).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}
