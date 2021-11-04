package job

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"

	"github.com/olivere/elastic/v7"
)

const (
	jobIndex = "test-nodehub-job"
)

type IJobElasticsearch interface {
	Create(ctx context.Context, documentID string, data map[string]interface{}) error
	GetJobByID(ctx context.Context, documentID string) (*models.Job, error)
	GetAllJob(ctx context.Context, from, size int64) ([]models.Job, int64, error)
	Update(ctx context.Context, documentID string, data map[string]interface{}) error
}

type JobES struct {
	ES     *elastic.Client
	Logger *zap.Logger
}

func NewJobES(es *elastic.Client, logger *zap.Logger) *JobES {
	return &JobES{
		ES:     es,
		Logger: logger,
	}
}

func (e *JobES) Create(ctx context.Context, documentID string, data map[string]interface{}) error {
	_, err := e.ES.Index().Index(jobIndex).BodyJson(data).Id(documentID).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Create Job error", zap.Error(err))
		return err
	}
	return nil
}

func (e *JobES) GetAllJob(ctx context.Context, from, size int64) ([]models.Job, int64, error) {
	searchService := e.ES.Search().Index(jobIndex)
	searchResult, err := searchService.Sort("hire_date", false).From(int(from)).Size(int(size)).Pretty(true).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Get Job List error", zap.Error(err))
		return nil, 0, err
	}

	jobs := make([]models.Job, 0, size)
	var j models.Job
	for _, item := range searchResult.Each(reflect.TypeOf(j)) {
		job := item.(models.Job)
		jobs = append(jobs, job)
	}
	return jobs, searchResult.TotalHits(), nil
}

func (e *JobES) GetJobByID(ctx context.Context, documentID string) (*models.Job, error) {
	res, err := e.ES.Get().Index(jobIndex).Id(documentID).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Get Job error", zap.Error(err))
		return nil, err
	}
	if !res.Found {
		e.Logger.Error("Job with ID not found", zap.String("ID", documentID))
		return nil, errors.New("Job with ID not found")
	}
	job := models.Job{}
	err = json.Unmarshal(res.Source, &job)
	if err != nil {
		e.Logger.Error("Unmarshal Job error", zap.Error(err))
		return nil, err
	}
	return &job, nil
}

func (e *JobES) Update(ctx context.Context, documentID string, data map[string]interface{}) error {
	_, err := e.ES.Update().Index(jobIndex).Id(documentID).Doc(data).DetectNoop(true).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Update Job error", zap.Error(err))
		return err
	}
	return nil
}
