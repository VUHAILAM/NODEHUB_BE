package notification

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
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
	data, err := n.notificationService.GetListNotificationByAccount(ctx, req.CandidateID, req.Page, req.Size)
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
