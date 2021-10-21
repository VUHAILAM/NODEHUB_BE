package blog

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IBlogService interface {
	GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error)
	CreateBlog(ctx context.Context, blog *models.RequestCreateBlog) error
	UpdateBlog(ctx context.Context, blog *models.RequestCreateBlog, Blog_id int64) error
}

type IBlogDatabase interface {
	GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error)
	Create(ctx context.Context, blog *models.Blog) error
	Update(ctx context.Context, blog *models.RequestUpdateBlog, Blog_id int64) error
}

type Blog struct {
	BlogGorm  *BlogGorm
	SecretKey string
	Logger    *zap.Logger
}

func NewBlog(blogGorm *BlogGorm, secretKey string, logger *zap.Logger) *Blog {
	return &Blog{
		BlogGorm:  blogGorm,
		SecretKey: secretKey,
		Logger:    logger,
	}
}

func (b *Blog) GetListBlog(ctx context.Context, title string, page int64, size int64) (*models.ResponsetListBlog, error) {
	acc, err := b.BlogGorm.GetListBlog(ctx, title, page, size)
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
