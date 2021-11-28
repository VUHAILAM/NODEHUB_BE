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
		s.Logger.Error("Parse request Apply error", zap.Error(err))
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

func (s *JobApplySerializer) GetJobAppliedByJobID(ginCtx *gin.Context) {
	//ctx := ginCtx.Request.Context()
	//req := models.RequestGetJobApplyByJobID{}
	//err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	//if err != nil {
	//	s.Logger.Error("Parse request Get Job Apply error", zap.Error(err))
	//	ginx.BuildErrorResponse(ginCtx, err, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//resp, err := s.JobApplyService.GetJobsByJobID(ctx, req)
	//if err != nil {
	//	s.Logger.Error("GetJobAppliedByJobID error", zap.Error(err))
	//	ginx.BuildErrorResponse(ginCtx, err, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}

func (s *JobApplySerializer) GetJobAppliedByCandidateID(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.RequestGetJobApplyByCandidateID{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Job Create error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models.Account).Id

	resp, err := s.JobApplyService.GetJobByCandidateID(ctx, req)
	if err != nil {
		s.Logger.Error("GetJobAppliedByCandidateID error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}

func (s *JobApplySerializer) GetCandidatesAppyJob(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetJobApplyByJobID{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Candidate get error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := s.JobApplyService.GetCandidatesApplyJob(ctx, req)
	if err != nil {
		s.Logger.Error("GetCandidatesApplyJob error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)

}

func (s *JobApplySerializer) UpdateStatus(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestUpdateStatusJobApplied{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Job Create error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = s.JobApplyService.UpdateStatusJobApplied(ctx, req)
	if err != nil {
		s.Logger.Error("Update status error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (s *JobApplySerializer) CountCandidateByStatus(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCountStatus{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Count status error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	count, err := s.JobApplyService.CountCandidateByStatus(ctx, req)
	if err != nil {
		s.Logger.Error("Count Status error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, count)
}

func (s *JobApplySerializer) CheckApplied(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCheckApply{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Check apply error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	jobapply, err := s.JobApplyService.CheckApplied(ctx, req)
	if err != nil {
		s.Logger.Error("Count Status error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, jobapply)
}
