package notification

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableNotification = "notification"

type INotificationDatabase interface {
	Create(ctx context.Context, notification *models.Notification) error
	GetListNotificationByAccount(ctx context.Context, candidateID int64, offset int64, size int64) ([]*models.Notification, int64, error)
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
	var notifications []*models.Notification
	db := n.db.WithContext(ctx).Table(tableNotification).Where("candidate_id=?", candidateID).Find(&notifications)
	total := db.RowsAffected
	notifications = make([]*models.Notification, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Find(&notifications).Order("created_at desc").Error

	if err != nil {
		n.logger.Error(err.Error())
		return nil, 0, err
	}
	return notifications, total, nil
}
