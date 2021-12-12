package recruiter

import (
	"context"
	"fmt"
	"math"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
)

const (
	tableRecruiter = "recruiter"
)
const (
	tableAccount = "account"
)

const (
	tableRecruiterSkill = "recruiter_skill"
	tableSkill          = "skill"
)

type IRecruiterDatabase interface {
	Create(ctx context.Context, recruiter *models.Recruiter) (int64, error)
	GetAllRecruiterName(ctx context.Context) ([]string, error)
	AddRecruiterSkill(ctx context.Context, recruiterSkill *models.RecruiterSkill) error
	UpdateProfile(ctx context.Context, recruiter *models.RequestUpdateRecruiter, recruiter_id int64) error
	GetRecruiterSkill(ctx context.Context, recruiter_id int64) ([]models.ResponseRecruiterSkill, error)
	GetProfile(ctx context.Context, id int64) (*models.Recruiter, error)
	GetAllRecruiterForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListRecruiter, error)
	UpdateRecruiterByAdmin(ctx context.Context, recruiter_id int64, data map[string]interface{}) error
	UpdateStatusRecruiter(ctx context.Context, updateRequest *models.RequestUpdateStatusRecruiter, recruiter_id int64) error
	GetAllRecruiterForCandidate(ctx context.Context, recruiterName string, skillName string, address string, page int64, size int64) (*models.ResponsetListRecruiterForCandidate, error)
	DeleteRecruiterSkill(ctx context.Context, recruiter_skill_id int64) error
	SearchRecruiter(ctx context.Context, text string, offset, size int64) ([]*models.Recruiter, int64, error)
	GetAllRecruiter(ctx context.Context, offset, size int64) ([]*models.Recruiter, int64, error)
	GetAllSkillByRecruiterID(ctx context.Context, recruiterID int64) ([]*models.Skill, error)
	Count(ctx context.Context) (int64, error)
	GetPremiumField(ctx context.Context, recruiterID int64) (bool, error)
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
func (r *RecruiterGorm) GetAllRecruiterName(ctx context.Context) ([]string, error) {
	db := r.db.WithContext(ctx)
	var names []struct {
		Name string
	}
	err := db.Table(tableRecruiter).Select("name").Find(&names).Error
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	res := make([]string, 0)
	for _, n := range names {
		res = append(res, n.Name)
	}
	return res, nil
}

func (r *RecruiterGorm) GetProfile(ctx context.Context, id int64) (*models.Recruiter, error) {
	db := r.db.WithContext(ctx)
	rec := models.Recruiter{}
	err := db.Table(tableRecruiter).Where("recruiter_id=?", id).First(&rec).Error
	if err != nil {
		r.logger.Error("RecruiterGorm: Get recruiter error", zap.Error(err))
		return nil, err
	}
	acc := models.Account{}
	err = db.Table(tableAccount).Where("id=?", id).Take(&acc).Error
	if err != nil {
		r.logger.Error(err.Error())
	}
	rec.Email = acc.Email
	return &rec, nil
}

