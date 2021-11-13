package job_skill

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gitlab.com/hieuxeko19991/job4e_be/services/job_skill"
	"go.uber.org/zap"
)

type JobSkillSerializer struct {
	JobSkillService job_skill.IJobSkillService
	Logger          *zap.Logger
}

func NewJobSkillSerializer(jskill *job_skill.JobSkill, logger *zap.Logger) *JobSkillSerializer {
	return &JobSkillSerializer{
		JobSkillService: jskill,
		Logger:          logger,
	}
}

func (s *JobSkillSerializer) GetJobsBySkill(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetJobsBySkill{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Jobs By Skill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := s.JobSkillService.GetJobBySkill(ctx, req)
	if err != nil {
		s.Logger.Error("GetJobsBySkill error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}

func (s *JobSkillSerializer) GetSkillsByJob(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetSkillsByJob{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		s.Logger.Error("Parse request Skills By Job error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := s.JobSkillService.GetSkillsByJob(ctx, req)
	if err != nil {
		s.Logger.Error("GetSkillsByJob error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, resp)
}
