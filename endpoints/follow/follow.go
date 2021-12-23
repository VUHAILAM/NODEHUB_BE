package follow

import (
	"encoding/json"
	"net/http"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
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

	req := models2.RequestFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Follow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models2.Account).Id
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

	req := models2.RequestUnfollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request UnFollow error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models2.Account).Id
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
	req := models2.RequestFollow{}
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
	req := models2.RequestFollow{}
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
	req := models2.RequestFollow{}
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
		ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, models2.Follow{})
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

	req := models2.RequestGetCandidateFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request List Follow Candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.RecruiterID = acc.(models2.Account).Id
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

	req := models2.RequestGetRecruiterFollow{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request List Follow Candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.CandidateID = acc.(models2.Account).Id
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
