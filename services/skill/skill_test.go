package skill

import (
	"testing"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

func TestSkill_CreateSkill(t *testing.T) {
	testcases := []struct {
		Name        string
		Req         *models.RequestCreateSkill
		MockFunc    func()
		ExpectedErr error
	}{
		{
			Name: "successful case",
			Req: &models.RequestCreateSkill{
				Name:        "Golang",
				Description: "description",
			},
		},
	}
}
