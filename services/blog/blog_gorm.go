package blog

import (
	"context"
	"math"

	"go.uber.org/zap"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gorm.io/gorm"
)

const tableAccount = "blog"

type BlogGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewBlogGorm(db *gorm.DB, logger *zap.Logger) *BlogGorm {
	return &BlogGorm{
		db:     db,
		logger: logger,
	}
}

func (g *BlogGorm) GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error) {
	db := g.db.WithContext(ctx)
	arr := []models.ResponseBlog{}
	resutl := models.ResponsetListBlog{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`SELECT b.blog_id, s.setting_id as 'category_id', s.name as 'category_name', b.title, b.icon, b.excerpts, b.description, b.status, b.created_at, b.updated_at 
	FROM nodehub.blog b INNER join setting s on b.category_id = s.setting_id  where  b.title like ? ORDER BY b.blog_id desc LIMIT ?, ?`, "%"+title+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.blog where  title like ?`, "%"+title+"%").Scan(&total)
	if err != nil {
		g.logger.Error("BlogGorm: Get blog error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.TotalBlog = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}

func (g *BlogGorm) Create(ctx context.Context, blog *models.Blog) error {
	db := g.db.WithContext(ctx)
	err := db.Table(tableAccount).Create(blog).Error
	if err != nil {
		g.logger.Error("BlogGorm: Create blog error", zap.Error(err))
		return err
	}
	return nil
}
func (g *BlogGorm) Update(ctx context.Context, blog *models.RequestUpdateBlog, Blog_id int64) error {
	db := g.db.WithContext(ctx)
	err := db.Table(tableAccount).Where("blog_id = ?", Blog_id).Updates(map[string]interface{}{
		"category_id": blog.Category_id,
		"title":       blog.Title,
		"icon":        blog.Icon,
		"excerpts":    blog.Excerpts,
		"description": blog.Description,
		"status":      blog.Status}).Error
	if err != nil {
		g.logger.Error("BlogGorm: Update blog error", zap.Error(err))
		return err
	}
	return nil
}

func (g *BlogGorm) GetListBlogUser(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error) {
	db := g.db.WithContext(ctx)
	arr := []models.ResponseBlog{}
	resutl := models.ResponsetListBlog{}
	offset := (page - 1) * size
	limit := size
	var total int64
	//search query
	data, err := db.Raw(`SELECT b.blog_id, s.name as 'category_name', b.title, b.icon, b.excerpts, b.description, b.status, b.created_at, b.updated_at FROM nodehub.blog b INNER join setting s on b.category_id = s.setting_id  where  b.title like ? and b.status = 1 ORDER BY b.blog_id desc LIMIT ?, ?`, "%"+title+"%", offset, limit).Rows()
	// count query
	db.Raw(`SELECT count(*) FROM nodehub.blog where  title like ? and status = 1`, "%"+title+"%").Scan(&total)
	if err != nil {
		g.logger.Error("BlogGorm: Get blog error", zap.Error(err))
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		// ScanRows scan a row into user
		db.ScanRows(data, &arr)
	}
	var temp float64 = math.Ceil(float64(total) / float64(size))
	resutl.TotalBlog = total
	resutl.TotalPage = temp
	resutl.CurrentPage = page
	resutl.Data = arr

	return &resutl, nil
}
