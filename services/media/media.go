package media

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IMediaService interface {
	CreateMedia(ctx context.Context, category *models.RequestCreateMedia) error
	// UpdateCategory(ctx context.Context, category *models.RequestCreateSetting, setting_id int64) error
	// GetListCategoryPaging(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error)
	// GetAllCategory(ctx context.Context) ([]models.Setting, error)
}

type IMediaDatabase interface {
	Create(ctx context.Context, category *models.Media) error
	// Update(ctx context.Context, category *models.RequestUpdateSetting, setting_id int64) error
	// Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error)
	// GetAll(ctx context.Context) ([]models.Setting, error)
}

type Media struct {
	MediaGorm *MediaGorm
	SecretKey string
	Logger    *zap.Logger
}

func NewMediaCategory(mediaGorm *MediaGorm, secretKey string, logger *zap.Logger) *Media {
	return &Media{
		MediaGorm: mediaGorm,
		SecretKey: secretKey,
		Logger:    logger,
	}
}

/*Create Media*/
func (m *Media) CreateMedia(ctx context.Context, media *models.Media) error {
	MediaModels := &models.Media{
		Name:   media.Name,
		Type:   media.Type,
		Status: media.Status}
	err := m.MediaGorm.Create(ctx, MediaModels)
	if err != nil {
		return err
	}
	return nil
}

// /*Update Category*/
// func (c *Category) UpdateCategory(ctx context.Context, category *models.RequestCreateSetting, setting_id int64) error {
// 	categoryModels := &models.RequestUpdateSetting{
// 		Name: category.Name,
// 		Type: category.Type}
// 	err := c.CategoryGorm.Update(ctx, categoryModels, setting_id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *Category) GetListCategoryPaging(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListSetting, error) {
// 	acc, err := c.CategoryGorm.Get(ctx, name, page, size)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }

// func (c *Category) GetAllCategory(ctx context.Context) ([]models.Setting, error) {
// 	acc, err := c.CategoryGorm.GetAll(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }
