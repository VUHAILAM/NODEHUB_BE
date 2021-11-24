package notification

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type INotificationService interface {
	CreateNotification(ctx context.Context, notification *models.RequestCreateNotification) error
	GetListNotificationByAccount(ctx context.Context, candidateID int64, page int64, size int64) (*models.ResponseListNotification, error)
}

type Notification struct {
	NotificationGorm *NotificationGorm
	Logger           *zap.Logger
}

func NewNotification(notificationGorm *NotificationGorm, logger *zap.Logger) *Notification {
	return &Notification{
		NotificationGorm: notificationGorm,
		Logger:           logger,
	}
}

/*Create Notification*/
func (n *Notification) CreateNotification(ctx context.Context, notification *models.RequestCreateNotification) error {
	notificationModels := &models.Notification{
		CandidateID: notification.CandidateID,
		Content:     notification.Content,
		Title:       notification.Title,
		Key:         notification.Key,
		CheckRead:   notification.CheckRead,
	}
	err := n.NotificationGorm.Create(ctx, notificationModels)
	if err != nil {
		return err
	}
	return nil
}

func (n *Notification) GetListNotificationByAccount(ctx context.Context, candidateID int64, page int64, size int64) (*models.ResponseListNotification, error) {
	offset := (page - 1) * size
	noti, total, err := n.NotificationGorm.GetListNotificationByAccount(ctx, candidateID, offset, size)
	if err != nil {
		n.Logger.Error(err.Error(), zap.Int64("candidate id", candidateID))
		return nil, err
	}
	resp := models.ResponseListNotification{
		Total:         total,
		Notifications: noti,
	}
	return &resp, nil
}
