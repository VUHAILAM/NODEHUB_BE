package follow

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/follow"
	"go.uber.org/zap"
)

type FollowSerializer struct {
	FollowService follow.IFollowService
	Logger        *zap.Logger
}

func NewFollowSerializer(service *follow.FollowService, logger *zap.Logger) *FollowSerializer {
	return &FollowSerializer{
		FollowService: service,
		Logger:        logger,
	}
}

func (s *FollowSerializer) Follow(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}

	req := models.RequestFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Follow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models.Account).Id
	err = s.FollowService.Follow(ctx, req)
	if err != nil {
		s.Logger.Error("Follow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (s *FollowSerializer) UnFollow(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}

	req := models.RequestUnfollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request UnFollow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models.Account).Id
	err = s.FollowService.UnFollow(ctx, req)
	if err != nil {
		s.Logger.Error("UnFollow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (s *FollowSerializer) CountFollowOfCandidate(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Follow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	count, err := s.FollowService.CountOfCandidate(ctx, req.CandidateID)
	if err != nil {
		s.Logger.Error("Count Follow of Candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, count)
}

func (s *FollowSerializer) CountFollowOfRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Follow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	count, err := s.FollowService.CountOfRecruiter(ctx, req.RecruiterID)
	if err != nil {
		s.Logger.Error("Count Follow of Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, count)
}

func (s *FollowSerializer) FollowExist(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Follow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	follow, err := s.FollowService.FollowExist(ctx, req)
	if err != nil {
		s.Logger.Error("Follow Exist error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, follow)
}

func (s *FollowSerializer) GetListCandidate(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}

	req := models.RequestGetCandidateFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request List Follow Candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.RecruiterID = acc.(models.Account).Id
	resp, err := s.FollowService.GetCandidate(ctx, req)
	if err != nil {
		s.Logger.Error("Get List Follow of Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}

func (s *FollowSerializer) GetListRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		s.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}

	req := models.RequestGetRecruiterFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request List Follow Candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models.Account).Id
	resp, err := s.FollowService.GetRecruiter(ctx, req)
	if err != nil {
		s.Logger.Error("Get List Follow of Candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}
