package skill

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/services/skill"
	"go.uber.org/zap"
)

type SkillSerializer struct {
	skillService skill.ISkillService
	Logger       *zap.Logger
}

func NewSkillSerializer(skillService skill.ISkillService, logger *zap.Logger) *SkillSerializer {
	return &SkillSerializer{
		skillService: skillService,
		Logger:       logger,
	}
}

/*Create Skill*/
func (s *SkillSerializer) CreateSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCreateSkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		s.Logger.Error("Parse request create skill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = s.skillService.CreateSkill(ctx, &req)
	if err != nil {
		s.Logger.Error("Create skill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

/*Update Skill*/
func (s *SkillSerializer) UpdateSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCreateSkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		s.Logger.Error("Parse request update skill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = s.skillService.UpdateSkill(ctx, &req, req.Skill_id)
	if err != nil {
		s.Logger.Error("Update blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

/*Get List SKill*/
func (s *SkillSerializer) GetlistSkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListSkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request get list skill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := s.skillService.GetListSkill(ctx, req.Name, req.Page, req.Size)
	if err != nil {
		s.Logger.Error("getlistSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

/*Get All SKill*/
func (s *SkillSerializer) GetAll(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := ginCtx.Query("name")
	data, err := s.skillService.GetAll(ctx, req)
	if err != nil {
		s.Logger.Error("getAllSkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}
