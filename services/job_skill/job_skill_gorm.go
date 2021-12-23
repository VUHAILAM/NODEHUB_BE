package job_skill

import (
	"context"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	jobSkillTable = "job_skill"
	jobTable      = "job"
	skillTable    = "skill"
)

type IJobSkillDatabase interface {
	Create(ctx context.Context, jobSkill []*models2.JobSkill) error
	GetAllSkillByJob(ctx context.Context, jobID int64) ([]*models2.Skill, error)
	GetAllJobBySkill(ctx context.Context, skillID, offset, size int64) ([]*models2.Job, int64, error)
	Delete(ctx context.Context, jobID int64) error
}

type JobSkillGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewJobSkillGorm(db *gorm.DB, logger *zap.Logger) *JobSkillGorm {
	return &JobSkillGorm{
		DB:     db,
		Logger: logger,
	}
}

func (g *JobSkillGorm) Create(ctx context.Context, jobSkill []*models2.JobSkill) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobSkillTable).Create(&jobSkill).Error
	if err != nil {
		g.Logger.Error("JobSkillGorm: Create job skill error", zap.Error(err))
		return err
	}
	return nil
}

func (g *JobSkillGorm) Delete(ctx context.Context, jobID int64) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobSkillTable).Where("job_id=?", jobID).Delete(&models2.JobSkill{}).Error
	if err != nil {
		g.Logger.Error("JobSkillGorm: Delete job skill error", zap.Error(err))
		return err
	}
	return nil
}

func (g *JobSkillGorm) GetAllSkillByJob(ctx context.Context, jobID int64) ([]*models2.Skill, error) {
	db := g.DB.WithContext(ctx)
	var skills []*models2.Skill
	err := db.Table(skillTable).
		Joins("INNER JOIN "+jobSkillTable+" ON "+skillTable+".skill_id = "+jobSkillTable+".job_id").
		Where(jobSkillTable+".job_id=?", jobID).
		Find(&skills).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: GetAppliedJobByJobID error", zap.Error(err))
		return nil, err
	}
	return skills, nil
}

func (g *JobSkillGorm) GetAllJobBySkill(ctx context.Context, skillID, offset, size int64) ([]*models2.Job, int64, error) {
	db := g.DB.WithContext(ctx)
	var jobs []*models2.Job
	err := db.Table(jobTable).
		Joins("JOIN "+jobSkillTable+" ON "+jobTable+".job_id = "+jobSkillTable+".job_id").
		Where(jobSkillTable+".skill_id=?", skillID).
		Offset(int(offset)).Limit(int(size)).Order(jobSkillTable + ".updated_at desc").
		Find(&jobs).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: GetAppliedJobByJobID error", zap.Error(err))
		return nil, 0, err
	}
	var total int64
	err = db.Table(jobSkillTable).Count(&total).Where("skill_id=?", skillID).Error
	if err != nil {
		g.Logger.Error("JobApplyGorm: Count total job error")
		return jobs, total, err
	}
	return jobs, total, nil
}
