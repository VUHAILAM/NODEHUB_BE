package category

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewCategory(t *testing.T) {
	cate := NewCategory(&CategoryGorm{}, zap.L())
	assert.NotNil(t, cate)
}
