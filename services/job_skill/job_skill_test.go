package job_skill

import (
	"context"
	"testing"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
)

type MockJobSkillGorm struct {
	mock.Mock
}

func (g *MockJobSkillGorm) Create(ctx context.Context, jobSkill []*models2.JobSkill) error {
	args := g.Called(ctx, jobSkill)
	return args.Error(0)
}

func (g *MockJobSkillGorm) GetAllSkillByJob(ctx context.Context, jobID int64) ([]*models2.Skill, error) {
	arg := g.Called(ctx, jobID)
	return arg.Get(0).([]*models2.Skill), arg.Error(1)
}

func (g *MockJobSkillGorm) GetAllJobBySkill(ctx context.Context, skillID, offset, size int64) ([]*models2.Job, int64, error) {
	args := g.Called(ctx, skillID, offset, size)
	return args.Get(0).([]*models2.Job), int64(args.Int(1)), args.Error(2)
}

func (g *MockJobSkillGorm) Delete(ctx context.Context, jobID int64) error {
	args := g.Called(ctx, jobID)
	return args.Error(0)
}

func TestNewJobSkill(t *testing.T) {
	jobSkill := NewJobSkill(&JobSkillGorm{}, zap.L())
	assert.NotNil(t, jobSkill)
}

func TestJobSkill_GetJobBySkill(t *testing.T) {
	testcases := []struct {
		Name        string
		MockObject  JobSkill
		Req         models2.RequestGetJobsBySkill
		ExpectedRes *models2.ResponseGetJobsBySkill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			MockObject: JobSkill{
				JSGorm: &MockJobSkillGorm{},
				Logger: zap.L(),
			},
			Req: models2.RequestGetJobsBySkill{
				SkillID: 1,
				Page:    1,
				Size:    5,
			},
			ExpectedRes: &models2.ResponseGetJobsBySkill{
				Total:  0,
				Result: []*models2.Job{},
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockJobSkillGorm)
			testObj.On("GetAllJobBySkill", context.Background(), test.Req.SkillID, (test.Req.Page-1)*test.Req.Size, test.Req.Size).
				Return([]*models2.Job{}, 0, nil)
			test.MockObject.JSGorm = testObj
			resp, err := test.MockObject.GetJobBySkill(context.Background(), test.Req)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}

func TestJobSkill_GetSkillsByJob(t *testing.T) {
	testcases := []struct {
		Name        string
		MockObject  JobSkill
		Req         models2.RequestGetSkillsByJob
		ExpectedRes []*models2.Skill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			MockObject: JobSkill{
				JSGorm: &MockJobSkillGorm{},
				Logger: zap.L(),
			},
			Req: models2.RequestGetSkillsByJob{
				JobID: 1,
			},
			ExpectedRes: []*models2.Skill{},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockJobSkillGorm)
			testObj.On("GetAllSkillByJob", context.Background(), test.Req.JobID).
				Return([]*models2.Skill{}, nil)
			test.MockObject.JSGorm = testObj
			resp, err := test.MockObject.GetSkillsByJob(context.Background(), test.Req)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}
