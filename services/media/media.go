package media

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
)

type IMediaService interface {
	CreateMedia(ctx context.Context, media *models.RequestCreateMedia) error
	UpdateMedia(ctx context.Context, media *models.RequestUpdateMedia, media_id int64) error
	GetListMediaPaging(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListMedia, error)
	GetSlide(ctx context.Context) ([]models.Media, error)
}

type IMediaDatabase interface {
	Create(ctx context.Context, media *models.Media) error
	Update(ctx context.Context, media *models.RequestUpdateMedia, media_id int64) error
	Get(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListMedia, error)
	GetSlide(ctx context.Context) ([]models.Media, error)
}

type Media struct {
	MediaGorm *MediaGorm
	Logger    *zap.Logger
}

func NewMediaCategory(mediaGorm *MediaGorm, logger *zap.Logger) *Media {
	return &Media{
		MediaGorm: mediaGorm,
		Logger:    logger,
	}
}

/*Create Media*/
func (m *Media) CreateMedia(ctx context.Context, media *models.RequestCreateMedia) error {
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

/*Update Media*/
func (m *Media) UpdateMedia(ctx context.Context, media *models.RequestUpdateMedia, media_id int64) error {
	mediaModels := &models.RequestUpdateMedia{
		Name:   media.Name,
		Type:   media.Type,
		Status: media.Status}
	err := m.MediaGorm.Update(ctx, mediaModels, media_id)
	if err != nil {
		return err
	}
	return nil
}

func (m *Media) GetListMediaPaging(ctx context.Context, name string, page int64, size int64) (*models.ResponsetListMedia, error) {
	acc, err := m.MediaGorm.Get(ctx, name, page, size)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (m *Media) GetSlide(ctx context.Context) ([]models.Media, error) {
	acc, err := m.MediaGorm.GetSlide(ctx)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
