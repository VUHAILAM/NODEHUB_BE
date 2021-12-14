package notification

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

type MockNotificationGorm struct {
	mock.Mock
}

func (g *MockNotificationGorm) Create(ctx context.Context, notification []*models.Notification) error {
	return nil
}

func (g *MockNotificationGorm) GetListNotificationByCandidate(ctx context.Context, candidateID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
	return nil, 0, nil
}

func (g *MockNotificationGorm) GetListNotificationByRecruiter(ctx context.Context, recruiterID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
	return nil, 0, nil
}

func (g *MockNotificationGorm) UpdateCheckRead(ctx context.Context, notificationID int64) error {
	return nil
}

func (g *MockNotificationGorm) UpdateCheckReadByAccountID(ctx context.Context, field string, accountID int64) error {
	return nil
}

func TestNotificationService_CreateNotification(t *testing.T) {
	testcases := []struct {
		Name        string
		Req         []*models.RequestCreateNotification
		MockObj     NotificationService
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			Req:  []*models.RequestCreateNotification{},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			gMock := new(MockNotificationGorm)
			gMock.On("Create", context.Background(), []*models.Notification{}).Return(nil)
			test.MockObj.NotificationGorm = gMock
			err := test.MockObj.CreateNotification(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestNotificationService_GetListNotificationByCandidate(t *testing.T) {
	testcases := []struct {
		Name        string
		CandidateID int64
		Page        int64
		Size        int64
		MockObj     NotificationService
		ExpectedRes *models.ResponseListNotification
		ExpectedErr error
	}{
		{
			Name:        "Happy case",
			CandidateID: 1,
			Page:        1,
			Size:        5,
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedErr: nil,
			ExpectedRes: &models.ResponseListNotification{
				Total:         0,
				Notifications: []*models.Notification{nil},
			},
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			gMock := new(MockNotificationGorm)
			gMock.On("GetListNotificationByCandidate", context.Background(), test.CandidateID, test.Page, test.Size).Return([]*models.Notification{}, 0, nil)
			test.MockObj.NotificationGorm = gMock
			resp, err := test.MockObj.GetListNotificationByCandidate(context.Background(), test.CandidateID, test.Page, test.Size)
			assert.Equal(t, resp, test.ExpectedRes)
			assert.Nil(t, err)
		})
	}
}
