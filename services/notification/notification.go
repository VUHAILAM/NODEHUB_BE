package notification

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type INotificationService interface {
	CreateNotification(ctx context.Context, notification []*models.RequestCreateNotification) error
	GetListNotificationByCandidate(ctx context.Context, candidateID int64, page int64, size int64) (*models.ResponseListNotification, error)
	GetListNotificationByRecruiter(ctx context.Context, recruiterID int64, page int64, size int64) (*models.ResponseListNotification, error)
}

type NotificationService struct {
	NotificationGorm *NotificationGorm
	Logger           *zap.Logger
}

func NewNotification(notificationGorm *NotificationGorm, logger *zap.Logger) *NotificationService {
	return &NotificationService{
		NotificationGorm: notificationGorm,
		Logger:           logger,
	}
}

/*Create NotificationService*/
func (n *NotificationService) CreateNotification(ctx context.Context, notification []*models.RequestCreateNotification) error {
	notifications := []*models.Notification{}
	for _, noti := range notification {
		notificationModel := &models.Notification{
			RecruiterID: noti.RecruiterID,
			CandidateID: noti.CandidateID,
			Content:     noti.Content,
			Title:       noti.Title,
			Key:         noti.Key,
			CheckRead:   noti.CheckRead,
		}
		notifications = append(notifications, notificationModel)
	}
	err := n.NotificationGorm.Create(ctx, notifications)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationService) GetListNotificationByCandidate(ctx context.Context, candidateID int64, page int64, size int64) (*models.ResponseListNotification, error) {
	offset := (page - 1) * size
	noti, total, err := n.NotificationGorm.GetListNotificationByCandidate(ctx, candidateID, offset, size)
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

func (n *NotificationService) GetListNotificationByRecruiter(ctx context.Context, recruiterID int64, page int64, size int64) (*models.ResponseListNotification, error) {
	offset := (page - 1) * size
	noti, total, err := n.NotificationGorm.GetListNotificationByRecruiter(ctx, recruiterID, offset, size)
	if err != nil {
		n.Logger.Error(err.Error(), zap.Int64("recruiter id", recruiterID))
		return nil, err
	}
	resp := models.ResponseListNotification{
		Total:         total,
		Notifications: noti,
	}
	return &resp, nil
}
