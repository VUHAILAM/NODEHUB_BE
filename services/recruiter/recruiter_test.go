package recruiter

import (
	"context"
	"testing"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
)

type MockRecruiterGorm struct {
	mock.Mock
}

func (g *MockRecruiterGorm) Create(ctx context.Context, recruiter *models2.Recruiter) (int64, error) {
	args := g.Called(ctx, recruiter)
	return int64(args.Int(0)), args.Error(1)
}

func (g *MockRecruiterGorm) GetAllRecruiterName(ctx context.Context) ([]string, error) {
	args := g.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}

func (g *MockRecruiterGorm) AddRecruiterSkill(ctx context.Context, recruiterSkill *models2.RecruiterSkill) error {
	args := g.Called(ctx, recruiterSkill)
	return args.Error(0)
}

func (g *MockRecruiterGorm) UpdateProfile(ctx context.Context, recruiter *models2.RequestUpdateRecruiter, recruiter_id int64) error {
	args := g.Called(ctx, recruiter, recruiter_id)
	return args.Error(0)
}

func (g *MockRecruiterGorm) GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models2.ResponseRecruiterSkill, error) {
	args := g.Called(ctx, recruiter_id)
	return args.Get(0).([]models2.ResponseRecruiterSkill), args.Error(1)
}

func (g *MockRecruiterGorm) GetProfile(ctx context.Context, id int64) (*models2.Recruiter, error) {
	args := g.Called(ctx, id)
	return args.Get(0).(*models2.Recruiter), args.Error(1)
}

func (g *MockRecruiterGorm) GetAllRecruiterForAdmin(ctx context.Context, name string, page int64, size int64) (*models2.ResponsetListRecruiter, error) {
	args := g.Called(ctx, name, page, size)
	return args.Get(0).(*models2.ResponsetListRecruiter), args.Error(1)
}

func (g *MockRecruiterGorm) UpdateRecruiterByAdmin(ctx context.Context, recruiter_id int64, data map[string]interface{}) error {
	args := g.Called(ctx, recruiter_id, data)
	return args.Error(0)
}

func (g *MockRecruiterGorm) UpdateStatusRecruiter(ctx context.Context, updateRequest *models2.RequestUpdateStatusRecruiter, recruiter_id int64) error {
	args := g.Called(ctx, updateRequest, recruiter_id)
	return args.Error(0)
}

func (g *MockRecruiterGorm) GetAllRecruiterForCandidate(ctx context.Context, recruiterName string, skillName string, address string, page int64, size int64) (*models2.ResponsetListRecruiterForCandidate, error) {
	args := g.Called(ctx, recruiterName, skillName, address, page, size)
	return args.Get(0).(*models2.ResponsetListRecruiterForCandidate), args.Error(1)
}

func (g *MockRecruiterGorm) DeleteRecruiterSkill(ctx context.Context, recruiter_skill_id int64) error {
	args := g.Called(ctx, recruiter_skill_id)
	return args.Error(0)
}

func (g *MockRecruiterGorm) SearchRecruiter(ctx context.Context, text string, offset, size int64) ([]*models2.Recruiter, int64, error) {
	args := g.Called(ctx, text, offset, size)
	return args.Get(0).([]*models2.Recruiter), int64(args.Int(1)), args.Error(2)
}

func (g *MockRecruiterGorm) GetAllRecruiter(ctx context.Context, offset, size int64) ([]*models2.Recruiter, int64, error) {
	args := g.Called(ctx, offset, size)
	return args.Get(0).([]*models2.Recruiter), int64(args.Int(1)), args.Error(2)
}

func (g *MockRecruiterGorm) GetAllSkillByRecruiterID(ctx context.Context, recruiterID int64) ([]*models2.Skill, error) {
	args := g.Called(ctx, recruiterID)
	return args.Get(0).([]*models2.Skill), args.Error(1)
}

func (g *MockRecruiterGorm) Count(ctx context.Context) (int64, error) {
	args := g.Called(ctx)
	return int64(args.Int(0)), args.Error(1)
}

func (g *MockRecruiterGorm) GetPremiumField(ctx context.Context, recruiterID int64) (bool, error) {
	args := g.Called(ctx, recruiterID)
	return args.Bool(0), args.Error(1)
}

type MockEmailService struct {
	mock.Mock
}

func (s *MockEmailService) CreateMail(mailReq *models2.Mail) []byte {
	args := s.Called(mailReq)
	return args.Get(0).([]byte)
}

