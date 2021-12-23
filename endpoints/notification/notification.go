package notification

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/services/notification"
	"go.uber.org/zap"
)

type NotificationSerializer struct {
	notificationService notification.INotificationService
	Logger              *zap.Logger
}

func NewNotificationSerializer(notificationService notification.INotificationService, logger *zap.Logger) *NotificationSerializer {
	return &NotificationSerializer{
		notificationService: notificationService,
		Logger:              logger,
	}
}

/*Get List notification*/
func (n *NotificationSerializer) GetListNotificationByAccount(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListNotification{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		n.Logger.Error("Parse request get list NotificationService error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := n.notificationService.GetListNotificationByCandidate(ctx, req.CandidateID, req.Page, req.Size)
	if err != nil {
		n.Logger.Error("getlistNotification error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (n *NotificationSerializer) GetListNotificationByRecruiter(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestGetListNotification{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		n.Logger.Error("Parse request get list NotificationService error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	data, err := n.notificationService.GetListNotificationByRecruiter(ctx, req.RecruiterID, req.Page, req.Size)
	if err != nil {
		n.Logger.Error("getlistNotification error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"data": data,
	})
}

func (n *NotificationSerializer) MarkRead(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestMarkRead{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		n.Logger.Error("Parse request Mark Read error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = n.notificationService.MarkRead(ctx, req)
	if err != nil {
		n.Logger.Error("Mark read error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (n *NotificationSerializer) MarkReadAll(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestMarkReadAll{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		n.Logger.Error("Parse request Mark Read All error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = n.notificationService.MarkReadAll(ctx, req)
	if err != nil {
		n.Logger.Error("Mark read All error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (n *NotificationSerializer) CountUnread(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestCountUnread{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		n.Logger.Error("Parse request Count unread error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	count, err := n.notificationService.CountUnread(ctx, req)
	if err != nil {
		n.Logger.Error("Count unread error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, count)
}
