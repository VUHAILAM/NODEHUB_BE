package notification

import (
	"context"
	"testing"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/pkg/errors"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
)

type MockNotificationGorm struct {
	mock.Mock
}

func (g *MockNotificationGorm) Create(ctx context.Context, notification []*models.Notification) error {
	args := g.Called(ctx, notification)
	return args.Error(0)
}

func (g *MockNotificationGorm) GetListNotificationByCandidate(ctx context.Context, candidateID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
	args := g.Called(ctx, candidateID, offset, size)
	return args.Get(0).([]*models.Notification), int64(args.Int(1)), args.Error(2)
}

func (g *MockNotificationGorm) GetListNotificationByRecruiter(ctx context.Context, recruiterID int64, offset int64, size int64) ([]*models.Notification, int64, error) {
	args := g.Called(ctx, recruiterID, offset, size)
	return args.Get(0).([]*models.Notification), int64(args.Int(1)), args.Error(2)
}

func (g *MockNotificationGorm) UpdateCheckRead(ctx context.Context, notificationID int64) error {
	args := g.Called(ctx, notificationID)
	return args.Error(0)
}

func (g *MockNotificationGorm) UpdateCheckReadByAccountID(ctx context.Context, field string, accountID int64) error {
	args := g.Called(ctx, field, accountID)
	return args.Error(0)
}

func (g *MockNotificationGorm) CountUnread(ctx context.Context, field string, accountID int64) (int64, error) {
	args := g.Called(ctx, field, accountID)
	return int64(args.Int(0)), args.Error(1)
}

func TestNewNotification(t *testing.T) {
	noti := NewNotification(&NotificationGorm{}, zap.L())
	assert.NotNil(t, noti)
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
			Req: []*models.RequestCreateNotification{
				{
					RecruiterID: 1,
					CandidateID: 2,
				},
			},
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
			gMock.On("Create", context.Background(), []*models.Notification{
				{
					RecruiterID: 1,
					CandidateID: 2,
				},
			}).Return(nil)
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
				Notifications: []*models.Notification{},
			},
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			gMock := new(MockNotificationGorm)
			gMock.On("GetListNotificationByCandidate", context.Background(), test.CandidateID, test.Page-1, test.Size).Return([]*models.Notification{}, 0, nil)
			test.MockObj.NotificationGorm = gMock
			resp, err := test.MockObj.GetListNotificationByCandidate(context.Background(), test.CandidateID, test.Page, test.Size)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}

func TestNotificationService_GetListNotificationByRecruiter(t *testing.T) {
	testcases := []struct {
		Name        string
		RecruiterID int64
		Page        int64
		Size        int64
		MockObj     NotificationService
		ExpectedRes *models.ResponseListNotification
		ExpectedErr error
	}{
		{
			Name:        "Happy case",
			RecruiterID: 1,
			Page:        1,
			Size:        5,
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedErr: nil,
			ExpectedRes: &models.ResponseListNotification{
				Total:         0,
				Notifications: []*models.Notification{},
			},
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			gMock := new(MockNotificationGorm)
			gMock.On("GetListNotificationByRecruiter", context.Background(), test.RecruiterID, test.Page-1, test.Size).Return([]*models.Notification{}, 0, nil)
			test.MockObj.NotificationGorm = gMock
			resp, err := test.MockObj.GetListNotificationByRecruiter(context.Background(), test.RecruiterID, test.Page, test.Size)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}

func TestNotificationService_MarkRead(t *testing.T) {
	testcases := []struct {
		Name        string
		Req         models.RequestMarkRead
		MockObj     NotificationService
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			Req: models.RequestMarkRead{
				NotificationID: 1,
			},
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
			gMock.On("UpdateCheckRead", context.Background(), test.Req.NotificationID).Return(nil)
			test.MockObj.NotificationGorm = gMock
			err := test.MockObj.MarkRead(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestNotificationService_MarkReadAll(t *testing.T) {
	testcases := []struct {
		Name        string
		Req         models.RequestMarkReadAll
		MockObj     NotificationService
		ExpectedErr error
	}{
		{
			Name: "Happy case for candidate role",
			Req: models.RequestMarkReadAll{
				AccountID: 1,
				Role:      auth.CandidateRole,
			},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedErr: nil,
		},
		{
			Name: "Happy case for recruiter role",
			Req: models.RequestMarkReadAll{
				AccountID: 1,
				Role:      auth.RecruiterRole,
			},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedErr: nil,
		},
		{
			Name: "Error case",
			Req: models.RequestMarkReadAll{
				AccountID: 1,
				Role:      auth.CommonRole,
			},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedErr: errors.New("Role not found"),
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			gMock := new(MockNotificationGorm)
			if test.Req.Role == auth.CandidateRole {
				gMock.On("UpdateCheckReadByAccountID", context.Background(), "candidate_id", test.Req.AccountID).Return(nil)
			} else if test.Req.Role == auth.RecruiterRole {
				gMock.On("UpdateCheckReadByAccountID", context.Background(), "recruiter_id", test.Req.AccountID).Return(nil)
			}
			test.MockObj.NotificationGorm = gMock
			err := test.MockObj.MarkReadAll(context.Background(), test.Req)
			if err != nil {
				assert.Equal(t, test.ExpectedErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestNotificationService_CountUnread(t *testing.T) {
	testcases := []struct {
		Name          string
		Req           models.RequestCountUnread
		MockObj       NotificationService
		ExpectedCount int64
		ExpectedErr   error
	}{
		{
			Name: "Happy case for candidate role",
			Req: models.RequestCountUnread{
				AccountID: 1,
				Role:      auth.CandidateRole,
			},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedCount: 10,
			ExpectedErr:   nil,
		},
		{
			Name: "Happy case for recruiter role",
			Req: models.RequestCountUnread{
				AccountID: 1,
				Role:      auth.RecruiterRole,
			},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedCount: 10,
			ExpectedErr:   nil,
		},
		{
			Name: "Error case",
			Req: models.RequestCountUnread{
				AccountID: 1,
				Role:      auth.CommonRole,
			},
			MockObj: NotificationService{
				NotificationGorm: &MockNotificationGorm{},
				Logger:           zap.L(),
			},
			ExpectedCount: 10,
			ExpectedErr:   errors.New("Role not found"),
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			gMock := new(MockNotificationGorm)
			if test.Req.Role == auth.CandidateRole {
				gMock.On("CountUnread", context.Background(), "candidate_id", test.Req.AccountID).Return(10, nil)
			} else if test.Req.Role == auth.RecruiterRole {
				gMock.On("CountUnread", context.Background(), "recruiter_id", test.Req.AccountID).Return(10, nil)
			}
			test.MockObj.NotificationGorm = gMock
			count, err := test.MockObj.CountUnread(context.Background(), test.Req)
			if err != nil {
				assert.Equal(t, test.ExpectedErr.Error(), err.Error())
			} else {
				assert.Equal(t, test.ExpectedCount, count)
				assert.Nil(t, err)
			}
		})
	}
}
