package skill

import (
	"context"
	"math"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableSkill = "skill"

type SkillGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewSkillGorm(db *gorm.DB, logger *zap.Logger) *SkillGorm {
	return &SkillGorm{
		db:     db,
		logger: logger,
	}
}

/*Create Skill*/

func (s *SkillGorm) Create(ctx context.Context, skill *models.Skill) error {
	db := s.db.WithContext(ctx)
	err := db.Table(tableSkill).Create(skill).Error
	if err != nil {
		s.logger.Error("SkillGorm: Create skill error", zap.Error(err))
		return err
	}
	return nil
}

/*Update Skill*/
func (s *SkillGorm) Update(ctx context.Context, skill *models.RequestUpdateSkill, skill_id int64) error {
	db := s.db.WithContext(ctx)
	err := db.Table(tableSkill).Where("skill_id = ?", skill_id).Updates(map[string]interface{}{
		"name":        skill.Name,
		"description": skill.Description,
		"questions":   skill.Questions,
		"icon":        skill.Icon,
		"status":      skill.Status}).Error
	if err != nil {
		s.logger.Error("SkillGorm: Update skill error", zap.Error(err))
		return err
	}
	return nil
}

func (s *SkillGorm) Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSkill, error) {
	db := s.db.WithContext(ctx)
	arr := []models.Skill{}
	resutl := models.ResponsetListSkill{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`select skill_id, name, description, questions, icon, status, created_at, updated_at FROM nodehub.skill where name like ? ORDER BY skill_id desc LIMIT ?, ?`, "%"+name+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.skill where name like ?`, "%"+name+"%").Scan(&total)
	if err != nil {
		s.logger.Error("BlogGorm: Get blog error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.TotalSkill = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}

func (s *SkillGorm) GetAll(ctx context.Context, name string) ([]models.Skill, error) {
	db := s.db.WithContext(ctx)
	arr := []models.Skill{}
	data, err := db.Raw(`select skill_id, name, description, questions, icon, status, created_at, updated_at FROM nodehub.skill where name like ? and status = 1`, "%"+name+"%").Rows()
	if err != nil {
		s.logger.Error("SkillGorm: Get Skill error", zap.Error(err))
		return nil, err
	}
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	return arr, nil
}

func (s *SkillGorm) GetSkillByIDs(ctx context.Context, skillIDs []int64) ([]models.Skill, error) {
	db := s.db.WithContext(ctx)
	var res []models.Skill
	err := db.Table(tableSkill).Where(skillIDs).Find(&res).Error
	if err != nil {
		s.logger.Error("SkillGorm: Get Skill error", zap.Error(err))
		return nil, err
	}
	return res, nil
}
