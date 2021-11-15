package job

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/services/job_skill"

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
	GetAllJobForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListJobAdmin, error)
	UpdateStatusJob(ctx context.Context, updateRequest *models.RequestUpdateStatusJob) error
	DeleteJob(ctx context.Context, job_id int64) error
}

type Job struct {
	JobGorm      *JobGorm
	JobES        *JobES
	JobSkillGorm job_skill.IJobSkillDatabase

	Conf   *config.Config
	Logger *zap.Logger
}

func NewJobService(jobGorm *JobGorm, jobES *JobES, js *job_skill.JobSkillGorm, conf *config.Config, logger *zap.Logger) *Job {
	return &Job{
		JobGorm:      jobGorm,
		JobES:        jobES,
		JobSkillGorm: js,
		Conf:         conf,
		Logger:       logger,
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
	esJob.SkillIDs = job.SkillIDs
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

	for _, sid := range job.SkillIDs {
		jsModel := &models.JobSkill{
			SkillID: sid,
			JobID:   newJob.JobID,
		}
		_, err = j.JobSkillGorm.Create(ctx, jsModel)
		if err != nil {
			j.Logger.Error("Create job skill error", zap.Error(err), zap.Int64("job_id", newJob.JobID), zap.Int64("skill_id", sid))
			continue
		}
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

//job admin
func (j *Job) GetAllJobForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListJobAdmin, error) {
	acc, err := j.JobGorm.GetAllJobForAdmin(ctx, name, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (j *Job) UpdateStatusJob(ctx context.Context, updateRequest *models.RequestUpdateStatusJob) error {
	jobModels := &models.RequestUpdateStatusJob{
		Status: updateRequest.Status}
	err := j.JobGorm.UpdateStatusJob(ctx, jobModels, updateRequest.JobID)
	if err != nil {
		return err
	}
	return nil
}
func (j *Job) DeleteJob(ctx context.Context, job_id int64) error {
	err := j.JobGorm.DeleteJob(ctx, job_id)
	if err != nil {
		j.Logger.Error("Can not delete to MySQL", zap.Error(err))
		return err
	}
	return nil
}
