package skill

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/hieuxeko19991/job4e_be/services/autocomplete"
	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

type MockSkillGorm struct {
	mock.Mock
}

func (m *MockSkillGorm) Create(ctx context.Context, skill *models.Skill) error {
	args := m.Called(ctx, skill)
	return args.Error(0)
}

func (m *MockSkillGorm) Update(ctx context.Context, skill *models.RequestUpdateSkill, skillID int64) error {
	args := m.Called(ctx, skill, skillID)
	return args.Error(0)
}

func (m *MockSkillGorm) Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSkill, error) {
	args := m.Called(ctx, name, page, size)
	return args.Get(0).(*models.ResponsetListSkill), args.Error(1)
}

func (m *MockSkillGorm) GetAll(ctx context.Context, name string) ([]models.Skill, error) {
	args := m.Called(ctx, name)
	return args.Get(0).([]models.Skill), args.Error(1)
}

func TestNewSkill(t *testing.T) {
	skill := NewSkill(&SkillGorm{}, &autocomplete.Trie{}, &autocomplete.Trie{}, &autocomplete.Trie{}, zap.L())
	assert.NotNil(t, skill)
}

func TestSkill_CreateSkill(t *testing.T) {
	testcases := []struct {
		Name          string
		Req           *models.RequestCreateSkill
		MockObject    Skill
		ExpectedError error
	}{
		{
			Name: "happy case",
			Req: &models.RequestCreateSkill{
				Name:        "Go",
				Description: "ABC",
				Questions:   "ABC",
				Icon:        "ABC",
				Status:      true,
			},
			MockObject: Skill{
				SkillGorm:     &MockSkillGorm{},
				CandidateTrie: autocomplete.NewTrie(),
				RecruiterTrie: autocomplete.NewTrie(),
				JobTrie:       autocomplete.NewTrie(),
				Logger:        zap.L(),
			},
			ExpectedError: nil,
		},
	}
	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockSkillGorm)
			testObj.On("Create", context.Background(), &models.Skill{
				Name:        test.Req.Name,
				Description: test.Req.Description,
				Questions:   test.Req.Questions,
				Icon:        test.Req.Icon,
				Status:      test.Req.Status,
			}).Return(nil)
			test.MockObject.SkillGorm = testObj
			err := test.MockObject.CreateSkill(context.Background(), test.Req)
			assert.Nil(t, err)
		})
	}
}

func TestSkill_UpdateSkill(t *testing.T) {
	testcases := []struct {
		Name          string
		SkillIO       int64
		Req           *models.RequestCreateSkill
		MockObject    Skill
		ExpectedError error
	}{
		{
			Name:    "Happy case",
			SkillIO: 1,
			Req: &models.RequestCreateSkill{
				Name:        "Go",
				Description: "ABC",
				Questions:   "ABC",
				Icon:        "ABC",
				Status:      true,
			},
			MockObject: Skill{
				SkillGorm:     &MockSkillGorm{},
				CandidateTrie: autocomplete.NewTrie(),
				RecruiterTrie: autocomplete.NewTrie(),
				JobTrie:       autocomplete.NewTrie(),
				Logger:        zap.L(),
			},
			ExpectedError: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockSkillGorm)
			testObj.On("Update", context.Background(), &models.RequestUpdateSkill{
				Name:        test.Req.Name,
				Description: test.Req.Description,
				Questions:   test.Req.Questions,
				Icon:        test.Req.Icon,
				Status:      test.Req.Status,
			}, test.SkillIO).Return(nil)
			test.MockObject.SkillGorm = testObj
			err := test.MockObject.UpdateSkill(context.Background(), test.Req, test.SkillIO)
			assert.Nil(t, err)
		})
	}
}

func TestSkill_GetListSkill(t *testing.T) {
	testcases := []struct {
		Name          string
		name          string
		page          int64
		size          int64
		MockObject    Skill
		ExpectedRes   *models.ResponsetListSkill
		ExpectedError error
	}{
		{
			Name: "happy case",
			name: "Go",
			page: 1,
			size: 5,
			MockObject: Skill{
				SkillGorm:     &MockSkillGorm{},
				CandidateTrie: autocomplete.NewTrie(),
				RecruiterTrie: autocomplete.NewTrie(),
				JobTrie:       autocomplete.NewTrie(),
				Logger:        zap.L(),
			},
			ExpectedRes: &models.ResponsetListSkill{
				TotalSkill:  0,
				TotalPage:   0,
				CurrentPage: 0,
				Data:        []models.Skill{},
			},
			ExpectedError: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockSkillGorm)
			testObj.On("Get", context.Background(), test.name, test.page, test.size).Return(&models.ResponsetListSkill{
				TotalSkill:  0,
				TotalPage:   0,
				CurrentPage: 0,
				Data:        []models.Skill{},
			}, nil)
			test.MockObject.SkillGorm = testObj
			resp, err := test.MockObject.GetListSkill(context.Background(), test.name, test.page, test.size)
			assert.Equal(t, resp, test.ExpectedRes)
			assert.Nil(t, err)
		})
	}
}

func TestSkill_GetAll(t *testing.T) {
	testcases := []struct {
		Name        string
		SkillName   string
		MockObject  Skill
		ExpectedRes []models.Skill
		ExpectedErr error
	}{
		{
			Name:      "Happy case",
			SkillName: "Go",
			MockObject: Skill{
				SkillGorm:     &MockSkillGorm{},
				CandidateTrie: autocomplete.NewTrie(),
				RecruiterTrie: autocomplete.NewTrie(),
				JobTrie:       autocomplete.NewTrie(),
				Logger:        zap.L(),
			},
			ExpectedRes: []models.Skill{},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockSkillGorm)
			testObj.On("GetAll", context.Background(), test.SkillName).Return([]models.Skill{}, nil)
			test.MockObject.SkillGorm = testObj
			resp, err := test.MockObject.GetAll(context.Background(), test.SkillName)
			assert.Equal(t, resp, test.ExpectedRes)
			assert.Nil(t, err)
		})
	}
}
