package job

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

const (
	jobIndex = "test-nodehub-job"
)

type IJobElasticsearch interface {
	Create(ctx context.Context, documentID string, data map[string]interface{}) error
	GetJobByID(ctx context.Context, documentID string) (*models.ESJob, error)
	GetAllJob(ctx context.Context, from, size int64) ([]models.ESJob, int64, error)
	Update(ctx context.Context, documentID string, data map[string]interface{}) error
	GetJobsByRecruiterID(ctx context.Context, recruiterID, from, size int64) ([]models.ESJob, int64, error)
	Delete(ctx context.Context, documentID string) error
	SearchJobs(ctx context.Context, text, location string, from, size int64) ([]models.ESJob, int64, error)
}

type JobES struct {
	ES       *elastic.Client
	JobIndex string
	Logger   *zap.Logger
}

func NewJobES(es *elastic.Client, jobindex string, logger *zap.Logger) *JobES {
	return &JobES{
		ES:       es,
		JobIndex: jobindex,
		Logger:   logger,
	}
}

func (e *JobES) Create(ctx context.Context, documentID string, data map[string]interface{}) error {
	_, err := e.ES.Index().Index(e.JobIndex).BodyJson(data).Id(documentID).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Create Job error", zap.Error(err))
		return err
	}
	return nil
}

func (e *JobES) GetAllJob(ctx context.Context, from, size int64) ([]models.ESJob, int64, error) {
	searchService := e.ES.Search().Index(e.JobIndex)
	searchResult, err := searchService.Sort("created_at", false).From(int(from)).Size(int(size)).Pretty(true).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Get Job List error", zap.Error(err))
		return nil, 0, err
	}

	jobs := make([]models.ESJob, 0, size)

	for _, hit := range searchResult.Hits.Hits {
		var j models.ESJob
		err := json.Unmarshal(hit.Source, &j)
		if err != nil {
			e.Logger.Error(err.Error())
			continue
		}
		e.Logger.Info("job es", zap.Reflect("job es", j))
		jobs = append(jobs, j)
	}
	return jobs, searchResult.TotalHits(), nil
}

func (e *JobES) GetJobByID(ctx context.Context, documentID string) (*models.ESJob, error) {
	res, err := e.ES.Get().Index(e.JobIndex).Id(documentID).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Get Job error", zap.Error(err))
		return nil, err
	}
	if !res.Found {
		e.Logger.Error("Job with ID not found", zap.String("ID", documentID))
		return nil, errors.New("Job with ID not found")
	}
	job := models.ESJob{}
	err = json.Unmarshal(res.Source, &job)
	if err != nil {
		e.Logger.Error("Unmarshal Job error", zap.Error(err))
		return nil, err
	}
	return &job, nil
}

func (e *JobES) GetJobsByRecruiterID(ctx context.Context, recruiterID, from, size int64) ([]models.ESJob, int64, error) {
	searchService := e.ES.Search().Index(e.JobIndex).Query(elastic.NewTermQuery("recruiter_id", recruiterID))
	searchResult, err := searchService.Sort("created_at", false).From(int(from)).Size(int(size)).Pretty(true).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Get Job List error", zap.Error(err))
		return nil, 0, err
	}
	jobs := make([]models.ESJob, 0, size)

	for _, hit := range searchResult.Hits.Hits {
		var j models.ESJob
		err := json.Unmarshal(hit.Source, &j)
		if err != nil {
			e.Logger.Error(err.Error())
			continue
		}
		e.Logger.Info("job es", zap.Reflect("job es", j))
		jobs = append(jobs, j)
	}
	return jobs, searchResult.TotalHits(), nil
}

func (e *JobES) Update(ctx context.Context, documentID string, data map[string]interface{}) error {
	_, err := e.ES.Update().Index(e.JobIndex).Id(documentID).Doc(data).DetectNoop(true).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Update Job error", zap.Error(err))
		return err
	}
	return nil
}

func (e *JobES) Delete(ctx context.Context, documentID string) error {
	_, err := e.ES.Delete().Index(e.JobIndex).Id(documentID).Do(ctx)
	if err != nil {
		e.Logger.Error("Job ES: Delete Job error", zap.Error(err))
		return err
	}
	return nil
}

func (e *JobES) SearchJobs(ctx context.Context, text, location string, from, size int64) ([]models.ESJob, int64, error) {
	txtQuery := elastic.NewMultiMatchQuery(text, "title", "role", "company_name").Type("most_fields")
	locationQuery := elastic.NewMatchQuery("location", location)
	skillQuery := elastic.NewNestedQuery("skills", elastic.NewMatchQuery("skills.name", text))
	generalQuery := elastic.NewBoolQuery().Must(locationQuery, elastic.NewBoolQuery().Should(skillQuery, txtQuery))
	var searchResult *elastic.SearchResult
	var err error
	if location == "" && text != "" {
		searchResult, err = e.ES.Search().Index(e.JobIndex).Query(elastic.NewBoolQuery().Should(skillQuery, txtQuery)).
			From(int(from)).Size(int(size)).Sort("_score", false).Do(ctx)
	} else if text == "" && location != "" {
		searchResult, err = e.ES.Search().Index(e.JobIndex).Query(locationQuery).
			From(int(from)).Size(int(size)).Sort("_score", false).Do(ctx)
	} else {
		searchResult, err = e.ES.Search().Index(e.JobIndex).Query(generalQuery).
			From(int(from)).Size(int(size)).Sort("_score", false).Do(ctx)
	}
	if err != nil {
		e.Logger.Error("Job ES: Get Job List error", zap.Error(err))
		return nil, 0, err
	}
	jobs := make([]models.ESJob, 0, size)

	for _, hit := range searchResult.Hits.Hits {
		var j models.ESJob
		err := json.Unmarshal(hit.Source, &j)
		if err != nil {
			e.Logger.Error(err.Error())
			continue
		}
		e.Logger.Info("job es", zap.Reflect("job es", j))
		jobs = append(jobs, j)
	}
	return jobs, searchResult.TotalHits(), nil
}