func (s *MockEmailService) SendMail(mailReq *models2.Mail) error {
	args := s.Called(mailReq)
	return args.Error(0)
}

func (s *MockEmailService) NewMail(from string, to []string, subject string, mailType models2.MailType, data *models2.MailData) *models2.Mail {
	return &models2.Mail{
		From:    from,
		To:      to,
		Subject: subject,
		Mtype:   mailType,
		Data:    data,
	}
}

func TestNewRecruiterCategory(t *testing.T) {
	recruiter := NewRecruiterCategory(&RecruiterGorm{}, nil, zap.L())
	assert.NotNil(t, recruiter)
}

func TestRecruiter_AddRecruiterSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		Req         *models2.RecruiterSkill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			Req: &models2.RecruiterSkill{
				RecruiterId: 1,
				SkillId:     2,
			},
			ExpectedErr: nil,
		},
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			Req: &models2.RecruiterSkill{
				RecruiterId: 1,
				SkillId:     2,
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("AddRecruiterSkill", context.Background(), test.Req).Return(nil)
			test.TestObj.RecruiterGorm = mockObj
			err := test.TestObj.AddRecruiterSkill(context.Background(), test.Req)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_UpdateProfile(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		Req         *models2.RequestUpdateRecruiter
		RecruiterID int64
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			Req: &models2.RequestUpdateRecruiter{
				Name: "One mount",
			},
			RecruiterID: 1,
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("UpdateProfile", context.Background(), test.Req, test.RecruiterID).Return(nil)
			test.TestObj.RecruiterGorm = mockObj
			err := test.TestObj.UpdateProfile(context.Background(), test.Req, test.RecruiterID)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_GetProfileRecruiter(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		ID          int64
		ExpectedRes *models2.Recruiter
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			ID: 1,
			ExpectedRes: &models2.Recruiter{
				RecruiterID: 1,
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("GetProfile", context.Background(), test.ID).Return(&models2.Recruiter{
				RecruiterID: 1,
			}, nil)
			test.TestObj.RecruiterGorm = mockObj
			res, err := test.TestObj.GetProfileRecruiter(context.Background(), test.ID)
			assert.Equal(t, test.ExpectedRes, res)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_GetRecruiterSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		ID          int64
		ExpectedRes []models2.ResponseRecruiterSkill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			ID: 1,
			ExpectedRes: []models2.ResponseRecruiterSkill{
				{
					Id: 1,
				},
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("GetRecruiterSkill", context.Background(), test.ID).Return([]models2.ResponseRecruiterSkill{
				{
					Id: 1,
				},
			}, nil)
			test.TestObj.RecruiterGorm = mockObj
			res, err := test.TestObj.GetRecruiterSkill(context.Background(), test.ID)
			assert.Equal(t, test.ExpectedRes, res)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_DeleteRecruiterSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		SkillID     int64
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			SkillID:     1,
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("DeleteRecruiterSkill", context.Background(), test.SkillID).Return(nil)
			test.TestObj.RecruiterGorm = mockObj
			err := test.TestObj.DeleteRecruiterSkill(context.Background(), test.SkillID)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_GetAllRecruiterForAdmin(t *testing.T) {
	testcases := []struct {
		TestName    string
		TestObj     Recruiter
		Name        string
		Page        int64
		Size        int64
		ExpectedRes *models2.ResponsetListRecruiter
		ExpectedErr error
	}{
		{
			TestName: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			Name: "One mount",
			Page: 1,
			Size: 5,
			ExpectedRes: &models2.ResponsetListRecruiter{
				Total:       1,
				TotalPage:   1,
				CurrentPage: 1,
				Data: []models2.RecruiterForAdmin{
					{
						RecruiterID: 1,
						Name:        "One mount",
					},
				},
			},
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.TestName, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("GetAllRecruiterForAdmin", context.Background(), test.Name, test.Page, test.Size).
				Return(&models2.ResponsetListRecruiter{
					Total:       1,
					TotalPage:   1,
					CurrentPage: 1,
					Data: []models2.RecruiterForAdmin{
						{
							RecruiterID: 1,
							Name:        "One mount",
						},
					},
				}, nil)
			test.TestObj.RecruiterGorm = mockObj
			resp, err := test.TestObj.GetAllRecruiterForAdmin(context.Background(), test.Name, test.Page, test.Size)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_UpdateReciuterByAdmin(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		Req         *models2.RequestUpdateRecruiterAdmin
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			Req: &models2.RequestUpdateRecruiterAdmin{
				RecruiterID:    1,
				Nodehub_review: "Good",
				Premium:        true,
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			updateData := map[string]interface{}{}
			mapStructureDecodeWithTextUnmarshaler(test.Req, &updateData)
			mockObj.On("UpdateRecruiterByAdmin", context.Background(), test.Req.RecruiterID, updateData).Return(nil)
			test.TestObj.RecruiterGorm = mockObj
			err := test.TestObj.UpdateReciuterByAdmin(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestRecruiter_UpdateStatusReciuter(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		RecruiterID int64
		Req         *models2.RequestUpdateStatusRecruiter
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Email:         &MockEmailService{},

				Conf:   &config.Config{},
				Logger: zap.L(),
			},
			RecruiterID: 1,
			Req: &models2.RequestUpdateStatusRecruiter{
				Status: true,
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockEmail := new(MockEmailService)
			mockObj.On("UpdateStatusRecruiter", context.Background(), test.Req, test.RecruiterID).Return(nil)
			mockObj.On("GetProfile", context.Background(), test.RecruiterID).Return(&models2.Recruiter{
				RecruiterID: test.RecruiterID,
				Email:       "abc@def.com",
			}, nil)
			mockEmail.On("SendMail", &models2.Mail{
				From:    "lamvhhe130764@fpt.edu.vn",
				Subject: "Approved your Company on NodeHub",
				Mtype:   3,
				To:      []string{"abc@def.com"},
				Data: &models2.MailData{
					Link: test.TestObj.Conf.Domain + "recruiter/login",
				},
			}).Return(nil)

			test.TestObj.RecruiterGorm = mockObj
			test.TestObj.Email = mockEmail
			err := test.TestObj.UpdateStatusReciuter(context.Background(), test.Req, test.RecruiterID)
			assert.Nil(t, err)
		})
	}
}

func TestRecruiter_GetAllRecruiterForCandidate(t *testing.T) {
	testcases := []struct {
		Name          string
		TestObj       Recruiter
		RecruiterName string
		SkillName     string
		Address       string
		Page          int64
		Size          int64
		ExpectedRes   *models2.ResponsetListRecruiterForCandidate
		ExpectedErr   error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			RecruiterName: "One mount",
			SkillName:     "Golang",
			Address:       "Hanoi",
			Page:          1,
			Size:          5,
			ExpectedRes: &models2.ResponsetListRecruiterForCandidate{
				Total:       1,
				TotalPage:   1,
				CurrentPage: 1,
				Data: []models2.RecruiterForCandidateCheck{
					{
						RecruiterID: 1,
						Name:        "One mount",
						Skill_name:  "Golang",
						Address:     "Hanoi",
					},
				},
			},
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("GetAllRecruiterForCandidate", context.Background(), test.RecruiterName, test.SkillName, test.Address, test.Page, test.Size).
				Return(&models2.ResponsetListRecruiterForCandidate{
					Total:       1,
					TotalPage:   1,
					CurrentPage: 1,
					Data: []models2.RecruiterForCandidateCheck{
						{
							RecruiterID: 1,
							Name:        "One mount",
							Skill_name:  "Golang",
							Address:     "Hanoi",
						},
					},
				}, nil)
			test.TestObj.RecruiterGorm = mockObj
			resp, err := test.TestObj.GetAllRecruiterForCandidate(context.Background(), test.RecruiterName, test.SkillName, test.Address, test.Page, test.Size)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}

func TestRecruiter_CountRecruiter(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		ExpectedRes int64
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			ExpectedRes: 10,
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("Count", context.Background()).Return(10, nil)
			test.TestObj.RecruiterGorm = mockObj
			count, err := test.TestObj.CountRecruiter(context.Background())
			assert.Equal(t, test.ExpectedRes, count)
			assert.Nil(t, err)
		})
	}
}

func TestRecruiter_CheckPremium(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     Recruiter
		RecruiterID int64
		ExpectedRes bool
		ExpectedErr error
	}{
		{
			Name: "happy case",
			TestObj: Recruiter{
				RecruiterGorm: &MockRecruiterGorm{},
				Logger:        zap.L(),
			},
			RecruiterID: 1,
			ExpectedRes: true,
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockRecruiterGorm)
			mockObj.On("GetPremiumField", context.Background(), test.RecruiterID).Return(true, nil)
			test.TestObj.RecruiterGorm = mockObj
			res, err := test.TestObj.CheckPremium(context.Background(), test.RecruiterID)
			assert.Equal(t, test.ExpectedRes, res)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
