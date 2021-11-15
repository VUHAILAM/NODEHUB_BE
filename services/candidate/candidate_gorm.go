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

type ICandidateDatabase interface {
	Create(ctx context.Context, createData *models.Candidate) (int64, error)
	Update(ctx context.Context, candidateID int64, updateData *models.Candidate) error
	GetByCandidateID(ctx context.Context, candidateID int64) (*models.Candidate, error)
	GetAllCandidateForAdmin(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListCandidateAdmin, error)
	UpdateReviewCandidateByAdmin(ctx context.Context, candidate_id int64, data map[string]interface{}) error
	UpdateStatusCandidate(ctx context.Context, candidate *models.RequestUpdateStatusCandidate, candidate_id int64) error
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
	err := db.Table(candidateTable).Where("candidate_id=?", candidateID).Take(&candidate).Error
	if err != nil {
		g.Logger.Error(err.Error())
		return nil, err
	}
	return &candidate, nil
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
	c.nodehub_review, c.cv_manage, c.experience_manage, c.social_manage, c.project_manage, 
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
