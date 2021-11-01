package job

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"

	"go.uber.org/zap"

	"github.com/olivere/elastic"
)

const (
	jobIndex = ""
)

type IJobElasticsearch interface {
	Create(ctx context.Context, documentID string, data map[string]interface{}) error
	GetJobList(ctx context.Context, queries []*elastic.TermQuery) (*elastic.SearchResult, error)
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
	_, err := e.ES.Index().Index(jobIndex).Type("job").Id(documentID).BodyJson(data).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Create Job error", zap.Error(err))
		return err
	}
	return nil
}

func (e *JobES) GetJobList(ctx context.Context, queries []*elastic.TermQuery) (*elastic.SearchResult, error) {
	searchService := e.ES.Search().Index(jobIndex)
	searchQueries := []elastic.Query{}
	for _, q := range queries {
		searchQueries = append(searchQueries, q)
	}
	searchResult, err := searchService.Query(elastic.NewBoolQuery().Must(searchQueries...)).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Get Job List error", zap.Error(err))
		return nil, err
	}
	return searchResult, nil
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
	err = json.Unmarshal(*res.Source, &job)
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
