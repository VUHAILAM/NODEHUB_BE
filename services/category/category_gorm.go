package category

import (
	"context"
	"math"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"

	"gorm.io/gorm"
)

const tableAccount = "setting"

type CategoryGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewCategoryGorm(db *gorm.DB, logger *zap.Logger) *CategoryGorm {
	return &CategoryGorm{
		db:     db,
		logger: logger,
	}
}

/*Create Category*/

func (c *CategoryGorm) Create(ctx context.Context, category *models.Setting) error {
	db := c.db.WithContext(ctx)
	err := db.Table(tableAccount).Create(category).Error
	if err != nil {
		c.logger.Error("CategoryGorm: Create category error", zap.Error(err))
		return err
	}
	return nil
}

/*Update Category*/
func (c *CategoryGorm) Update(ctx context.Context, category *models.RequestUpdateSetting, setting_id int64) error {
	db := c.db.WithContext(ctx)
	err := db.Table(tableAccount).Where("setting_id = ?", setting_id).Updates(map[string]interface{}{
		"name": category.Name,
		"type": category.Type}).Error
	if err != nil {
		c.logger.Error("CategoryGorm: Update category error", zap.Error(err))
		return err
	}
	return nil
}

/*Get Category*/
func (c *CategoryGorm) Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error) {
	db := c.db.WithContext(ctx)
	arr := []models.Setting{}
	resutl := models.ResponsetListSetting{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`select * FROM nodehub.setting where  name like ? and type = "blog" ORDER BY setting_id desc LIMIT ?, ?`, "%"+name+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.setting where name like ? and type = "blog"`, "%"+name+"%").Scan(&total)
	if err != nil {
		c.logger.Error("CategoryGorm: Get Category error", zap.Error(err))
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

/*Get all category*/
func (c *CategoryGorm) GetAll(ctx context.Context) ([]models.Setting, error) {
	db := c.db.WithContext(ctx)
	arr := []models.Setting{}
	data, err := db.Raw(`select * FROM nodehub.setting where type = "blog"`).Rows()
	if err != nil {
		c.logger.Error("CategoryGorm: Get Category error", zap.Error(err))
		return nil, err
	}
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	return arr, nil
}
