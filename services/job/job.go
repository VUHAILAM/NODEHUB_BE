package job

import (
	"context"

	"github.com/mitchellh/mapstructure"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IJobService interface {
	CreateNewJob(ctx context.Context, job *models.CreateJobRequest) error
	GetDetailJob(ctx context.Context, jobID int64) (*models.Job, error)
	UpdateJob(ctx context.Context, updateRequest *models.RequestUpdateJob) error
	GetAllJob(ctx context.Context, getRequest *models.RequestGetAllJob) (*models.ResponseGetAllJob, error)
}

type Job struct {
	JobGorm *JobGorm
	JobES   *JobES
	Conf    *config.Config
	Logger  *zap.Logger
}

func NewJobService(jobGorm *JobGorm, jobES *JobES, conf *config.Config, logger *zap.Logger) *Job {
	return &Job{
		JobGorm: jobGorm,
		JobES:   jobES,
		Conf:    conf,
		Logger:  logger,
	}
}

func (j *Job) CreateNewJob(ctx context.Context, job *models.CreateJobRequest) error {
	jobData := &models.Job{
		RecruiterID: job.RecruiterID,
		Title:       job.Title,
		Description: job.Description,
		SalaryRange: job.SalaryRange,
		Quantity:    job.Quantity,
		Role:        job.Role,
		Experience:  job.Experience,
		Location:    job.Location,
		HireDate:    job.HireDate,
		Status:      job.Status,
	}
	newJob, err := j.JobGorm.Create(ctx, jobData)
	if err != nil {
		j.Logger.Error("Create job to SQL error", zap.Error(err))
		return err
	}

	jobData.JobID = newJob.JobID
	jobData.CreatedAt = newJob.CreatedAt
	esJob := models.ToESJobCreate(jobData)
	jobInput := map[string]interface{}{}
	err = mapStructureDecodeWithTextUnmarshaler(esJob, &jobInput)
	if err != nil {
		j.Logger.Error("Cannot convert map to Job log struct", zap.Error(err))
		return err
	}
	j.Logger.Info("Data input", zap.Reflect("Input", jobInput))
	err = j.JobES.Create(ctx, string(jobData.JobID), jobInput)
	if err != nil {
		j.Logger.Error("Create job to ES error", zap.Error(err))
		return err
	}
	return nil
}

func (j *Job) GetDetailJob(ctx context.Context, jobID int64) (*models.Job, error) {
	job, err := j.JobES.GetJobByID(ctx, string(jobID))
	if err != nil {
		j.Logger.Error("Can not get Job from ES", zap.Error(err), zap.Int64("job_id", jobID))
		job, err = j.JobGorm.Get(ctx, jobID)
		if err != nil {
			j.Logger.Error("Can not Get Job", zap.Error(err), zap.Int64("job_id", jobID))
			return nil, err
		}
		return job, nil
	}
	return job, nil
}

func (j *Job) UpdateJob(ctx context.Context, updateRequest *models.RequestUpdateJob) error {
	j.Logger.Info("updateReq", zap.Reflect("Req", updateRequest))
	updateData := map[string]interface{}{}
	err := mapStructureDecodeWithTextUnmarshaler(updateRequest, &updateData)
	if err != nil {
		j.Logger.Error("Can not convert to map", zap.Error(err))
		return err
	}
	j.Logger.Info("updateData", zap.Reflect("DATA", updateData))
	err = j.JobES.Update(ctx, string(updateRequest.JobID), updateData)
	if err != nil {
		j.Logger.Error("Can not Update to ES", zap.Error(err))
		return err
	}

	err = j.JobGorm.Update(ctx, updateRequest.JobID, updateData)
	if err != nil {
		j.Logger.Error("Can not Update to MySQL", zap.Error(err))
		return err
	}
	return nil
}

func (j *Job) GetAllJob(ctx context.Context, getRequest *models.RequestGetAllJob) (*models.ResponseGetAllJob, error) {
	offset := (getRequest.Page - 1) * getRequest.Size
	jobs, total, err := j.JobES.GetAllJob(ctx, offset, getRequest.Size)
	if err != nil {
		j.Logger.Error("Can not Get all Job from ES", zap.Error(err))
		return nil, err
	}
	resp := models.ResponseGetAllJob{
		Total:  total,
		Result: jobs,
	}
	return &resp, nil
}

func mapStructureDecodeWithTextUnmarshaler(input, output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:     output,
		DecodeHook: mapstructure.TextUnmarshallerHookFunc(),
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
