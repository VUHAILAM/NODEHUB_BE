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

const (
	tableRecruiterSkill = "recruiter_skill"
)

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

func (r *RecruiterGorm) GetProfile(ctx context.Context, id int64) (*models.Recruiter, error) {
	db := r.db.WithContext(ctx)
	acc := models.Recruiter{}
	err := db.Table(tableRecruiter).Where("recruiter_id=?", id).First(&acc).Error
	if err != nil {
		r.logger.Error("RecruiterGorm: Get recruiter error", zap.Error(err))
		return nil, err
	}
	return &acc, nil
}

func (r *RecruiterGorm) UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error {
	db := r.db.WithContext(ctx)
	err := db.Table(tableRecruiter).Where("recruiter_id = ?", recruiter_id).Updates(map[string]interface{}{
		"name":              recruiter.Name,
		"address":           recruiter.Address,
		"avartar":           recruiter.Avartar,
		"banner":            recruiter.Banner,
		"phone":             recruiter.Phone,
		"website":           recruiter.Website,
		"description":       recruiter.Description,
		"employee_quantity": recruiter.EmployeeQuantity,
		"contacter_name":    recruiter.ContacterName,
		"contacter_phone":   recruiter.ContacterPhone,
		"media":             recruiter.Media}).Error
	if err != nil {
		r.logger.Error("MediaGorm: Update media error", zap.Error(err))
		return err
	}
	return nil
}

//recruiter skill
func (r *RecruiterGorm) AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error {
	db := r.db.WithContext(ctx)
	err := db.Table(tableRecruiterSkill).Create(recruiterSkill).Error
	if err != nil {
		r.logger.Error("AccountGorm: Create recruiter error", zap.Error(err))
		return err
	}
	return nil
}

func (r *RecruiterGorm) GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error) {
	db := r.db.WithContext(ctx)
	arr := []models.ResponseRecruiterSkill{}
	data, err := db.Raw(`SELECT rs.id , rs.recruiter_id , rs.skill_id , s.name , s.description , s.questions , s.icon , s.status , rs.created_at ,rs.updated_at 
	FROM nodehub.recruiter_skill rs
	LEFT JOIN nodehub.skill s
	ON rs.skill_id = s.skill_id
	where s.status = 1 and rs.recruiter_id = ?`, recruiter_id).Rows()
	if err != nil {
		r.logger.Error("MediaGorm: Get slide error", zap.Error(err))
		return nil, err
	}
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	return arr, nil
}
