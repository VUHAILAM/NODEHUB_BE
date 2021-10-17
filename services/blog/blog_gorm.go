package blog

import (
	"context"

	"go.uber.org/zap"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"gorm.io/gorm"
)

const tableAccount = "account"

type AccountGorm struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAccountGorm(db *gorm.DB, logger *zap.Logger) *AccountGorm {
	return &AccountGorm{
		db:     db,
		logger: logger,
	}
}

func (g *AccountGorm) Create(ctx context.Context, account *models.Account) error {
	db := g.db.WithContext(ctx)
	err := db.Table(tableAccount).Create(account).Error
	if err != nil {
		g.logger.Error("AccountGorm: Create account error", zap.Error(err))
		return err
	}
	return nil
}

func (g *AccountGorm) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	db := g.db.WithContext(ctx)
	acc := models.Account{}
	err := db.Table(tableAccount).Where("email=?", email).First(&acc).Error
	if err != nil {
		g.logger.Error("AccountGorm: Get account error", zap.Error(err))
		return nil, err
	}

	return &acc, nil
}

func (g *AccountGorm) GetAccountByID(ctx context.Context, id string) (*models.Account, error) {
	db := g.db.WithContext(ctx)
	acc := models.Account{}
	err := db.Table(tableAccount).Where("id=?", id).First(&acc).Error
	if err != nil {
		g.logger.Error("AccountGorm: Get account error", zap.Error(err))
		return nil, err
	}
	return &acc, nil
}
