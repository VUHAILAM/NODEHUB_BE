package job_apply

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type IJobApplyService interface {
	CreateJobApply(ctx context.Context, req models.RequestApply) error
}

type JobApply struct {
	JobApplyGorm *JobApplyGorm
	Logger       *zap.Logger
}

func NewJobApplyService(gorm *JobApplyGorm, logger *zap.Logger) *JobApply {
	return &JobApply{
		JobApplyGorm: gorm,
		Logger:       logger,
	}
}

func (ja *JobApply) CreateJobApply(ctx context.Context, req models.RequestApply) error {
	jobApply := models.JobApply{
		JobID:       req.JobID,
		CandidateID: req.CandidateID,
		Status:      req.Status,
	}
	_, err := ja.JobApplyGorm.Create(ctx, &jobApply)
	if err != nil {
		ja.Logger.Error("Can not create apply candidate", zap.Error(err), zap.Reflect("request", req))
		return err
	}
	return nil
}
