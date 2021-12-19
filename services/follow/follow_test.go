package follow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/hieuxeko19991/job4e_be/services/candidate"
	"gitlab.com/hieuxeko19991/job4e_be/services/notification"
	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"
	"go.uber.org/zap"
)

func TestNewFollowService(t *testing.T) {
	follow := NewFollowService(&FollowGorm{}, &notification.NotificationGorm{}, &candidate.CandidateGorm{}, &recruiter.RecruiterGorm{}, zap.L())
	assert.NotNil(t, follow)
}