func (r *RecruiterGorm) GetAllRecruiter(ctx context.Context, offset, size int64) ([]*models.Recruiter, int64, error) {
	var recruiters []*models.Recruiter
	db := r.db.WithContext(ctx).Table(tableRecruiter).Select("recruiter.*").Find(&recruiters)
	total := db.RowsAffected
	recruiters = make([]*models.Recruiter, 0)
	err := db.Offset(int(offset)).Limit(int(size)).Order("updated_at desc").Find(&recruiters).Error
	if err != nil {
		r.logger.Error(err.Error())
		return nil, 0, err
	}

	return recruiters, total, nil
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

func (r *RecruiterGorm) DeleteRecruiterSkill(ctx context.Context, recruiter_skill_id int64) error {
	db := r.db.WithContext(ctx)
	recruiter_skill := models.RecruiterSkill{}
	err := db.Table(tableRecruiterSkill).Delete(&recruiter_skill, recruiter_skill_id).Error
	if err != nil {
		r.logger.Error("RecruiterGorm: Delete skill error", zap.Error(err), zap.Int64("recruiter_skill_id", recruiter_skill_id))
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

/*Get list recruiter for admin*/
func (r *RecruiterGorm) GetAllRecruiterForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListRecruiter, error) {
	db := r.db.WithContext(ctx)
	arr := []models.RecruiterForAdmin{}
	resutl := models.ResponsetListRecruiter{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`select r.recruiter_id, r.name, r.address, r.avartar, r.banner, 
	r.phone, r.website, r.description, r.employee_quantity, r.contacter_name, r.contacter_phone, 
	r.media, r.premium, r.nodehub_review, a.status, r.created_at, r.updated_at
	FROM nodehub.recruiter r
	left join nodehub.account a on r.recruiter_id = a.id
	where r.name like ? ORDER BY r.created_at desc LIMIT ?, ?`, "%"+name+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.recruiter where name like ?`, "%"+name+"%").Scan(&total)
	if err != nil {
		r.logger.Error("RecruiterGorm: Get List Recruiter error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.Total = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}

func (r *RecruiterGorm) UpdateRecruiterByAdmin(ctx context.Context, recruiter_id int64, data map[string]interface{}) error {
	db := r.db.WithContext(ctx)
	err := db.Table(tableRecruiter).Where("recruiter_id=?", recruiter_id).Updates(data).Error
	if err != nil {
		r.logger.Error("RecruiterGorm: Update recruiter error", zap.Error(err), zap.Int64("recruiter_id", recruiter_id))
		return err
	}
	return nil
}

func (r *RecruiterGorm) UpdateStatusRecruiter(ctx context.Context, recruiter *models.RequestUpdateStatusRecruiter, recruiter_id int64) error {
	db := r.db.WithContext(ctx)
	err := db.Table(tableAccount).Where("id = ?", recruiter_id).Updates(map[string]interface{}{
		"status": recruiter.Status}).Error
	if err != nil {
		r.logger.Error("RecruiterGorm: Update status recruiter error", zap.Error(err))
		return err
	}
	return nil
}

/*Get list recruiter for candidate*/
func (r *RecruiterGorm) GetAllRecruiterForCandidate(ctx context.Context, recruiterName string, skillName string, address string, page int64, size int64) (*models.ResponsetListRecruiterForCandidate, error) {
	db := r.db.WithContext(ctx)
	arr := []models.RecruiterForCandidateCheck{}
	resutl := models.ResponsetListRecruiterForCandidate{}
	offset := (page - 1) * size
	limit := size

	query := `select distinct r.recruiter_id, r.name, r.address, r.avartar, r.banner, 
	r.description, r.employee_quantity, s.name as skill_name, s.icon as skill_icon
	FROM nodehub.recruiter r
	left JOIN nodehub.recruiter_skill rs on rs.recruiter_id = r.recruiter_id 
	left join nodehub.skill s on rs.skill_id  = s.skill_id
	where r.name like @RName`

	querysum := `select count(*) FROM nodehub.recruiter r
	left JOIN nodehub.recruiter_skill rs on rs.recruiter_id = r.recruiter_id 
	left join nodehub.skill s on rs.skill_id  = s.skill_id
	where r.name like @RName`

	if skillName != "" {
		query += `and s.name like @SName`
		querysum += `and s.name like @SName`
	}

	if address != "" {
		query += `and r.address like @ADD`
		querysum += `and r.address like @ADD`
	}

	type NamedArgument struct {
		RName string
		SName string
		ADD   string
		OFF   int64
		LI    int64
	}

	var total int64
	//search query
	data, err := db.Raw(query+` ORDER BY r.name LIMIT @OFF, @LI`, NamedArgument{RName: "%" + recruiterName + "%", SName: "%" + skillName + "%", ADD: "%" + address + "%", OFF: offset, LI: limit}).Rows()
	// count query
	db.Raw(querysum, NamedArgument{RName: "%" + recruiterName + "%", SName: "%" + skillName + "%", ADD: "%" + address + "%", OFF: offset, LI: limit}).Scan(&total)
	if err != nil {
		r.logger.Error("RecruiterGorm: Get List Recruiter error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	fmt.Println("checkk data: ", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println("checkk data 1: ", arr[i].Skill_name)
		//TODO
	}

	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.Total = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}

func (r *RecruiterGorm) SearchRecruiter(ctx context.Context, text string, offset, size int64) ([]*models.Recruiter, int64, error) {
	var recruiters []*models.Recruiter
	db := r.db.WithContext(ctx).Table(tableRecruiter).Joins("Join "+tableRecruiterSkill+" on recruiter_skill.recruiter_id=recruiter.recruiter_id").
		Joins("Join "+tableSkill+" on recruiter_skill.skill_id=skill.skill_id").
		Where("MATCH(recruiter.name) AGAINST(?) OR MATCH(skill.name) AGAINST(?)", text, text).
		Group("recruiter.recruiter_id").Find(&recruiters)
	total := db.RowsAffected
	recruiters = make([]*models.Recruiter, 0)
	res := db.Offset(int(offset)).Limit(int(size)).Find(&recruiters)
	err := res.Error
	if err != nil {
		r.logger.Error(err.Error())
		return nil, 0, err
	}
	return recruiters, total, nil
}

func (r *RecruiterGorm) GetAllSkillByRecruiterID(ctx context.Context, recruiterID int64) ([]*models.Skill, error) {
	var skills []*models.Skill
	db := r.db.WithContext(ctx).Table(tableSkill).Joins("Join "+tableRecruiterSkill+" on recruiter_skill.skill_id=skill.skill_id").
		Where("recruiter_skill.recruiter_id=?", recruiterID).Find(&skills)
	if db.Error != nil {
		r.logger.Error(db.Error.Error())
		return nil, db.Error
	}
	return skills, nil
}

func (r *RecruiterGorm) Count(ctx context.Context) (int64, error) {
	var count int64
	db := r.db.WithContext(ctx)
	err := db.Table(tableRecruiter).Count(&count).Error
	if err != nil {
		r.logger.Error(err.Error())
		return 0, err
	}

	return count, nil
}

func (r *RecruiterGorm) GetPremiumField(ctx context.Context, recruiterID int64) (bool, error) {
	var premium bool
	db := r.db.WithContext(ctx)
	err := db.Table(tableRecruiter).Select("recruiter.premium").Where("recruiter_id=?", recruiterID).Take(&premium).Error
	if err != nil {
		r.logger.Error(err.Error(), zap.Error(err))
		return false, err
	}
	return premium, nil
}
