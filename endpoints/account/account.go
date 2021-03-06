package account

import (
	"encoding/json"
	"net/http"

	"gitlab.com/hieuxeko19991/job4e_be/models"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/auth"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/services/account"
	"go.uber.org/zap"
)

const cookieName = "account"

type AccountSerializer struct {
	accountService account.IAccountService
	Logger         *zap.Logger
}

func NewAccountSerializer(accountService account.IAccountService, logger *zap.Logger) *AccountSerializer {
	return &AccountSerializer{
		accountService: accountService,
		Logger:         logger,
	}
}

func (as *AccountSerializer) Login(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestLogin{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		as.Logger.Error("Parse request Login error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	as.Logger.Info("Request", zap.Reflect("request", req))
	accessToken, refreshToken, err := as.accountService.Login(ctx, req.Email, req.Password, req.Type)
	if err != nil {
		as.Logger.Error("Login error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, &auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (as *AccountSerializer) Logout(ginCtx *gin.Context) {
	ginCtx.SetCookie(cookieName, "", -1, "/", "", true, true)
	ginx.BuildSuccessResponse(ginCtx, http.StatusOK, nil)
}

func (as *AccountSerializer) Register(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestRegisterAccount{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)

	if err != nil {
		as.Logger.Error("Parse request Register account error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = as.accountService.Register(ctx, &req)
	if err != nil {
		as.Logger.Error("Register account error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (as *AccountSerializer) ForgotPassword(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestForgotPassword{}

	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		as.Logger.Error("Parse request Forgot Password account error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = as.accountService.ForgotPassword(ctx, req.Email)
	if err != nil {
		as.Logger.Error("Forgot Password error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (as *AccountSerializer) ChangePassword(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		as.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	req := models.RequestChangePassword{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		as.Logger.Error("Parse request Change Password account error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.Email = acc.(models.Account).Email
	err = as.accountService.ChangePassword(ctx, &req)
	if err != nil {
		as.Logger.Error("Change password error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (as *AccountSerializer) ResetPassword(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestResetPassword{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		as.Logger.Error("Parse request Change Password account error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = as.accountService.ResetPassword(ctx, req.Token, req.NewPassword)
	if err != nil {
		as.Logger.Error("Reset password error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}

func (as *AccountSerializer) GetAccessToken(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	acc, ok := ginCtx.Get(auth.AccountKey)
	if !ok {
		as.Logger.Error("Can not get account infor from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	as.Logger.Info("acc", zap.Reflect("acc", acc))
	customKey, ok := ginCtx.Get(auth.VerificationDataKey)
	if !ok {
		as.Logger.Error("Can not get customKey from context")
		ginx.BuildErrorResponse(ginCtx, errors.New("Can not get account infor from context"), gin.H{
			"message": "Can not get account infor from context",
		})
		return
	}
	id := acc.(models.Account).Id
	as.Logger.Info("id", zap.Reflect("id", id))
	accessToken, err := as.accountService.GetAccessToken(ctx, id, customKey.(string))
	if err != nil {
		as.Logger.Error("Can not get access token", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"access_token": accessToken,
	})
}

func (as *AccountSerializer) VerifyEmail(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	req := models.RequestVerifyEmail{}
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {
		as.Logger.Error("Parse request Change Password account error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = as.accountService.VerifyEmail(ctx, req.Email)
	if err != nil {
		as.Logger.Error("Verify email error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, nil)
}
