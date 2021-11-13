package candidate

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/candidate"
	"go.uber.org/zap"
)

type CandidateSerializer struct {
	CandidateService candidate.ICandidateService
	Logger           *zap.Logger
}

func NewCandidateSerializer(service *candidate.CandidateService, logger *zap.Logger) *CandidateSerializer {
	return &CandidateSerializer{
		CandidateService: service,
		Logger:           logger,
	}
}

func (cs *CandidateSerializer) CreateProfile(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		cs.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.CandidateRequest{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request Candidate Create error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.CandidateID = acc.(models.Account).Id
	_, err = cs.CandidateService.CreateCandidateProfile(ctx, req)
	if err != nil {
		cs.Logger.Error(err.Error())
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (cs *CandidateSerializer) UpdateProfile(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		cs.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.CandidateRequest{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request Candidate Update error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.CandidateID = acc.(models.Account).Id
	err = cs.CandidateService.UpdateCandidateProfile(ctx, req)
	if err != nil {
		cs.Logger.Error(err.Error())
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (cs *CandidateSerializer) GetProfile(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		cs.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.CandidateRequest{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request Candidate Get detail error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.CandidateID = acc.(models.Account).Id
	candidate, err := cs.CandidateService.GetCandidateProfile(ctx, req.CandidateID)
	if err != nil {
		cs.Logger.Error(err.Error())
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, candidate)
}
