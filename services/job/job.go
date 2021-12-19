package job

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"gitlab.com/hieuxeko19991/job4e_be/services/autocomplete"

	"gitlab.com/hieuxeko19991/job4e_be/services/follow"

	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"

	"gitlab.com/hieuxeko19991/job4e_be/services/notification"

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
	CountJob(ctx context.Context) (int64, error)
}

type Job struct {
	JobGorm       *JobGorm
	JobES         *JobES
	JobSkillGorm  job_skill.IJobSkillDatabase
	SkillGorm     *skill.SkillGorm
	NotiGorm      notification.INotificationDatabase
	RecruiterGorm recruiter.IRecruiterDatabase
	FollowGorm    follow.IFollowDatabase

	jobTrie *autocomplete.Trie

	Conf   *config.Config
	Logger *zap.Logger
}

func NewJobService(jobGorm *JobGorm, jobES *JobES, js *job_skill.JobSkillGorm, skillgorm *skill.SkillGorm, notiGorm *notification.NotificationGorm, recruiterGorm *recruiter.RecruiterGorm, followGorm *follow.FollowGorm, conf *config.Config, logger *zap.Logger, jobTrie *autocomplete.Trie) *Job {
	return &Job{
		JobGorm:       jobGorm,
		JobES:         jobES,
		SkillGorm:     skillgorm,
		JobSkillGorm:  js,
		NotiGorm:      notiGorm,
		RecruiterGorm: recruiterGorm,
		FollowGorm:    followGorm,
		jobTrie:       jobTrie,

		Conf:   conf,
		Logger: logger,
	}
}

