package candidate

import (
	"context"
	"math"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	candidateTable = "candidate"
)

const (
	tableAccount = "account"
)

const (
	tableCandidateSkill = "candidate_skill"
	tableSkill          = "skill"
)

type ICandidateDatabase interface {
	Create(ctx context.Context, createData *models.Candidate) (int64, error)
	Update(ctx context.Context, candidateID int64, updateData *models.Candidate) error
	GetByCandidateID(ctx context.Context, candidateID int64) (*models.Candidate, error)
	GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListCandidateAdmin, error)
	UpdateReviewCandidateByAdmin(ctx context.Context, candidate_id int64, data map[string]interface{}) error
	UpdateStatusCandidate(ctx context.Context, candidate *models.RequestUpdateStatusCandidate, candidate_id int64) error
	AddCandidateSkill(ctx context.Context, candidateSkill *models.CandidateSkill) error
	DeleteCandidateSkill(ctx context.Context, candidate_skill_id int64) error
	UpdateCandidateSkill(ctx context.Context, candidate_skill_id int64, data map[string]interface{}) error
	GetCandidateSkill(ctx context.Context, candidate_id int64) ([]models.ResponseCandidateSkill, error)
	SearchCandidate(ctx context.Context, text string, offset, page int64) ([]*models.Candidate, int64, error)
	GetAllCandidate(ctx context.Context, offset, size int64) ([]*models.Candidate, int64, error)
	GetAllSkillByCandidateID(ctx context.Context, candidateID int64) ([]*models.Skill, error)
	Count(ctx context.Context) (int64, error)
}

type CandidateGorm struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewCandidateGorm(db *gorm.DB, logger *zap.Logger) *CandidateGorm {
	return &CandidateGorm{
		DB:     db,
		Logger: logger,
	}
}

func (g *CandidateGorm) Create(ctx context.Context, createData *models.Candidate) (int64, error) {
	db := g.DB.WithContext(ctx)
	err := db.Table(candidateTable).Create(createData).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return 0, err
	}
	return createData.CandidateID, nil
}

func (g *CandidateGorm) Update(ctx context.Context, candidateID int64, updateData *models.Candidate) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(candidateTable).Where("candidate_id=?", candidateID).Updates(updateData).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (g *CandidateGorm) GetByCandidateID(ctx context.Context, candidateID int64) (*models.Candidate, error) {
	db := g.DB.WithContext(ctx)
	candidate := models.Candidate{}
	err := db.Table(candidateTable).Select("candidate.*").Where("candidate_id=?", candidateID).Take(&candidate).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, err
	}
	acc := models.Account{}
	err = db.Table(tableAccount).Where("id=?", candidateID).Take(&acc).Error
	if err != nil {
		g.Logger.Error(err.Error())
	}
	candidate.Email = acc.Email
	return &candidate, nil
}

func (g *CandidateGorm) GetAllCandidate(ctx context.Context, offset, size int64) ([]*models.Candidate, int64, error) {
	var candidates []*models.Candidate
	db := g.DB.WithContext(ctx).Table(candidateTable).Select("candidate.*").Find(&candidates)
	total := db.RowsAffected
	candidates = make([]*models.Candidate, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Order("updated_at desc").Find(&candidates).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, 0, err
	}

	return candidates, total, nil
}

