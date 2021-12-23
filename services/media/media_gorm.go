package media

import (
	"context"
	"math"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableAccount = "media"

type IMediaDatabase interface {
	Create(ctx context.Context, media *models.Media) error
	Update(ctx context.Context, media *models.RequestUpdateMedia, media_id int64) error
	Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListMedia, error)
	GetSlide(ctx context.Context) ([]models.Media, error)
}

type MediaGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewMediaGorm(db *gorm.DB, logger *zap.Logger) *MediaGorm {
	return &MediaGorm{
		db:     db,
		logger: logger,
	}
}

/*Create Media*/

func (m *MediaGorm) Create(ctx context.Context, media *models.Media) error {
	db := m.db.WithContext(ctx)
	err := db.Table(tableAccount).Create(media).Error
	if err != nil {
		m.logger.Error("MediaGorm: Create media error", zap.Error(err))
		return err
	}
	return nil
}

// /*Update Category*/
func (m *MediaGorm) Update(ctx context.Context, media *models.RequestUpdateMedia, media_id int64) error {
	db := m.db.WithContext(ctx)
	err := db.Table(tableAccount).Where("media_id = ?", media_id).Updates(map[string]interface{}{
		"name":   media.Name,
		"type":   media.Type,
		"status": media.Status}).Error
	if err != nil {
		m.logger.Error("MediaGorm: Update media error", zap.Error(err))
		return err
	}
	return nil
}

/*Get Media*/
func (m *MediaGorm) Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListMedia, error) {
	db := m.db.WithContext(ctx)
	arr := []models.Media{}
	resutl := models.ResponsetListMedia{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`select media_id, type, name, status, created_at, updated_at FROM nodehub.media where  name like ? ORDER BY media_id desc LIMIT ?, ?`, "%"+name+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.media where name like ?`, "%"+name+"%").Scan(&total)
	if err != nil {
		m.logger.Error("MediaGorm: Get Media error", zap.Error(err))
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

/*Get all slide media*/
func (m *MediaGorm) GetSlide(ctx context.Context) ([]models.Media, error) {
	db := m.db.WithContext(ctx)
	arr := []models.Media{}
	data, err := db.Raw(`select * FROM nodehub.media where type = "slide" and status = 1`).Rows()
	if err != nil {
		m.logger.Error("MediaGorm: Get slide error", zap.Error(err))
		return nil, err
	}
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	return arr, nil
}
