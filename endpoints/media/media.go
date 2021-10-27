package media

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/services/media"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
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

// func (c *CategorySerializer) UpdateCategory(ginCtx *gin.Context) {
// 	ctx := ginCtx.Request.Context()
// 	req := models.RequestCreateSetting{}
// 	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

// 	if err != nil {
// 		c.Logger.Error("Parse request update Category error", zap.Error(err))
// 		ginx.BuildErrorResponse(ginCtx, err, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	err = c.categoryService.UpdateCategory(ctx, &req, req.Setting_id)
// 	if err != nil {
// 		c.Logger.Error("Update Category error", zap.Error(err))
// 		ginx.BuildErrorResponse(ginCtx, err, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
// }

// func (c *CategorySerializer) GetListCategoryPaging(ginCtx *gin.Context) {
// 	ctx := ginCtx.Request.Context()
// 	req := models.RequestGetListSetting{}
// 	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
// 	if err != nil {
// 		c.Logger.Error("Parse request get list error", zap.Error(err))
// 		ginx.BuildErrorResponse(ginCtx, err, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	data, err := c.categoryService.GetListCategoryPaging(ctx, req.Name, req.Page, req.Size)
// 	if err != nil {
// 		c.Logger.Error("GetListCategoryAdmin error", zap.Error(err))
// 		ginx.BuildErrorResponse(ginCtx, err, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
// 		"data": data,
// 	})
// }

// func (c *CategorySerializer) GetAllCategory(ginCtx *gin.Context) {
// 	ctx := ginCtx.Request.Context()
// 	data, err := c.categoryService.GetAllCategory(ctx)
// 	if err != nil {
// 		c.Logger.Error("GetAllCategory error", zap.Error(err))
// 		ginx.BuildErrorResponse(ginCtx, err, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
// 		"data": data,
// 	})
// }
