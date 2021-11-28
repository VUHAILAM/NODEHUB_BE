package account

import (
	"context"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	tableSetting = "setting"
	tableAccount = "account"
)

type IAccountDatabase interface {
	Create(ctx context.Context, account *models.Account) (int64, error)
	GetAccountByEmail(ctx context.Context, email string) (*models.Account, error)
	GetAccountByID(ctx context.Context, id int64) (*models.Account, error)
	UpdatePassword(ctx context.Context, email, password, tokenHash string) error
	UpdateVerifyEmail(ctx context.Context, email string) error
}

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

func (g *AccountGorm) Create(ctx context.Context, account *models.Account) (int64, error) {
	db := g.db.WithContext(ctx)
	err := db.Table(tableAccount).Create(account).Error
	if err != nil {
		g.logger.Error("AccountGorm: Create account error", zap.Error(err))
		return 0, err
	}
	return account.Id, nil
}

func (g *AccountGorm) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	db := g.db.WithContext(ctx)
	acc := models.Account{}
	err := db.Table(tableAccount).Where("email=?", email).First(&acc).Error
	if err != nil {
		g.logger.Error("AccountGorm: Get account error", zap.Error(err))
		return nil, err
	}
	setting := models.Setting{}
	err = db.Table(tableSetting).Where("setting_id=?", acc.Type).First(&setting).Error
	if err != nil {
		g.logger.Error("AccountGorm: Get setting error", zap.Error(err))
		return nil, err
	}
	acc.RoleName = setting.Name
	return &acc, nil
}

func (g *AccountGorm) GetAccountByID(ctx context.Context, id int64) (*models.Account, error) {
	db := g.db.WithContext(ctx)
	acc := models.Account{}
	err := db.Table(tableAccount).Where("id=?", id).First(&acc).Error
	if err != nil {
		g.logger.Error("AccountGorm: Get account error", zap.Error(err))
		return nil, err
	}
	setting := models.Setting{}
	err = db.Table(tableSetting).Where("setting_id=?", acc.Type).First(&setting).Error
	if err != nil {
		g.logger.Error("AccountGorm: Get setting error", zap.Error(err))
		return nil, err
	}
	acc.RoleName = setting.Name
	return &acc, nil
}

func (g *AccountGorm) UpdatePassword(ctx context.Context, email, password, tokenHash string) error {
	db := g.db.WithContext(ctx)
	err := db.Table(tableAccount).Where("email=?", email).Updates(map[string]interface{}{"password": password, "token_hash": tokenHash}).Error
	if err != nil {
		g.logger.Error("AccountGorm: Update account password error", zap.Error(err))
		return err
	}
	return nil
}

func (g *AccountGorm) UpdateVerifyEmail(ctx context.Context, email string) error {
	db := g.db.WithContext(ctx)
	err := db.Table(tableAccount).Where("email=?", email).Update("status", true).Error
	if err != nil {
		g.logger.Error("AccountGorm: Update verify email error", zap.Error(err))
		return err
	}
	return nil
}
