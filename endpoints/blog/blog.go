package blog

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"gitlab.com/hieuxeko19991/job4e_be/services/blog"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"go.uber.org/zap"
)

type BlogSerializer struct {
	blogService blog.IBlogService
	Logger      *zap.Logger
}

func NewBlogSerializer(blogService blog.IBlogService, logger *zap.Logger) *BlogSerializer {
	return &BlogSerializer{
		blogService: blogService,
		Logger:      logger,
	}
}

func (bl *BlogSerializer) GetDetail(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetDetailBlog{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		bl.Logger.Error("Parse request get detail blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	blog, err := bl.blogService.GetDetailBlog(ctx, req.BlogID)
	if err != nil {
		bl.Logger.Error("Get detail blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, blog)
}

func (bl *BlogSerializer) Getlist(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListBlog{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		bl.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := bl.blogService.GetListBlog(ctx, req.Title, req.Page, req.Size)
	if err != nil {
		bl.Logger.Error("getlistBlog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (bl *BlogSerializer) GetListBlogUser(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListBlog{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		bl.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := bl.blogService.GetListBlogUser(ctx, req.Title, req.Category_id, req.Page, req.Size)
	if err != nil {
		bl.Logger.Error("GetListBlogUser error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}
func (bl *BlogSerializer) GetListBlogByCategory(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListBlogByCategoryId{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		bl.Logger.Error("Parse request get list error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := bl.blogService.GetListBlogByCategory(ctx, req.Category_id, req.Page, req.Size)
	if err != nil {
		bl.Logger.Error("GetListBlogByCategory error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (bl *BlogSerializer) CreateBlog(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCreateBlog{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		bl.Logger.Error("Parse request create blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = bl.blogService.CreateBlog(ctx, &req)
	if err != nil {
		bl.Logger.Error("Create blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (bl *BlogSerializer) UpdateBlog(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCreateBlog{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		bl.Logger.Error("Parse request update blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = bl.blogService.UpdateBlog(ctx, &req, req.Blog_id)
	if err != nil {
		bl.Logger.Error("Update blog error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}
