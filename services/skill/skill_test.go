package skill

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

type MockSkillGorm struct {
	mock.Mock
}

func (m *MockSkillGorm) Create(ctx context.Context, skill *models.Skill) error {
	args := m.Called(ctx, skill)

}

func TestSkill_CreateSkill(t *testing.T) {
	testcases := []struct {
		Name          string
		Req           *models.RequestCreateSkill
		MockObject    Skill
		ExpectedError error
	}{
		{},
	}

}