/*Get list candidate for admin*/
func (g *CandidateGorm) GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListCandidateAdmin, error) {
	db := g.DB.WithContext(ctx)
	arr := []models.CandidateAdmin{}
	arr2 := []models.CandidateRequestAdmin{}
	resutl := models.ResponsetListCandidateAdmin{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`select  c.candidate_id, a.email, CONCAT_WS(" ", c.last_name, c.first_name) AS 'fullname', c.first_name, c.last_name, c.birth_day, c.address, c.avatar, c.banner, c.phone, c.find_job, 
	c.nodehub_review, c.cv_manage,c.education_manage, c.experience_manage, c.social_manage, c.project_manage, 
	c.certificate_manage, c.prize_manage, a.status ,c.created_at, c.updated_at
		FROM nodehub.candidate c
		left join nodehub.account a on c.candidate_id = a.id
		where CONCAT_WS(" ", c.last_name, c.first_name) like ? ORDER BY c.created_at desc LIMIT ?, ?`, "%"+name+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.candidate where CONCAT_WS(" ", last_name, first_name) like ?`, "%"+name+"%").Scan(&total)
	if err != nil {
		g.Logger.Error("CandidateGorm: Get List Candidate error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// data = candi.ToCandidateRequest()
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}

	for i := 0; i < len(arr); i++ {
		can, err := arr[i].ToCandidateRequestAdmin()
		if err != nil {
			g.Logger.Error(err.Error())
			return nil, err
		}
		arr2 = append(arr2, can)
	}

	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.Total = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr2

	return &resutl, nil
}

func (g *CandidateGorm) UpdateReviewCandidateByAdmin(ctx context.Context, candidate_id int64, data map[string]interface{}) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(candidateTable).Where("candidate_id=?", candidate_id).Updates(data).Error
	if err != nil {
		g.Logger.Error("CandidateGorm: Update ReviewCandidateByAdmin error", zap.Error(err), zap.Int64("candidate_id", candidate_id))
		return err
	}
	return nil
}

func (g *CandidateGorm) UpdateStatusCandidate(ctx context.Context, candidate *models.RequestUpdateStatusCandidate, candidate_id int64) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(tableAccount).Where("id = ?", candidate_id).Updates(map[string]interface{}{
		"status": candidate.Status}).Error
	if err != nil {
		g.Logger.Error("RecruiterGorm: Update status recruiter error", zap.Error(err))
		return err
	}
	return nil
}

//candidate skill
func (g *CandidateGorm) AddCandidateSkill(ctx context.Context, candidateSkill *models.CandidateSkill) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(tableCandidateSkill).Create(candidateSkill).Error
	if err != nil {
		g.Logger.Error("CandidateGorm: Create recruiter skill error", zap.Error(err))
		return err
	}
	return nil
}

func (g *CandidateGorm) DeleteCandidateSkill(ctx context.Context, candidate_skill_id int64) error {
	db := g.DB.WithContext(ctx)
	CandidateSkill := models.CandidateSkill{}
	err := db.Table(tableCandidateSkill).Delete(&CandidateSkill, candidate_skill_id).Error
	if err != nil {
		g.Logger.Error("RecruiterGorm: Delete skill error", zap.Error(err), zap.Int64("candidate_skill_id", candidate_skill_id))
		return err
	}
	return nil
}

func (g *CandidateGorm) UpdateCandidateSkill(ctx context.Context, candidate_skill_id int64, data map[string]interface{}) error {
	db := g.DB.WithContext(ctx)
	err := db.Table(tableCandidateSkill).Where("id=?", candidate_skill_id).Updates(data).Error
	if err != nil {
		g.Logger.Error("RecruiterGorm: Update recruiter error", zap.Error(err), zap.Int64("candidate_skill_id", candidate_skill_id))
		return err
	}
	return nil
}

func (g *CandidateGorm) GetCandidateSkill(ctx context.Context, candidate_id int64) ([]models.ResponseCandidateSkill, error) {
	db := g.DB.WithContext(ctx)
	arr := []models.ResponseCandidateSkill{}
	data, err := db.Raw(`SELECT cs.id , cs.candidate_id , cs.skill_id , s.name , s.description , s.questions , s.icon , cs.media ,s.status , cs.created_at ,cs.updated_at
	FROM nodehub.candidate_skill cs 
	LEFT JOIN nodehub.skill s
	ON cs.skill_id = s.skill_id
	where s.status = 1 and cs.candidate_id = ?`, candidate_id).Rows()
	if err != nil {
		g.Logger.Error("MediaGorm: Get slide error", zap.Error(err))
		return nil, err
	}
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	return arr, nil
}

func (g *CandidateGorm) SearchCandidate(ctx context.Context, text string, offset, page int64) ([]*models.Candidate, int64, error) {
	var candidates []*models.Candidate
	db := g.DB.WithContext(ctx).Table(candidateTable).Select("candidate.*").Joins("Join "+tableCandidateSkill+" on candidate_skill.candidate_id=candidate.candidate_id").
		Joins("Join "+tableSkill+" on candidate_skill.skill_id=skill.skill_id").
		Where("MATCH(candidate.first_name, candidate.last_name) AGAINST(?) OR MATCH(skill.name) AGAINST(?)", text, text).
		Group("candidate.candidate_id").Find(&candidates)
	total := db.RowsAffected
	candidates = make([]*models.Candidate, 0)
	err := db.Offset(int(offset)).Limit(int(page)).Find(&candidates).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, 0, err
	}
	return candidates, total, nil
}

func (g *CandidateGorm) GetAllSkillByCandidateID(ctx context.Context, candidateID int64) ([]*models.Skill, error) {
	var skills []*models.Skill
	db := g.DB.WithContext(ctx).Table(tableSkill).Joins("Join "+tableCandidateSkill+" on candidate_skill.skill_id=skill.skill_id").
		Where("candidate_skill.candidate_id=?", candidateID).Find(&skills)
	if db.Error != nil {
		g.Logger.Error(db.Error.Error())
		return nil, db.Error
	}
	return skills, nil
}

func (g *CandidateGorm) Count(ctx context.Context) (int64, error) {
	var count int64
	db := g.DB.WithContext(ctx)
	err := db.Table(candidateTable).Count(&count).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return 0, err
	}

	return count, nil
}
