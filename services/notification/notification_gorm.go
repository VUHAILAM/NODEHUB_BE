package notification

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableNotification = "notification"

type INotificationDatabase interface {
	Create(ctx context.Context, notification []*models.Notification) error
	GetListNotificationByCandidate(ctx context.Context, candidateID int64, offset int64, size int64) ([]*models.Notification, int64, error)
	GetListNotificationByRecruiter(ctx context.Context, recruiterID int64, offset int64, size int64) ([]*models.Notification, int64, error)
	UpdateCheckRead(ctx context.Context, notificationID int64) error
	UpdateCheckReadByAccountID(ctx context.Context, field string, accountID int64) error
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

/*Create NotificationService*/
func (n *NotificationGorm) Create(ctx context.Context, notification []*models.Notification) error {
	db := n.db.WithContext(ctx)
	err := db.Table(tableNotification).Create(&notification).Error
	if err != nil {
		n.logger.Error("NotificationGorm: Create notification error", zap.Error(err))
		return err
	}
	return nil
}

/*Get NotificationService*/
func (n *NotificationGorm) GetListNotificationByCandidate(ctx context.Context, candidateID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
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

func (n *NotificationGorm) GetListNotificationByRecruiter(ctx context.Context, recruiterID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
	var notifications []*models.Notification
	db := n.db.WithContext(ctx).Table(tableNotification).Where("recruiter_id=?", recruiterID).Find(&notifications)
	total := db.RowsAffected
	notifications = make([]*models.Notification, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Find(&notifications).Order("created_at desc").Error

	if err != nil {
		n.logger.Error(err.Error())
		return nil, 0, err
	}
	return notifications, total, nil
}

func (n *NotificationGorm) UpdateCheckRead(ctx context.Context, notificationID int64) error {
	db := n.db.WithContext(ctx)
	err := db.Table(tableNotification).Where("notification_id=?", notificationID).Update("check_read", true).Error
	if err != nil {
		n.logger.Error(err.Error())
		return err
	}
	return nil
}

func (n *NotificationGorm) UpdateCheckReadByAccountID(ctx context.Context, field string, accountID int64) error {
	db := n.db.WithContext(ctx)
	err := db.Table(tableNotification).Where(field+"=?", accountID).Update("check_read", true).Error
	if err != nil {
		n.logger.Error(err.Error())
		return err
	}
	return nil
}
