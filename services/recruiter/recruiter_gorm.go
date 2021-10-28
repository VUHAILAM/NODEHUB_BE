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
	Create(ctx context.Context, recruiter *models.Recruiter) (int64, error)
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

func (r *RecruiterGorm) Create(ctx context.Context, recruiter *models.Recruiter) (int64, error) {
	db := r.db.WithContext(ctx)
	err := db.Table(tableRecruiter).Create(recruiter).Error
	if err != nil {
		r.logger.Error("AccountGorm: Create recruiter error", zap.Error(err))
		return 0, err
	}
	return recruiter.RecruiterID, nil
}
