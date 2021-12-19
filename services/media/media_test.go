package media

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewMediaCategory(t *testing.T) {
	media := NewMediaCategory(&MediaGorm{}, zap.L())
	assert.NotNil(t, media)
}
