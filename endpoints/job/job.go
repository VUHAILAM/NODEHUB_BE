package job

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/job"
	"go.uber.org/zap"
)

type JobSerializer struct {
	JobService job.IJobService
	Logger     *zap.Logger
}

func NewJobSerializer(jobService job.IJobService, logger *zap.Logger) *JobSerializer {
	return &JobSerializer{
		JobService: jobService,
		Logger:     logger,
	}
}

func (s *JobSerializer) Create(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.CreateJobRequest{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Job Create error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.RecruiterID = acc.(models.Account).Id
	err = s.JobService.CreateNewJob(ctx, &req)
	if err != nil {
		s.Logger.Error("Create New Job error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (s *JobSerializer) GetDetailJob(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetJobDetail{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Get Detail Job error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	job, err := s.JobService.GetDetailJob(ctx, req.JobID)
	if err != nil {
		s.Logger.Error("Get Detail Job error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, job)
}

func (s *JobSerializer) UpdateJob(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestUpdateJob{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Update Job error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = s.JobService.UpdateJob(ctx, req)
	if err != nil {
		s.Logger.Error("Update Job error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}
