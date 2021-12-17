package job_skill

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"

	"github.com/stretchr/testify/mock"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

type MockJobSkillGorm struct {
	mock.Mock
}

func (g *MockJobSkillGorm) Create(ctx context.Context, jobSkill []*models.JobSkill) error {
	args := g.Called(ctx, jobSkill)
	return args.Error(0)
}

func (g *MockJobSkillGorm) GetAllSkillByJob(ctx context.Context, jobID int64) ([]*models.Skill, error) {
	arg := g.Called(ctx, jobID)
	return arg.Get(0).([]*models.Skill), arg.Error(1)
}

func (g *MockJobSkillGorm) GetAllJobBySkill(ctx context.Context, skillID, offset, size int64) ([]*models.Job, int64, error) {
	args := g.Called(ctx, skillID, offset, size)
	return args.Get(0).([]*models.Job), int64(args.Int(1)), args.Error(2)
}

func (g *MockJobSkillGorm) Delete(ctx context.Context, jobID int64) error {
	args := g.Called(ctx, jobID)
	return args.Error(0)
}

func TestJobSkill_GetJobBySkill(t *testing.T) {
	testcases := []struct {
		Name        string
		MockObject  JobSkill
		Req         models.RequestGetJobsBySkill
		ExpectedRes *models.ResponseGetJobsBySkill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			MockObject: JobSkill{
				JSGorm: &MockJobSkillGorm{},
				Logger: zap.L(),
			},
			Req: models.RequestGetJobsBySkill{
				SkillID: 1,
				Page:    1,
				Size:    5,
			},
			ExpectedRes: &models.ResponseGetJobsBySkill{
				Total:  0,
				Result: []*models.Job{},
			},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockJobSkillGorm)
			testObj.On("GetAllJobBySkill", context.Background(), test.Req.SkillID, (test.Req.Page-1)*test.Req.Size, test.Req.Size).
				Return([]*models.Job{}, 0, nil)
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
		Req         models.RequestGetSkillsByJob
		ExpectedRes []*models.Skill
		ExpectedErr error
	}{
		{
			Name: "Happy case",
			MockObject: JobSkill{
				JSGorm: &MockJobSkillGorm{},
				Logger: zap.L(),
			},
			Req: models.RequestGetSkillsByJob{
				JobID: 1,
			},
			ExpectedRes: []*models.Skill{},
			ExpectedErr: nil,
		},
	}

	for _, test := range testcases {
		t.Run(test.Name, func(t *testing.T) {
			testObj := new(MockJobSkillGorm)
			testObj.On("GetAllSkillByJob", context.Background(), test.Req.JobID).
				Return([]*models.Skill{}, nil)
			test.MockObject.JSGorm = testObj
			resp, err := test.MockObject.GetSkillsByJob(context.Background(), test.Req)
			assert.Equal(t, test.ExpectedRes, resp)
			assert.Nil(t, err)
		})
	}
}
