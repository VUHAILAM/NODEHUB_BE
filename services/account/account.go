package account

import (
	"context"
	"time"

	"gitlab.com/hieuxeko19991/job4e_be/services/recruiter"

	email2 "gitlab.com/hieuxeko19991/job4e_be/services/email"

	"github.com/dgrijalva/jwt-go"

	"gitlab.com/hieuxeko19991/job4e_be/cmd/config"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"

	"gitlab.com/hieuxeko19991/job4e_be/pkg/utils"

	"github.com/pkg/errors"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	Login(ctx context.Context, email string, password string) (string, string, error)
	Logout(ctx context.Context, email string) error
	Register(ctx context.Context, account *models.RequestRegisterAccount) error
	ForgotPassword(ctx context.Context, email string) error
	ChangePassword(ctx context.Context, req *models.RequestChangePassword) error
	ResetPassword(ctx context.Context, token string, newPassword string) error
	GetAccessToken(ctx context.Context, accountID int64, customKey string) (string, error)
	VerifyEmail(ctx context.Context, email string) error
}

type Account struct {
	AccountGorm   *AccountGorm
	RecruiterGorm *recruiter.RecruiterGorm
	Auth          *auth.AuthHandler
	MailService   *email2.SGMailService
	Conf          *config.Config
	Logger        *zap.Logger
}

func NewAccount(accountGorm *AccountGorm, recruiterGorm *recruiter.RecruiterGorm, auth *auth.AuthHandler, conf *config.Config, mailSV *email2.SGMailService, logger *zap.Logger) *Account {
	return &Account{
		AccountGorm:   accountGorm,
		RecruiterGorm: recruiterGorm,
		Auth:          auth,
		Conf:          conf,
		MailService:   mailSV,
		Logger:        logger,
	}
}

