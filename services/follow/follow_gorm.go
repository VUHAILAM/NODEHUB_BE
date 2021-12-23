package follow

import (
	"context"

	models2 "gitlab.com/hieuxeko19991/job4e_be/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	followTable    = "follow"
	candidateTable = "candidate"
	recruiterTable = "recruiter"
)

type IFollowDatabase interface {
	Create(ctx context.Context, follow *models2.Follow) error
	Delete(ctx context.Context, follow *models2.Follow) error
	GetFollow(ctx context.Context, candidateID int64, recruiterID int64) (*models2.Follow, error)
	CountFollowOfRecruiter(ctx context.Context, recruiterID int64) (int64, error)
	CountFollowOfCandidate(ctx context.Context, candidateID int64) (int64, error)
	GetFollowedRecruiter(ctx context.Context, recruiterID int64, offset, size int64) ([]*models2.Candidate, int64, error)
	GetFollowingRecruiter(ctx context.Context, candidateID int64, offset, size int64) ([]*models2.Recruiter, int64, error)
	GetListCandidateID(ctx context.Context, recruiterID int64) ([]*models2.Follow, error)
}

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

func (fg *FollowGorm) Create(ctx context.Context, follow *models2.Follow) error {
	db := fg.DB.WithContext(ctx)
	err := db.Table(followTable).Create(follow).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (fg *FollowGorm) Delete(ctx context.Context, follow *models2.Follow) error {
	db := fg.DB.WithContext(ctx)
	err := db.Table(followTable).Where("candidate_id=? and recruiter_id=?", follow.CandidateID, follow.RecruiterID).Delete(follow).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (fg *FollowGorm) GetFollow(ctx context.Context, candidateID int64, recruiterID int64) (*models2.Follow, error) {
	db := fg.DB.WithContext(ctx)
	follow := models2.Follow{}
	err := db.Table(followTable).Where("candidate_id=? and recruiter_id=?", candidateID, recruiterID).Take(&follow).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return nil, err
	}
	return &follow, nil
}

func (fg *FollowGorm) CountFollowOfRecruiter(ctx context.Context, recruiterID int64) (int64, error) {
	var count int64
	err := fg.DB.WithContext(ctx).Table(followTable).Where("recruiter_id=?", recruiterID).Count(&count).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}

func (fg *FollowGorm) CountFollowOfCandidate(ctx context.Context, candidateID int64) (int64, error) {
	var count int64
	err := fg.DB.WithContext(ctx).Table(followTable).Where("candidate_id=?", candidateID).Count(&count).Error
	if err != nil {
		fg.Logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}

func (fg *FollowGorm) GetFollowedRecruiter(ctx context.Context, recruiterID int64, offset, size int64) ([]*models2.Candidate, int64, error) {
	var candidates []*models2.Candidate
	db := fg.DB.WithContext(ctx).Table(candidateTable).Select("candidate.*").
		Joins("JOIN "+followTable+" ON candidate.candidate_id=follow.candidate_id").
		Where("follow.recruiter_id=?", recruiterID).Find(&candidates)
	total := db.RowsAffected
	candidates = make([]*models2.Candidate, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Order(followTable + ".updated_at desc").Find(&candidates).Error
	if err != nil {
		fg.Logger.Error("FollowGorm: GetFollowedRecruiter error", zap.Error(err))
		return nil, 0, err
	}

	return candidates, total, nil
}

func (fg *FollowGorm) GetFollowingRecruiter(ctx context.Context, candidateID int64, offset, size int64) ([]*models2.Recruiter, int64, error) {
	var recruiters []*models2.Recruiter
	db := fg.DB.WithContext(ctx).Table(recruiterTable).Select("recruiter.*").
		Joins("JOIN "+followTable+" ON recruiter.recruiter_id=follow.recruiter_id").
		Where("follow.candidate_id=?", candidateID).Find(&recruiters)
	total := db.RowsAffected
	recruiters = make([]*models2.Recruiter, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Order(followTable + ".updated_at desc").Find(&recruiters).Error
	if err != nil {
		fg.Logger.Error("FollowGorm: GetFollowingRecruiter error", zap.Error(err))
		return nil, 0, err
	}

	return recruiters, total, nil
}

func (fg *FollowGorm) GetListCandidateID(ctx context.Context, recruiterID int64) ([]*models2.Follow, error) {
	var candidateIDs []*models2.Follow
	db := fg.DB.WithContext(ctx).Table(followTable).Where("recruiter_id=?", recruiterID).Find(&candidateIDs)
	if db.Error != nil {
		fg.Logger.Error(db.Error.Error())
		return nil, db.Error
	}
	return candidateIDs, nil
}
