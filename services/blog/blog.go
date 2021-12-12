package blog

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IBlogService interface {
	GetDetailBlog(ctx context.Context, blogID int64) (*models.Blog, error)
	GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error)
	GetListBlogUser(ctx context.Context, title string, category_id int64, page int64, size int64) (*models.ResponsetListBlog, error)
	CreateBlog(ctx context.Context, blog *models.RequestCreateBlog) error
	UpdateBlog(ctx context.Context, blog *models.RequestCreateBlog, Blog_id int64) error
	GetListBlogByCategory(ctx context.Context, category_id int64, page int64, size int64) (*models.ResponsetListBlog, error)
}

type Blog struct {
	BlogGorm *BlogGorm
	Logger   *zap.Logger
}

func NewBlog(blogGorm *BlogGorm, logger *zap.Logger) *Blog {
	return &Blog{
		BlogGorm: blogGorm,
		Logger:   logger,
	}
}

func (b *Blog) GetDetailBlog(ctx context.Context, blogID int64) (*models.Blog, error) {
	return b.BlogGorm.GetDetailBlog(ctx, blogID)
}

func (b *Blog) GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error) {
	acc, err := b.BlogGorm.GetListBlog(ctx, title, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (b *Blog) GetListBlogUser(ctx context.Context, title string, category_id int64, page int64, size int64) (*models.ResponsetListBlog, error) {
	acc, err := b.BlogGorm.GetListBlogUser(ctx, title, category_id, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
func (b *Blog) GetListBlogByCategory(ctx context.Context, category_id int64, page int64, size int64) (*models.ResponsetListBlog, error) {
	acc, err := b.BlogGorm.GetListBlogByCategory(ctx, category_id, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (b *Blog) CreateBlog(ctx context.Context, blog *models.RequestCreateBlog) error {
	blogModels := &models.Blog{
		Title:       blog.Title,
		Category_id: blog.Category_id,
		Icon:        blog.Icon,
		Description: blog.Description,
		Excerpts:    blog.Excerpts,
		Status:      blog.Status,
	}
	err := b.BlogGorm.Create(ctx, blogModels)
	if err != nil {
		return err
	}
	return nil
}
func (b *Blog) UpdateBlog(ctx context.Context, blog *models.RequestCreateBlog, Blog_id int64) error {
	blogModels := &models.RequestUpdateBlog{
		Title:       blog.Title,
		Category_id: blog.Category_id,
		Icon:        blog.Icon,
		Description: blog.Description,
		Excerpts:    blog.Excerpts,
		Status:      blog.Status,
	}
	err := b.BlogGorm.Update(ctx, blogModels, Blog_id)
	if err != nil {
		return err
	}
	return nil
}
