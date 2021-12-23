package candidate

import (
	"context"
	"testing"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
)

type MockCandidateGorm struct {
	mock.Mock
}

func (g *MockCandidateGorm) Create(ctx context.Context, createData *models2.Candidate) (int64, error) {
	args := g.Called(ctx, createData)
	return int64(args.Int(0)), args.Error(1)
}

func (g *MockCandidateGorm) Update(ctx context.Context, candidateID int64, updateData *models2.Candidate) error {
	args := g.Called(ctx, candidateID, updateData)
	return args.Error(0)
}
func (g *MockCandidateGorm) GetByCandidateID(ctx context.Context, candidateID int64) (*models2.Candidate, error) {
	args := g.Called(ctx, candidateID)
	return args.Get(0).(*models2.Candidate), args.Error(1)
}
func (g *MockCandidateGorm) GetAllName(ctx context.Context) ([]string, error) {
	args := g.Called(ctx)
	return args.Get(0).([]string), args.Error(1)
}
func (g *MockCandidateGorm) GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models2.ResponsetListCandidateAdmin, error) {
	args := g.Called(ctx, name, page, size)
	return args.Get(0).(*models2.ResponsetListCandidateAdmin), args.Error(1)
}
func (g *MockCandidateGorm) UpdateReviewCandidateByAdmin(ctx context.Context, candidate_id int64, data map[string]interface{}) error {
	args := g.Called(ctx, candidate_id, data)
	return args.Error(0)
}
func (g *MockCandidateGorm) UpdateStatusCandidate(ctx context.Context, candidate *models2.RequestUpdateStatusCandidate, candidate_id int64) error {
	args := g.Called(ctx, candidate, candidate_id)
	return args.Error(0)
}
func (g *MockCandidateGorm) AddCandidateSkill(ctx context.Context, candidateSkill *models2.CandidateSkill) error {
	args := g.Called(ctx, candidateSkill)
	return args.Error(0)
}
func (g *MockCandidateGorm) DeleteCandidateSkill(ctx context.Context, candidate_skill_id int64) error {
	args := g.Called(ctx, candidate_skill_id)
	return args.Error(0)
}
func (g *MockCandidateGorm) UpdateCandidateSkill(ctx context.Context, candidate_skill_id int64, data map[string]interface{}) error {
	args := g.Called(ctx, candidate_skill_id, data)
	return args.Error(0)
}
func (g *MockCandidateGorm) GetCandidateSkill(ctx context.Context, candidate_id int64) ([]models2.ResponseCandidateSkill, error) {
	args := g.Called(ctx, candidate_id)
	return args.Get(0).([]models2.ResponseCandidateSkill), args.Error(1)
}
func (g *MockCandidateGorm) SearchCandidate(ctx context.Context, text string, offset, page int64) ([]*models2.Candidate, int64, error) {
	args := g.Called(ctx, text, offset, page)
	return args.Get(0).([]*models2.Candidate), int64(args.Int(1)), args.Error(2)
}
func (g *MockCandidateGorm) GetAllCandidate(ctx context.Context, offset, size int64) ([]*models2.Candidate, int64, error) {
	args := g.Called(ctx, offset, size)
	return args.Get(0).([]*models2.Candidate), int64(args.Int(1)), args.Error(2)
}
func (g *MockCandidateGorm) GetAllSkillByCandidateID(ctx context.Context, candidateID int64) ([]*models2.Skill, error) {
	args := g.Called(ctx, candidateID)
	return args.Get(0).([]*models2.Skill), args.Error(1)
}
func (g *MockCandidateGorm) Count(ctx context.Context) (int64, error) {
	args := g.Called(ctx)
	return int64(args.Int(0)), args.Error(1)
}

func TestNewCandidateService(t *testing.T) {
	candi := NewCandidateService(&CandidateGorm{}, zap.L())
	assert.NotNil(t, candi)
}

