package job

import (
	"context"
	"math"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	jobTable = "job"
)

type IJobDatabase interface {
	Create(ctx context.Context, job *models.Job) (*models.Job, error)
	Get(ctx context.Context, jobID int64) (*models.Job, error)
	Update(ctx context.Context, jobID int64, data map[string]interface{}) error
	Count(ctx context.Context) (int64, error)
}

type JobGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewJobGorm(db *gorm.DB, logger *zap.Logger) *JobGorm {
	return &JobGorm{
		DB:     db,
		Logger: logger,
	}
}

func (g *JobGorm) Create(ctx context.Context, job *models.Job) (*models.Job, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobTable).Create(job).Error
	if err != nil {
		g.Logger.Error("JobGorm: Create job error", zap.Error(err))
		return nil, err
	}
	return job, nil
}

func (g *JobGorm) Get(ctx context.Context, jobID int64) (*models.Job, error) {
	db := g.DB.WithContext(ctx)
	job := models.Job{}
	err := db.Table(jobTable).Select("job.*").Where("job_id=?", jobID).First(&job).Error
	if err != nil {
		g.Logger.Error("JobGorm: Get job error", zap.Error(err), zap.Int64("job_id", jobID))
		return nil, err
	}
	return &job, nil
}

func (g *JobGorm) Update(ctx context.Context, jobID int64, data map[string]interface{}) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(jobTable).Where("job_id=?", jobID).Updates(data).Error
	if err != nil {
		g.Logger.Error("JobGorm: Update job error", zap.Error(err), zap.Int64("job_id", jobID))
		return err
	}
	return nil
}

/*Get list Job for admin*/
func (j *JobGorm) GetAllJobForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListJobAdmin, error) {
	db := j.DB.WithContext(ctx)
	arr := []models.JobForAdmin{}
	resutl := models.ResponsetListJobAdmin{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`SELECT j.job_id, j.recruiter_id ,r.name as recruiter_name, j.title, j.description, j.salary_range, j.quantity, j.role, j.experience, j.location, j.hire_date, j.status, j.created_at, j.updated_at 
	FROM nodehub.job j
	LEFT JOIN nodehub.recruiter r
	ON j.recruiter_id  = r.recruiter_id
	where r.name like ?
	ORDER BY j.updated_at desc LIMIT ?, ?`, "%"+name+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) 
	FROM nodehub.job j 
	LEFT JOIN nodehub.recruiter r
	ON j.recruiter_id  = r.recruiter_id
	where r.name like ?`, "%"+name+"%").Scan(&total)
	if err != nil {
		j.Logger.Error("JobGorm: Get List job error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.Total = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}

func (j *JobGorm) UpdateStatusJob(ctx context.Context, job *models.RequestUpdateStatusJob, job_id int64) error {
	db := j.DB.WithContext(ctx)
	err := db.Table(jobTable).Where("job_id=?", job_id).Updates(map[string]interface{}{
		"status": job.Status}).Error
	if err != nil {
		j.Logger.Error("RecruiterGorm: Update status job error", zap.Error(err), zap.Int64("job_id", job_id))
		return err
	}
	return nil
}

func (j *JobGorm) DeleteJob(ctx context.Context, job_id int64) error {
	db := j.DB.WithContext(ctx)
	job := models.Job{}
	err := db.Table(jobTable).Delete(&job, job_id).Error
	if err != nil {
		j.Logger.Error("RecruiterGorm: Delete job error", zap.Error(err), zap.Int64("job_id", job_id))
		return err
	}
	return nil
}

func (j *JobGorm) Count(ctx context.Context) (int64, error) {
	var count int64
	db := j.DB.WithContext(ctx)
	err := db.Table(jobTable).Count(&count).Error
	if err != nil {
		j.Logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}
