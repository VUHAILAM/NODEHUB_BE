package notification

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type INotificationService interface {
	CreateNotification(ctx context.Context, notification *models.RequestCreateNotification) error
	GetListNotificationByAccount(ctx context.Context, account_id int64, page int64, size int64) (*models.ResponsetListNotification, error)
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
		Account_id: notification.Account_id,
		Content:    notification.Content,
		Title:      notification.Title,
		Key:        notification.Key,
		CheckRead:  notification.CheckRead,
	}
	err := n.NotificationGorm.Create(ctx, notificationModels)
	if err != nil {
		return err
	}
	return nil
}

func (n *Notification) GetListNotificationByAccount(ctx context.Context, account_id int64, page int64, size int64) (*models.ResponsetListNotification, error) {
	acc, err := n.NotificationGorm.GetListNotificationByAccount(ctx, account_id, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
