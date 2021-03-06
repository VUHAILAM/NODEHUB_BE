package recruiter

import (
	"encoding/json"
	"net/http"
	"strconv"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"

	"go.uber.org/zap"
)

type RecruiterSerializer struct {
	recruiterService recruiter.IRecruiterService
	Logger           *zap.Logger
}

func NewRecruiterSerializer(recruiterService recruiter.IRecruiterService, logger *zap.Logger) *RecruiterSerializer {
	return &RecruiterSerializer{
		recruiterService: recruiterService,
		Logger:           logger,
	}
}

func (r *RecruiterSerializer) GetProfileRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := ginCtx.Query("recruiter_id")
	n, err1 := strconv.ParseInt(req, 10, 64)
	data, err := r.recruiterService.GetProfileRecruiter(ctx, n)
	if err1 != nil {
		r.Logger.Error("GetProfileRecruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		r.Logger.Error("GetProfileRecruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (r *RecruiterSerializer) UpdateProfile(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestUpdateRecruiter{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		r.Logger.Error("Parse request update recruiter profile error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = r.recruiterService.UpdateProfile(ctx, &req, req.RecruiterID)
	if err != nil {
		r.Logger.Error("Update recruiter profile error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

// RecruiterSkill
func (r *RecruiterSerializer) AddRecruiterSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RecruiterSkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		r.Logger.Error("Parse request AddRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = r.recruiterService.AddRecruiterSkill(ctx, &req)
	if err != nil {
		r.Logger.Error("AddRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (r *RecruiterSerializer) GetRecruiterSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := ginCtx.Query("recruiter_id")
	n, err1 := strconv.ParseInt(req, 10, 64)
	data, err := r.recruiterService.GetRecruiterSkill(ctx, n)
	if err1 != nil {
		r.Logger.Error("GetRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		r.Logger.Error("GetRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (r *RecruiterSerializer) DeleteRecruiterSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := ginCtx.Query("recruiter_skill_id")
	n, err1 := strconv.ParseInt(req, 10, 64)
	err := r.recruiterService.DeleteRecruiterSkill(ctx, n)
	if err1 != nil {
		r.Logger.Error(" DeleteRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		r.Logger.Error(" DeleteRecruiterSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

//recruiter admin
func (r *RecruiterSerializer) GetListRecruiterForAdmin(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestGetListRecruiter{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := r.recruiterService.GetAllRecruiterForAdmin(ctx, req.Name, req.Page, req.Size)
	if err != nil {
		r.Logger.Error("GetListRecruiterForAdmin error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (r *RecruiterSerializer) UpdateReciuterByAdmin(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestUpdateRecruiterAdmin{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request Update Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = r.recruiterService.UpdateReciuterByAdmin(ctx, &req)
	if err != nil {
		r.Logger.Error("Update Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}
func (r *RecruiterSerializer) UpdateStatusReciuter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestUpdateStatusRecruiter{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request Update status Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = r.recruiterService.UpdateStatusReciuter(ctx, &req, req.ID)
	if err != nil {
		r.Logger.Error("Update status Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (r *RecruiterSerializer) GetAllRecruiterForCandidate(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestGetListRecruiterForCandidate{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := r.recruiterService.GetAllRecruiterForCandidate(ctx, req.RecruiterName, req.SkillName, req.Address, req.Page, req.Size)
	if err != nil {
		r.Logger.Error("GetListRecruiterForAdmin error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (r *RecruiterSerializer) SearchRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestSearchRecruiter{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request Search Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := r.recruiterService.SearchRecruiter(ctx, req)
	if err != nil {
		r.Logger.Error("Search Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}

func (r *RecruiterSerializer) GetAllRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestSearchRecruiter{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request Search Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := r.recruiterService.GetAllRecruiter(ctx, req)
	if err != nil {
		r.Logger.Error("Search Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}

func (r *RecruiterSerializer) CountRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	count, err := r.recruiterService.CountRecruiter(ctx)
	if err != nil {
		r.Logger.Error("Count Recruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, count)
}

func (r *RecruiterSerializer) PublicProfile(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models2.RequestPublicProfile{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		r.Logger.Error("Parse request PublicProfile error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := r.recruiterService.GetProfileRecruiter(ctx, req.ID)
	if err != nil {
		r.Logger.Error("GetProfileRecruiter error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})

}

func (r *RecruiterSerializer) CheckPremium(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		r.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}

	recruiterID := acc.(models2.Account).Id

	premium, err := r.recruiterService.CheckPremium(ctx, recruiterID)
	if err != nil {
		r.Logger.Error("CheckPremium error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"premium": premium,
	})
}
