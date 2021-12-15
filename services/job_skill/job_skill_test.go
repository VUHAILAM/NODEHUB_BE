package job_skill

import (
	"context"
	"testing"

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
		{},
	}
}
