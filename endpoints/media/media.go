package media

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"gitlab.com/hieuxeko19991/job4e_be/services/media"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"go.uber.org/zap"
)

type MediaSerializer struct {
	mediaService media.IMediaService
	Logger       *zap.Logger
}

func NewMediaSerializer(mediaService media.IMediaService, logger *zap.Logger) *MediaSerializer {
	return &MediaSerializer{
		mediaService: mediaService,
		Logger:       logger,
	}
}

func (m *MediaSerializer) CreateMedia(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCreateMedia{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		m.Logger.Error("Parse request create media error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = m.mediaService.CreateMedia(ctx, &req)

	if err != nil {
		m.Logger.Error("Create meida error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (m *MediaSerializer) UpdateMedia(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestUpdateMedia{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		m.Logger.Error("Parse request update media error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = m.mediaService.UpdateMedia(ctx, &req, req.Media_id)
	if err != nil {
		m.Logger.Error("Update Media error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (m *MediaSerializer) GetListMediaPaging(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListMedia{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		m.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := m.mediaService.GetListMediaPaging(ctx, req.Name, req.Page, req.Size)
	if err != nil {
		m.Logger.Error("GetListMediaPaging error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (m *MediaSerializer) GetSlide(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	data, err := m.mediaService.GetSlide(ctx)
	if err != nil {
		m.Logger.Error("GetSlide error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}
