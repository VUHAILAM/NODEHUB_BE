package category

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"go.uber.org/zap"
)

type ICategoryService interface {
	CreateCategory(ctx context.Context, category *models.RequestCreateSetting) error
	UpdateCategory(ctx context.Context, category *models.RequestCreateSetting, setting_id int64) error
	GetListCategoryPaging(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error)
	GetAllCategory(ctx context.Context) ([]models.Setting, error)
}

type ICategoryDatabase interface {
	Create(ctx context.Context, category *models.Setting) error
	Update(ctx context.Context, category *models.RequestUpdateSetting, setting_id int64) error
	Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error)
	GetAll(ctx context.Context) ([]models.Setting, error)
}

type Category struct {
	CategoryGorm *CategoryGorm
	Logger       *zap.Logger
}

func NewCategory(categoryGorm *CategoryGorm, logger *zap.Logger) *Category {
	return &Category{
		CategoryGorm: categoryGorm,
		Logger:       logger,
	}
}

/*Create Category*/
func (c *Category) CreateCategory(ctx context.Context, category *models.RequestCreateSetting) error {
	CategoryModels := &models.Setting{
		Name: category.Name,
		Type: category.Type}
	err := c.CategoryGorm.Create(ctx, CategoryModels)
	if err != nil {
		return err
	}
	return nil
}

/*Update Category*/
func (c *Category) UpdateCategory(ctx context.Context, category *models.RequestCreateSetting, setting_id int64) error {
	categoryModels := &models.RequestUpdateSetting{
		Name: category.Name,
		Type: category.Type}
	err := c.CategoryGorm.Update(ctx, categoryModels, setting_id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) GetListCategoryPaging(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error) {
	acc, err := c.CategoryGorm.Get(ctx, name, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (c *Category) GetAllCategory(ctx context.Context) ([]models.Setting, error) {
	acc, err := c.CategoryGorm.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