func (j *Job) CreateNewJob(ctx context.Context, job *models.CreateJobRequest) error {
	recruiterInfo, err := j.RecruiterGorm.GetProfile(ctx, job.RecruiterID)
	if err != nil {
		j.Logger.Error(err.Error())
		return err
	}
	jobData := &models.Job{
		RecruiterID: job.RecruiterID,
		CompanyName: recruiterInfo.Name,
		Avatar:      recruiterInfo.Avartar,
		Title:       job.Title,
		Description: job.Description,
		SalaryRange: job.SalaryRange,
		Quantity:    job.Quantity,
		Role:        job.Role,
		Experience:  job.Experience,
		Location:    job.Location,
		Questions:   job.Questions,
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
	err = j.JobES.Create(ctx, strconv.FormatInt(jobData.JobID, 10), jobInput)
	if err != nil {
		j.Logger.Error("Create job to ES error", zap.Error(err))
		return err
	}

	candidates, err := j.FollowGorm.GetListCandidateID(ctx, job.RecruiterID)
	if err != nil {
		j.Logger.Error(err.Error())
		return err
	}
	if len(candidates) == 0 {
		j.Logger.Info("recruiter did not have follower", zap.Int64("recruiter_id", job.RecruiterID))
		return nil
	}

	notis := make([]*models.Notification, 0)
	for _, cid := range candidates {
		noti := &models.Notification{
			CandidateID: cid.CandidateID,
			Title:       recruiterInfo.Name + " has a new job maybe fit to you. Let check it!!",
			Content:     recruiterInfo.Name + " has a new job maybe fit to you. Let check it!!",
			Key:         "/public/job/getJob?jobID=" + strconv.FormatInt(jobData.JobID, 10),
			CheckRead:   false,
		}
		notis = append(notis, noti)
	}

	err = j.NotiGorm.Create(ctx, notis)
	if err != nil {
		j.Logger.Error(err.Error())
		return err
	}

	j.jobTrie.Insert(job.Title)
	j.jobTrie.Insert(recruiterInfo.Name)
	return nil
}

func (j *Job) GetDetailJob(ctx context.Context, jobID int64) (*models.ESJob, error) {
	job, err := j.JobES.GetJobByID(ctx, strconv.FormatInt(jobID, 10))
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
		CompanyName: updateRequest.CompanyName,
		Avatar:      updateRequest.Avatar,
		Title:       updateRequest.Title,
		Description: updateRequest.Description,
		SalaryRange: updateRequest.SalaryRange,
		Quantity:    updateRequest.Quantity,
		Role:        updateRequest.Role,
		Experience:  updateRequest.Experience,
		Location:    updateRequest.Location,
		Status:      updateRequest.Status,
		Questions:   updateRequest.Questions,
		HireDate:    time.Time(updateRequest.HireDate).Format("2006-01-02"),
	}
	var err error
	if updateRequest.SkillIDs != nil && len(updateRequest.SkillIDs) != 0 {
		skillList, err := j.SkillGorm.GetSkillByIDs(ctx, updateRequest.SkillIDs)
		if err != nil {
			j.Logger.Error(err.Error())
			return err
		}
		var esSkill []models.ESSkill
		for _, skill := range skillList {
			esSk := models.ToESSkill(&skill)
			esSkill = append(esSkill, esSk)
		}
		updateES.Skills = esSkill

		err = j.JobSkillGorm.Delete(ctx, updateRequest.JobID)
		if err != nil {
			j.Logger.Error(err.Error())
			return err
		}
		var jobSkill []*models.JobSkill
		for _, skillID := range updateRequest.SkillIDs {
			jsk := &models.JobSkill{
				SkillID: skillID,
				JobID:   updateRequest.JobID,
			}
			jobSkill = append(jobSkill, jsk)
		}
		err = j.JobSkillGorm.Create(ctx, jobSkill)
		if err != nil {
			j.Logger.Error(err.Error())
			return err
		}
	}
	updateData := map[string]interface{}{}
	err = mapStructureDecodeWithTextUnmarshaler(updateES, &updateData)
	if err != nil {
		j.Logger.Error("Can not convert to map", zap.Error(err))
		return err
	}
	j.Logger.Info("updateData", zap.Reflect("DATA", updateData))
	err = j.JobES.Update(ctx, strconv.FormatInt(updateRequest.JobID, 10), updateData)
	if err != nil {
		j.Logger.Error("Can not Update to ES", zap.Error(err))
		return err
	}
	_, ok := updateData["skills"]
	if ok {
		delete(updateData, "skills")
	}
	_, ok = updateData["questions"]
	if ok {
		questions, err := json.Marshal(updateES.Questions)
		if err != nil {
			j.Logger.Error("Marshal question error", zap.Error(err))
			delete(updateData, "questions")
		} else {
			updateData["questions"] = string(questions)
		}
	}
	_, ok = updateData["company_name"]
	if ok {
		delete(updateData, "company_name")
	}
	_, ok = updateData["avatar"]
	if ok {
		delete(updateData, "avatar")
	}
	err = j.JobGorm.Update(ctx, updateRequest.JobID, updateData)
	if err != nil {
		j.Logger.Error("Can not Update to MySQL", zap.Error(err))
		return err
	}

	j.jobTrie.Insert(updateRequest.Title)
	j.jobTrie.Insert(updateRequest.CompanyName)
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
	err = j.JobES.Update(ctx, strconv.FormatInt(updateRequest.JobID, 10), updateData)
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
	err = j.JobES.Delete(ctx, strconv.FormatInt(job_id, 10))
	if err != nil {
		return err
	}
	return nil
}

func (j *Job) SearchJob(ctx context.Context, searchReq models.RequestSearchJob) (*models.ResponseGetJob, error) {
	offset := (searchReq.Page - 1) * searchReq.Size
	jobs := make([]models.ESJob, 0)
	var total int64
	var err error
	if searchReq.Text == "" && searchReq.Location == "" {
		jobs, total, err = j.JobES.GetAllJob(ctx, offset, searchReq.Size)
	} else {
		jobs, total, err = j.JobES.SearchJobs(ctx, searchReq.Text, searchReq.Location, offset, searchReq.Size)
	}
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

func (j *Job) CountJob(ctx context.Context) (int64, error) {
	return j.JobGorm.Count(ctx)
}
