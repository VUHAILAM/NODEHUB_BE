package job_skill

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type IJobSkillService interface {
	GetJobBySkill(ctx context.Context, req models.RequestGetJobsBySkill) (*models.ResponseGetJobsBySkill, error)
	GetSkillsByJob(ctx context.Context, req models.RequestGetSkillsByJob) ([]*models.Skill, error)
}

type JobSkill struct {
	JSGorm IJobSkillDatabase
	Logger *zap.Logger
}

func NewJobSkill(gorm *JobSkillGorm, logger *zap.Logger) *JobSkill {
	return &JobSkill{
		JSGorm: gorm,
		Logger: logger,
	}
}

func (js *JobSkill) GetJobBySkill(ctx context.Context, req models.RequestGetJobsBySkill) (*models.ResponseGetJobsBySkill, error) {
	offset := (req.Page - 1) * req.Size
	jobs, total, err := js.JSGorm.GetAllJobBySkill(ctx, req.SkillID, offset, req.Size)
	if err != nil {
		js.Logger.Error("Can not get jobs", zap.Error(err), zap.Int64("skill_id", req.SkillID))
		return nil, err
	}
	resp := models.ResponseGetJobsBySkill{
		Total:  total,
		Result: jobs,
	}
	return &resp, nil
}

func (js *JobSkill) GetSkillsByJob(ctx context.Context, req models.RequestGetSkillsByJob) ([]*models.Skill, error) {
	skills, err := js.JSGorm.GetAllSkillByJob(ctx, req.JobID)
	if err != nil {
		js.Logger.Error("Can not get skill", zap.Error(err), zap.Int64("job_id", req.JobID))
		return nil, err
	}
	return skills, nil
}
