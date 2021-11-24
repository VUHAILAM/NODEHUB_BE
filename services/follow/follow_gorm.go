package follow

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	followTable = "follow"
)

type FollowGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewFollowGorm(db *gorm.DB, logger *zap.Logger) *FollowGorm {
	return &FollowGorm{
		DB:     db,
		Logger: logger,
	}
}

func (fg *FollowGorm) Create(ctx context.Context, follow *models.Follow) error {
	db := fg.DB.WithContext(ctx)
	err := db.Table(followTable).Create(follow).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (fg *FollowGorm) Delete(ctx context.Context, follow *models.Follow) error {
	db := fg.DB.WithContext(ctx)
	err := db.Table(followTable).Where("candidate_id=? and recruiter_id=?", follow.CandidateID, follow.RecruiterID).Delete(follow).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (fg *FollowGorm) GetFollow(ctx context.Context, candidateID int64, recruiterID int64) (*models.Follow, error) {
	db := fg.DB.WithContext(ctx)
	follow := models.Follow{}
	err := db.Table(followTable).Where("candidate_id=? and recruiter_id=?", candidateID, recruiterID).Take(&follow).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return nil, err
	}
	return &follow, nil
}
