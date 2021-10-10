package account

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	Login(ctx context.Context, email string, password string) (string, error)
	Logout(ctx context.Context, email string) error
	Register(ctx context.Context, account *models.RequestRegisterAccount) error
	GetUserFromCookie(ctx context.Context, cookie string) (*models.Account, error)
}

type IAccountDatabase interface {
	Create(ctx context.Context, account *models.Account) error
	GetAccountByEmail(ctx context.Context, email string) (*models.Account, error)
	GetAccountByID(ctx context.Context, id string) (*models.Account, error)
}

type Account struct {
	AccountGorm *AccountGorm
	SecretKey   string
	Logger      *zap.Logger
}

func NewAccount(accountGorm *AccountGorm, secretKey string, logger *zap.Logger) *Account {
	return &Account{
		AccountGorm: accountGorm,
		SecretKey:   secretKey,
		Logger:      logger,
	}
}

func (a *Account) Login(ctx context.Context, email string, password string) (string, error) {
	acc, err := a.AccountGorm.GetAccountByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if acc.Id == 0 {
		a.Logger.Error("Account by email not found")
		return "", errors.New("Account not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		a.Logger.Error("Incorrect password")
		return "", errors.Wrap(err, "Incorrect password")
	}
	jsonAccount, err := json.Marshal(acc)
	if err != nil {
		a.Logger.Error("Can not convert account to json", zap.Error(err))
		return "", err
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    string(jsonAccount),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(a.SecretKey))
	if err != nil {
		a.Logger.Error("Cannot get token", zap.Error(err))
		return "", err
	}

	return token, nil
}

func (a *Account) Logout(ctx context.Context, email string) error {
	return nil
}

func (a *Account) Register(ctx context.Context, account *models.RequestRegisterAccount) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), 8)
	accountModels := &models.Account{
		Email:    account.Email,
		Phone:    account.Phone,
		Password: string(hashedPassword),
		Type:     account.Type,
	}
	err := a.AccountGorm.Create(ctx, accountModels)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) GetUserFromCookie(ctx context.Context, cookie string) (*models.Account, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.SecretKey), nil
	})
	if err != nil {
		a.Logger.Error("Unauthenticated!!", zap.Error(err))
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	acc, err := a.AccountGorm.GetAccountByID(ctx, claims.Issuer)
	if err != nil {
		a.Logger.Error("Can not get Account by ID", zap.Error(err))
		return nil, err
	}
	return acc, nil
}
