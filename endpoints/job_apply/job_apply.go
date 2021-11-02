package job_apply

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/job_apply"
	"go.uber.org/zap"
)

type JobApplySerializer struct {
	JobApplyService job_apply.IJobApplyService
	Logger          *zap.Logger
}

func NewJobApplySerializer(service job_apply.IJobApplyService, logger *zap.Logger) *JobApplySerializer {
	return &JobApplySerializer{
		JobApplyService: service,
		Logger:          logger,
	}
}

func (s *JobApplySerializer) Apply(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.RequestApply{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Job Create error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models.Account).Id
	err = s.JobApplyService.CreateJobApply(ctx, req)
	if err != nil {
		s.Logger.Error("ApplyJob error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}