func (a *Account) Login(ctx context.Context, email string, password string) (string, string, error) {
	acc, err := a.AccountGorm.GetAccountByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if acc.Id == 0 {
		a.Logger.Error("Account by email not found")
		return "", "", errors.New("Account not found")
	}

	if acc.Type == auth.AdminRole {
		a.Logger.Error("Account is not author")
		return "", "", errors.New("Account is not author")
	}
	a.Logger.Info("Account", zap.Reflect("account", acc))

	if !acc.Status {
		a.Logger.Error("Account not verified")
		return "", "", errors.New("Account not verified")
	}
	reqAccount := &models.Account{
		Email:    email,
		Password: password,
	}
	if valid := a.Auth.Authenticate(reqAccount, acc); !valid {
		a.Logger.Error("Wrong password!!")
		return "", "", errors.New("Wrong password")
	}

	accessToken, err := a.Auth.GenerateAccessToken(acc)
	if err != nil {
		a.Logger.Error("Can not generate Access Token", zap.Error(err))
		return "", "", err
	}

	refreshToken, err := a.Auth.GenerateRefreshToken(acc)
	if err != nil {
		a.Logger.Error("Can not generate Refresh Token", zap.Error(err))
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (a *Account) Logout(ctx context.Context, email string) error {
	return nil
}

type VerifyAccountClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (a *Account) Register(ctx context.Context, account *models.RequestRegisterAccount) error {
	token, err := a.generateVerifyToken(ctx, account.Email)
	if err != nil {
		a.Logger.Error("Cannot gen Verify Email token", zap.Error(err))
		return err
	}
	linkReset := a.Conf.Origin + "verify-email?token=" + token
	a.Logger.Info("Link verify email", zap.String("url", linkReset))

	from := "lamvhhe130764@fpt.edu.vn"
	to := []string{account.Email}
	subject := "Verify Email for NodeHub"
	mailType := email2.MailConfirmation
	mailData := models.MailData{
		Link: linkReset,
	}

	mailReq := a.MailService.NewMail(from, to, subject, mailType, &mailData)
	err = a.MailService.SendMail(mailReq)
	if err != nil {
		a.Logger.Error("Cannot send email", zap.Error(err))
		return err
	}
	var inforID int64 = 0
	if account.Type == auth.RecruiterRole {
		recruiterModel := &models.Recruiter{
			AccountID:        0,
			CompanyName:      account.RecruiterInfor.CompanyName,
			Address:          account.RecruiterInfor.Address,
			Avartar:          account.RecruiterInfor.Avartar,
			Banner:           account.RecruiterInfor.Banner,
			Phone:            account.RecruiterInfor.Phone,
			Website:          account.RecruiterInfor.Website,
			Description:      account.RecruiterInfor.Description,
			EmployeeQuantity: account.RecruiterInfor.EmployeeQuantity,
			ContacterName:    account.RecruiterInfor.ContacterName,
			ContacterPhone:   account.RecruiterInfor.ContacterPhone,
			Media:            account.RecruiterInfor.Media,
		}
		inforID, err = a.RecruiterGorm.Create(ctx, recruiterModel)
		if err != nil {
			a.Logger.Error("Create Recruiter error", zap.Error(err))
			return err
		}
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), 8)
	accountModels := &models.Account{
		Email:         account.Email,
		Phone:         account.Phone,
		Password:      string(hashedPassword),
		Status:        false,
		TokenHash:     utils.GenerateRandomString(15),
		Type:          account.Type,
		InformationID: inforID,
	}
	_, err = a.AccountGorm.Create(ctx, accountModels)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) generateVerifyToken(ctx context.Context, email string) (string, error) {
	claims := ResetPasswordClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(a.Conf.VerifyEmailExpiration)).Unix(),
		},
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(a.Conf.AccessTokenPrivateKey))
	if err != nil {
		a.Logger.Error("unable to parse private key", zap.Error(err))
		return "", errors.New("could not generate access token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

func (a *Account) ForgotPassword(ctx context.Context, email string) error {
	acc, err := a.AccountGorm.GetAccountByEmail(ctx, email)
	if err != nil {
		a.Logger.Error("Verify Account exist", zap.Error(err))
		return err
	}

	if acc.Id == 0 {
		a.Logger.Error("Account not existed")
		return errors.New("Account " + email + " not existed")
	}

	token, err := a.generateResetToken(ctx, email)
	if err != nil {
		a.Logger.Error("Cannot gen Reset password token", zap.Error(err))
		return err
	}
	linkReset := a.Conf.Origin + "reset-password?token=" + token
	a.Logger.Info("Link reset password", zap.String("url", linkReset))

	from := "lamvhhe130764@fpt.edu.vn"
	to := []string{email}
	subject := "Password Reset for NodeHub"
	mailType := email2.PassReset
	mailData := models.MailData{
		Link: linkReset,
	}

	mailReq := a.MailService.NewMail(from, to, subject, mailType, &mailData)
	err = a.MailService.SendMail(mailReq)
	if err != nil {
		a.Logger.Error("Cannot send email", zap.Error(err))
		return err
	}
	return nil
}

type ResetPasswordClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (a *Account) generateResetToken(ctx context.Context, email string) (string, error) {
	claims := ResetPasswordClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(a.Conf.ResetPasswordExpiration)).Unix(),
		},
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(a.Conf.AccessTokenPrivateKey))
	if err != nil {
		a.Logger.Error("unable to parse private key", zap.Error(err))
		return "", errors.New("could not generate access token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

func (a *Account) ChangePassword(ctx context.Context, req *models.RequestChangePassword) error {
	acc, err := a.AccountGorm.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		a.Logger.Error("Get Account error", zap.Error(err))
		return err
	}

	if acc.Id == 0 {
		a.Logger.Error("Account not existed")
		return errors.New("Account " + req.Email + " not existed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.OldPassword)); err != nil {
		a.Logger.Error("Password not same", zap.Error(err))
		return err
	}

	hashNewPassword, _ := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 8)
	tokenHash := utils.GenerateRandomString(15)
	err = a.AccountGorm.UpdatePassword(ctx, req.Email, string(hashNewPassword), tokenHash)
	if err != nil {
		a.Logger.Error("Update Password error", zap.Error(err))
		return err
	}
	return nil
}

func (a *Account) ResetPassword(ctx context.Context, token string, newPassword string) error {
	jwt, err := jwt.ParseWithClaims(token, &ResetPasswordClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			a.Logger.Error("Unexpected signing method in reset token")
			return nil, errors.New("Unexpected signing method in reset token")
		}
		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(a.Conf.AccessTokenPublicKey))
		if err != nil {
			a.Logger.Error("unable to parse public key", zap.Error(err))
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		a.Logger.Error("unable to parse Reset Password claims", zap.Error(err))
		return err
	}
	claims, ok := jwt.Claims.(*ResetPasswordClaims)
	if !jwt.Valid || !ok || claims.Email == "" {
		return errors.New("invalid token: get Reset token failed")
	}

	hashNewPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 8)
	tokenHash := utils.GenerateRandomString(15)
	err = a.AccountGorm.UpdatePassword(ctx, claims.Email, string(hashNewPassword), tokenHash)
	if err != nil {
		a.Logger.Error("Update Password error", zap.Error(err))
		return err
	}
	return nil
}

func (a *Account) GetAccessToken(ctx context.Context, accountID int64, customKey string) (string, error) {
	account, err := a.AccountGorm.GetAccountByID(ctx, accountID)
	if err != nil {
		a.Logger.Error("Can not use AccountID to get", zap.Error(err), zap.Int64("account_id", accountID))
		return "", err
	}
	actualCustomKey := a.Auth.GenerateCustomKey(string(account.Id), account.TokenHash)
	if customKey != actualCustomKey {
		a.Logger.Error("Wrong token: Authentication failed")
		return "", errors.New("Authentication failed. Invalid token")
	}

	accessToken, err := a.Auth.GenerateAccessToken(account)
	if err != nil {
		a.Logger.Error("Can not Get Access Token", zap.Error(err))
		return "", err
	}

	return accessToken, nil
}

func (a *Account) VerifyEmail(ctx context.Context, email string) error {
	err := a.AccountGorm.UpdateVerifyEmail(ctx, email)
	if err != nil {
		a.Logger.Error("Update verify email failed", zap.Error(err), zap.String("email", email))
		return err
	}
	return nil
}
