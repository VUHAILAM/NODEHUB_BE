package job_apply

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"
)

type IJobApplyService interface {
	CreateJobApply(ctx context.Context, req models.RequestApply) error
	GetJobsByJobID(ctx context.Context, req models.RequestGetJobApplyByJobID) (*models.ResponseGetJobApply, error)
	GetJobByCandidateID(ctx context.Context, req models.RequestGetJobApplyByCandidateID) (*models.ResponseGetJobApply, error)
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

func (ja *JobApply) GetJobsByJobID(ctx context.Context, req models.RequestGetJobApplyByJobID) (*models.ResponseGetJobApply, error) {
	offset := (req.Page - 1) * req.Size
	jobs, total, err := ja.JobApplyGorm.GetAppliedJobByJobID(ctx, req.JobID, offset, req.Size)
	if err != nil {
		ja.Logger.Error("Can not get jobs", zap.Error(err), zap.Int64("job_id", req.JobID))
		return nil, err
	}
	resp := models.ResponseGetJobApply{
		Total:  total,
		Result: jobs,
	}
	return &resp, nil
}

func (ja *JobApply) GetJobByCandidateID(ctx context.Context, req models.RequestGetJobApplyByCandidateID) (*models.ResponseGetJobApply, error) {
	offset := (req.Page - 1) * req.Size
	jobs, total, err := ja.JobApplyGorm.GetAppliedJobByCandidateID(ctx, req.CandidateID, offset, req.Size)
	if err != nil {
		ja.Logger.Error("Can not get jobs", zap.Error(err), zap.Int64("candidate_id", req.CandidateID))
		return nil, err
	}
	resp := models.ResponseGetJobApply{
		Total:  total,
		Result: jobs,
	}
	return &resp, nil
}
