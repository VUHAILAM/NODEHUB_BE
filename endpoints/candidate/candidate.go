package candidate

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	cs.Logger.Info("", zap.Reflect("canReq", req))
	if req.CandidateID == 0 {
		req.CandidateID = acc.(models.Account).Id
	}
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

//recruiter admin
func (cs *CandidateSerializer) GetAllCandidateForAdmin(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListCandidateAdmin{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := cs.CandidateService.GetAllCandidateForAdmin(ctx, req.Name, req.Page, req.Size)
	if err != nil {
		cs.Logger.Error("GetAllCandidateForAdmin error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (cs *CandidateSerializer) UpdateReviewCandidateByAdmin(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestUpdateReviewCandidateAdmin{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request Update review candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = cs.CandidateService.UpdateReviewCandidateByAdmin(ctx, &req)
	if err != nil {
		cs.Logger.Error("Update Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (cs *CandidateSerializer) UpdateStatusCandidate(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestUpdateStatusCandidate{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request Update status candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = cs.CandidateService.UpdateStatusCandidate(ctx, &req, req.ID)
	if err != nil {
		cs.Logger.Error("Update status candidate error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

// CandidateSkill
func (cs *CandidateSerializer) AddCandidateSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.CandidateSkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		cs.Logger.Error("Parse request AddCandidateSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = cs.CandidateService.AddCandidateSkill(ctx, &req)
	if err != nil {
		cs.Logger.Error("AddRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (cs *CandidateSerializer) DeleteCandidateSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := ginCtx.Query("candidate_skill_id")
	n, err1 := strconv.ParseInt(req, 10, 64)
	err := cs.CandidateService.DeleteCandidateSkill(ctx, n)
	if err1 != nil {
		cs.Logger.Error(" DeleteCandidateSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		cs.Logger.Error(" DeleteCandidateSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (cs *CandidateSerializer) UpdateCandidateSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestUpdateCandidateSkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		cs.Logger.Error("Parse request Update Candidate Skill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = cs.CandidateService.UpdateCandidateSkill(ctx, &req)
	if err != nil {
		cs.Logger.Error("Update Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (cs *CandidateSerializer) GetCandidateSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := ginCtx.Query("candidate_id")
	n, err1 := strconv.ParseInt(req, 10, 64)
	data, err := cs.CandidateService.GetCandidateSkill(ctx, n)
	if err1 != nil {
		cs.Logger.Error("GetCandidateSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		cs.Logger.Error("GetCandidateSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}
