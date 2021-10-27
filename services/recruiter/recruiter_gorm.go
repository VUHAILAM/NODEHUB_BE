package recruiter

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

const (
	tableRecruiter = "recruiter"
)

type IRecruiterDatabase interface {
	Create(ctx context.Context, recruiter models.Recruiter) (int64, error)
}

type RecruiterGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewRecruiterGorm(db *gorm.DB, logger *zap.Logger) *RecruiterGorm {
	return &RecruiterGorm{
		db:     db,
		logger: logger,
	}
}
