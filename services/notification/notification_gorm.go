package notification

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableNotification = "notification"

type INotificationDatabase interface {
	CreateNotification(ctx context.Context, notification *models.RequestCreateNotification) error
	GetListNotificationByAccount(ctx context.Context, candidateID int64, offset int64, size int64) (*models.ResponsetListNotification, error)
}

type NotificationGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewNotificationGorm(db *gorm.DB, logger *zap.Logger) *NotificationGorm {
	return &NotificationGorm{
		db:     db,
		logger: logger,
	}
}

/*Create Notification*/
func (n *NotificationGorm) Create(ctx context.Context, notification *models.Notification) error {
	db := n.db.WithContext(ctx)
	err := db.Table(tableNotification).Create(notification).Error
	if err != nil {
		n.logger.Error("NotificationGorm: Create notification error", zap.Error(err))
		return err
	}
	return nil
}

/*Get Notification*/
func (n *NotificationGorm) GetListNotificationByAccount(ctx context.Context, candidateID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
	db := n.db.WithContext(ctx).Table(tableNotification)
}
