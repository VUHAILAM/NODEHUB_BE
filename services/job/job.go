package job

import (
	"context"
	"time"

	"gitlab.com/hieuxeko19991/job4e_be/services/skill"

	"gitlab.com/hieuxeko19991/job4e_be/services/job_skill"

	"github.com/mitchellh/mapstructure"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IJobService interface {
	CreateNewJob(ctx context.Context, job *models.CreateJobRequest) error
	GetDetailJob(ctx context.Context, jobID int64) (*models.ESJob, error)
	UpdateJob(ctx context.Context, updateRequest *models.RequestUpdateJob) error
	GetAllJob(ctx context.Context, getRequest *models.RequestGetAllJob) (*models.ResponseGetJob, error)
	GetJobsByRecruiterID(ctx context.Context, req *models.RequestGetJobsByRecruiter) (*models.ResponseGetJob, error)
	GetAllJobForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListJobAdmin, error)
	UpdateStatusJob(ctx context.Context, updateRequest *models.RequestUpdateStatusJob) error
	DeleteJob(ctx context.Context, job_id int64) error
	SearchJob(ctx context.Context, searchReq models.RequestSearchJob) (*models.ResponseGetJob, error)
}

type Job struct {
	JobGorm      *JobGorm
	JobES        *JobES
	JobSkillGorm job_skill.IJobSkillDatabase
	SkillGorm    *skill.SkillGorm

	Conf   *config.Config
	Logger *zap.Logger
}

func NewJobService(jobGorm *JobGorm, jobES *JobES, js *job_skill.JobSkillGorm, skillgorm *skill.SkillGorm, conf *config.Config, logger *zap.Logger) *Job {
	return &Job{
		JobGorm:      jobGorm,
		JobES:        jobES,
		SkillGorm:    skillgorm,
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
		HireDate:    time.Time(job.HireDate),
		Status:      job.Status,
	}
	newJob, err := j.JobGorm.Create(ctx, jobData)
	if err != nil {
		j.Logger.Error("Create job to SQL error", zap.Error(err))
		return err
	}

	var jobSkill []*models.JobSkill
	for _, skillID := range job.SkillIDs {
		jsk := &models.JobSkill{
			SkillID: skillID,
			JobID:   newJob.JobID,
		}
		jobSkill = append(jobSkill, jsk)
	}

	err = j.JobSkillGorm.Create(ctx, jobSkill)
	if err != nil {
		j.Logger.Error("Create job skill error", zap.Error(err), zap.Int64("job_id", newJob.JobID))
		return err
	}

	jobData.JobID = newJob.JobID
	jobData.CreatedAt = newJob.CreatedAt
	esJob := models.ToESJobCreate(jobData)
	skillList, err := j.SkillGorm.GetSkillByIDs(ctx, job.SkillIDs)
	if err != nil {
		j.Logger.Error(err.Error())
		return err
	}
	var esSkill []models.ESSkill
	for _, skill := range skillList {
		esSk := models.ToESSkill(&skill)
		esSkill = append(esSkill, esSk)
	}
	esJob.Skills = esSkill
	j.Logger.Info("EsJob", zap.Reflect("es_job", esJob))
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

func (j *Job) GetDetailJob(ctx context.Context, jobID int64) (*models.ESJob, error) {
	job, err := j.JobES.GetJobByID(ctx, string(jobID))
	if err != nil {
		j.Logger.Error("Can not get Job from ES", zap.Error(err), zap.Int64("job_id", jobID))
		return nil, err
	}
	return job, nil
}

func (j *Job) UpdateJob(ctx context.Context, updateRequest *models.RequestUpdateJob) error {
	j.Logger.Info("updateReq", zap.Reflect("Req", updateRequest))
	updateES := models.ESJobUpdate{
		JobID:       updateRequest.JobID,
		RecruiterID: updateRequest.RecruiterID,
		Title:       updateRequest.Title,
		Description: updateRequest.Description,
		SalaryRange: updateRequest.SalaryRange,
		Quantity:    updateRequest.Quantity,
		Role:        updateRequest.Role,
		Experience:  updateRequest.Experience,
		Location:    updateRequest.Location,
		Status:      updateRequest.Status,
		HireDate:    time.Time(updateRequest.HireDate).Format("2006-01-02"),
	}
	updateData := map[string]interface{}{}
	err := mapStructureDecodeWithTextUnmarshaler(updateES, &updateData)
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

func (j *Job) GetAllJob(ctx context.Context, getRequest *models.RequestGetAllJob) (*models.ResponseGetJob, error) {
	offset := (getRequest.Page - 1) * getRequest.Size
	jobs, total, err := j.JobES.GetAllJob(ctx, offset, getRequest.Size)
	if err != nil {
		j.Logger.Error("Can not Get all Job from ES", zap.Error(err))
		return nil, err
	}
	resp := models.ResponseGetJob{
		Total:  total,
		Result: jobs,
	}
	return &resp, nil
}

func (j *Job) GetJobsByRecruiterID(ctx context.Context, req *models.RequestGetJobsByRecruiter) (*models.ResponseGetJob, error) {
	offset := (req.Page - 1) * req.Size
	jobs, total, err := j.JobES.GetJobsByRecruiterID(ctx, req.RecruiterID, offset, req.Size)
	if err != nil {
		j.Logger.Error("Can not Get all Job by recruiter from ES", zap.Error(err))
		return nil, err
	}
	resp := models.ResponseGetJob{
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
	updateES := models.ESJobUpdate{
		JobID:  updateRequest.JobID,
		Status: int(updateRequest.Status),
	}
	err := j.JobGorm.UpdateStatusJob(ctx, jobModels, updateRequest.JobID)
	if err != nil {
		return err
	}
	updateData := map[string]interface{}{}
	err = mapStructureDecodeWithTextUnmarshaler(updateES, &updateData)
	if err != nil {
		j.Logger.Error("Can not convert to map", zap.Error(err))
		return err
	}
	err = j.JobES.Update(ctx, string(updateRequest.JobID), updateData)
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
	err = j.JobES.Delete(ctx, string(job_id))
	if err != nil {
		return err
	}
	return nil
}

func (j *Job) SearchJob(ctx context.Context, searchReq models.RequestSearchJob) (*models.ResponseGetJob, error) {
	offset := (searchReq.Page - 1) * searchReq.Size
	jobs, total, err := j.JobES.SearchJobs(ctx, searchReq.Text, searchReq.Location, offset, searchReq.Size)
	if err != nil {
		j.Logger.Error("Can not Searhc Job from ES", zap.Error(err))
		return nil, err
	}
	resp := models.ResponseGetJob{
		Total:  total,
		Result: jobs,
	}
	return &resp, nil
}
