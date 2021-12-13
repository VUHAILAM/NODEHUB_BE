package job_apply

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/services/job_skill"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/job"
	"gitlab.com/hieuxeko19991/job4e_be/services/notification"
	"go.uber.org/zap"
)

type IJobApplyService interface {
	CreateJobApply(ctx context.Context, req models.RequestApply) error
	GetJobByCandidateID(ctx context.Context, req models.RequestGetJobApplyByCandidateID) (*models.ResponseGetJobApply, error)
	GetCandidatesApplyJob(ctx context.Context, req models.RequestGetJobApplyByJobID) (*models.ResponseGetCandidateApply, error)
	UpdateStatusJobApplied(ctx context.Context, req models.RequestUpdateStatusJobApplied) error
	CountCandidateByStatus(ctx context.Context, req models.RequestCountStatus) (int64, error)
	CheckApplied(ctx context.Context, req models.RequestCheckApply) (*models.JobApply, error)
	GetApply(ctx context.Context, req models.RequestCheckApply) (*models.JobApply, error)
}

type JobApply struct {
	JobApplyGorm *JobApplyGorm
	NotiGorm     notification.INotificationDatabase
	JobGorm      job.IJobDatabase
	JobSkill     job_skill.IJobSkillDatabase

	Logger *zap.Logger
}

func NewJobApplyService(gorm *JobApplyGorm, jobGorm *job.JobGorm, notiGorm *notification.NotificationGorm, jobSkill *job_skill.JobSkillGorm, logger *zap.Logger) *JobApply {
	return &JobApply{
		JobApplyGorm: gorm,
		NotiGorm:     notiGorm,
		JobGorm:      jobGorm,
		JobSkill:     jobSkill,
		Logger:       logger,
	}
}

func (ja *JobApply) CreateJobApply(ctx context.Context, req models.RequestApply) error {
	jobApply := models.JobApply{
		JobID:       req.JobID,
		CandidateID: req.CandidateID,
		Status:      req.Status,
		Answers:     req.Answers,
		Media:       req.Media,
	}

	_, err := ja.JobApplyGorm.Create(ctx, &jobApply)
	if err != nil {
		ja.Logger.Error("Can not create apply candidate", zap.Error(err), zap.Reflect("request", req))
		return err
	}
	job, err := ja.JobGorm.Get(ctx, req.JobID)
	if err != nil {
		ja.Logger.Error(err.Error())
		return err
	}
	notiCandidate := models.Notification{
		CandidateID: req.CandidateID,
		Title:       "Apply job " + job.Title + " successful",
		Content:     "The recruiter has received your CV!! Good luck!!",
		Key:         "job",
		CheckRead:   false,
	}
	notiRecruiter := models.Notification{
		RecruiterID: job.RecruiterID,
		Title:       "A candidate apply to your job",
		Content:     job.Title + ":This job has a new candidate. Let check it!!",
		Key:         "job-apply",
		CheckRead:   false,
	}
	notis := make([]*models.Notification, 0)
	notis = append(notis, &notiCandidate)
	notis = append(notis, &notiRecruiter)
	err = ja.NotiGorm.Create(ctx, notis)
	if err != nil {
		ja.Logger.Error(err.Error())
	}
	return nil
}

func (ja *JobApply) GetJobByCandidateID(ctx context.Context, req models.RequestGetJobApplyByCandidateID) (*models.ResponseGetJobApply, error) {
	offset := (req.Page - 1) * req.Size
	jobs, total, err := ja.JobApplyGorm.GetAppliedJobByCandidateID(ctx, req.CandidateID, offset, req.Size)
	if err != nil {
		ja.Logger.Error("Can not get jobs", zap.Error(err), zap.Int64("candidate_id", req.CandidateID))
		return nil, err
	}
	jobsWithSkills := make([]models.JobWithSkill, 0)
	for _, job := range jobs {
		skills, err := ja.JobSkill.GetAllSkillByJob(ctx, job.JobID)
		if err != nil {
			ja.Logger.Error(err.Error(), zap.Int64("Job ID", job.JobID))
			continue
		}
		rwk := models.JobWithSkill{
			Job:    job,
			Skills: skills,
		}
		jobsWithSkills = append(jobsWithSkills, rwk)
	}
	resp := models.ResponseGetJobApply{
		Total:  total,
		Result: jobsWithSkills,
	}
	return &resp, nil
}

func (ja *JobApply) GetCandidatesApplyJob(ctx context.Context, req models.RequestGetJobApplyByJobID) (*models.ResponseGetCandidateApply, error) {
	offset := (req.Page - 1) * req.Size
	candidates, total, err := ja.JobApplyGorm.GetCandidateApplyJob(ctx, req.JobID, offset, req.Size)
	if err != nil {
		ja.Logger.Error("Can not get candidates", zap.Error(err), zap.Int64("job_id", req.JobID))
		return nil, err
	}
	resp := models.ResponseGetCandidateApply{
		Total:  total,
		Result: candidates,
	}
	return &resp, nil
}

func (ja *JobApply) UpdateStatusJobApplied(ctx context.Context, req models.RequestUpdateStatusJobApplied) error {
	err := ja.JobApplyGorm.UpdateStatus(ctx, req.Status, req.JobID, req.CandidateID)
	if err != nil {
		ja.Logger.Error(err.Error())
		return err
	}
	job, err := ja.JobGorm.Get(ctx, req.JobID)
	if err != nil {
		ja.Logger.Error(err.Error())
		return err
	}
	noti := models.Notification{
		CandidateID: req.CandidateID,
		Title:       "Your apply to " + job.Title + " has new!!",
		Content:     "The recruiter has update status your application!! Let check it!!",
		Key:         "job-apply",
		CheckRead:   false,
	}
	notis := make([]*models.Notification, 0)
	notis = append(notis, &noti)
	err = ja.NotiGorm.Create(ctx, notis)
	if err != nil {
		ja.Logger.Error(err.Error())
	}
	return nil
}

func (ja *JobApply) CountCandidateByStatus(ctx context.Context, req models.RequestCountStatus) (int64, error) {
	return ja.JobApplyGorm.CountByStatus(ctx, req.Status)
}
func (ja *JobApply) CheckApplied(ctx context.Context, req models.RequestCheckApply) (*models.JobApply, error) {
	return ja.JobApplyGorm.CheckApplied(ctx, req.JobID, req.CandidateID)
}

func (ja *JobApply) GetApply(ctx context.Context, req models.RequestCheckApply) (*models.JobApply, error) {
	return ja.JobApplyGorm.GetApply(ctx, req.JobID, req.CandidateID)
}
