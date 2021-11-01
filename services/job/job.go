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
		Expereience: job.Expereience,
		Location:    job.Location,
		HireDate:    job.HireDate,
		Status:      job.Status,
	}
	jobID, err := j.JobGorm.Create(ctx, jobData)
	if err != nil {
		j.Logger.Error("Create job to SQL error", zap.Error(err))
		return err
	}

	jobData.JobID = jobID
	esJob := models.ToESJobCreate(jobData)
	jobInput := map[string]interface{}{}
	err = mapStructureDecodeWithTextUnmarshaler(esJob, jobInput)
	if err != nil {
		j.Logger.Error("Cannot convert map to Job log struct", zap.Error(err))
		return err
	}
	err = j.JobES.Create(ctx, string(jobID), jobInput)
	if err != nil {
		j.Logger.Error("Create job to ES error", zap.Error(err))
		return err
	}
	return nil
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