func TestCandidateService_CreateCandidateProfile(t *testing.T) {
	testcases := []struct {
		Name        string
		Req         models2.CandidateRequest
		TestObj     CandidateService
		ExpectedRes int64
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			Req: models2.CandidateRequest{
				CandidateID: 1,
			},
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			ExpectedRes: 1,
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			can, _ := test.Req.ToCandidate()
			mockObj.On("Create", context.Background(), &can).Return(1, nil)
			test.TestObj.CanGorm = mockObj
			res, err := test.TestObj.CreateCandidateProfile(context.Background(), test.Req)
			assert.Equal(t, test.ExpectedRes, res)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_UpdateCandidateProfile(t *testing.T) {
	testcases := []struct {
		Name        string
		Req         models2.CandidateRequest
		TestObj     CandidateService
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			Req: models2.CandidateRequest{
				CandidateID: 1,
				FirstName:   "Hai Lam",
			},
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			can, _ := test.Req.ToCandidate()
			mockObj.On("Update", context.Background(), test.Req.CandidateID, &can).Return(nil)
			test.TestObj.CanGorm = mockObj
			err := test.TestObj.UpdateCandidateProfile(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_GetCandidateProfile(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		CandidateID int64
		ExpectedRes *models2.CandidateResponse
		ExpectedErr error
	}{
		{
			Name: "happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			CandidateID: 1,
			ExpectedRes: &models2.CandidateResponse{
				CandidateID:       1,
				LastName:          "Hai Lam",
				CVManage:          []models2.CV{},
				ExperienceManage:  []models2.Experience{},
				EducationManage:   []models2.Education{},
				SocialManage:      []models2.Social{},
				ProjectManage:     []models2.Project{},
				CertificateManage: []models2.Certificate{},
				PrizeManage:       []models2.Prize{},
			},
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)

			mockObj.On("GetByCandidateID", context.Background(), test.CandidateID).
				Return(&models2.Candidate{
					CandidateID:       1,
					LastName:          "Hai Lam",
					CvManage:          "[]",
					ExperienceManage:  "[]",
					EducationManage:   "[]",
					SocialManage:      "[]",
					ProjectManage:     "[]",
					CertificateManage: "[]",
					PrizeManage:       "[]",
				}, nil)

			test.TestObj.CanGorm = mockObj
			resp, err := test.TestObj.GetCandidateProfile(context.Background(), test.CandidateID)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_GetAllCandidateForAdmin(t *testing.T) {
	testcases := []struct {
		TestName    string
		TestObj     CandidateService
		Name        string
		Page        int64
		Size        int64
		ExpectedRes *models2.ResponsetListCandidateAdmin
		ExpectedErr error
	}{
		{
			TestName: "Happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			Name: "Lam",
			Page: 1,
			Size: 5,
			ExpectedRes: &models2.ResponsetListCandidateAdmin{
				Total:       0,
				TotalPage:   1,
				CurrentPage: 1,
				Data:        []models2.CandidateRequestAdmin{},
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.TestName, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			mockObj.On("GetAllCandidateForAdmin", context.Background(), test.Name, test.Page, test.Size).Return(&models2.ResponsetListCandidateAdmin{
				Total:       0,
				TotalPage:   1,
				CurrentPage: 1,
				Data:        []models2.CandidateRequestAdmin{},
			}, nil)
			test.TestObj.CanGorm = mockObj
			resp, err := test.TestObj.GetAllCandidateForAdmin(context.Background(), test.Name, test.Page, test.Size)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_UpdateReviewCandidateByAdmin(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		Req         *models2.RequestUpdateReviewCandidateAdmin
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			Req: &models2.RequestUpdateReviewCandidateAdmin{
				CandidateID:    1,
				Nodehub_review: "Good",
				NodehubScore:   5,
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			updateData := map[string]interface{}{}
			mapStructureDecodeWithTextUnmarshaler(test.Req, &updateData)
			mockObj.On("UpdateReviewCandidateByAdmin", context.Background(), test.Req.CandidateID, updateData).Return(nil)
			test.TestObj.CanGorm = mockObj
			err := test.TestObj.UpdateReviewCandidateByAdmin(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_UpdateStatusCandidate(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		CandidateId int64
		Req         *models2.RequestUpdateStatusCandidate
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			CandidateId: 1,
			Req: &models2.RequestUpdateStatusCandidate{
				Status: true,
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			mockObj.On("UpdateStatusCandidate", context.Background(), test.Req, test.CandidateId).Return(nil)
			test.TestObj.CanGorm = mockObj
			err := test.TestObj.UpdateStatusCandidate(context.Background(), test.Req, test.CandidateId)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_AddCandidateSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		Req         *models2.CandidateSkill
		ExpectedErr error
	}{
		{
			Name: "happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			Req: &models2.CandidateSkill{
				Id:          1,
				CandidateId: 1,
				SkillId:     1,
				Media:       "abc",
			},
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			mockObj.On("AddCandidateSkill", context.Background(), test.Req).Return(nil)
			test.TestObj.CanGorm = mockObj
			err := test.TestObj.AddCandidateSkill(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_DeleteCandidateSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		SkillID     int64
		ExpectedErr error
	}{
		{
			Name: "happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			SkillID:     1,
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			mockObj.On("DeleteCandidateSkill", context.Background(), test.SkillID).Return(nil)
			test.TestObj.CanGorm = mockObj
			err := test.TestObj.DeleteCandidateSkill(context.Background(), test.SkillID)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_UpdateCandidateSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		Req         *models2.RequestUpdateCandidateSkill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			Req: &models2.RequestUpdateCandidateSkill{
				ID:    1,
				Media: "def",
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			updateData := map[string]interface{}{}
			mapStructureDecodeWithTextUnmarshaler(test.Req, &updateData)
			mockObj.On("UpdateCandidateSkill", context.Background(), test.Req.ID, updateData).Return(nil)
			test.TestObj.CanGorm = mockObj
			err := test.TestObj.UpdateCandidateSkill(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_GetCandidateSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		CandidateId int64
		ExpectedRes []models2.ResponseCandidateSkill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			CandidateId: 1,
			ExpectedRes: []models2.ResponseCandidateSkill{},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			mockObj.On("GetCandidateSkill", context.Background(), test.CandidateId).Return([]models2.ResponseCandidateSkill{}, nil)
			test.TestObj.CanGorm = mockObj
			resp, err := test.TestObj.GetCandidateSkill(context.Background(), test.CandidateId)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}

func TestCandidateService_CountCandidate(t *testing.T) {
	testcases := []struct {
		Name        string
		TestObj     CandidateService
		ExpectedRes int64
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			TestObj: CandidateService{
				CanGorm: &MockCandidateGorm{},
				Logger:  zap.L(),
			},
			ExpectedRes: 10,
			ExpectedErr: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			mockObj := new(MockCandidateGorm)
			mockObj.On("Count", context.Background()).Return(10, nil)
			test.TestObj.CanGorm = mockObj
			count, err := test.TestObj.CountCandidate(context.Background())
			assert.Equal(t, test.ExpectedRes, count)
			assert.Nil(t, err)
		})
	}
}
