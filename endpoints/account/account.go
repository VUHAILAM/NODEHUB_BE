package account

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/hieuxeko19991/job4e_be/services/account"

	"github.com/gin-gonic/gin"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/ginx"
	"gitlab.com/hieuxeko19991/job4e_be/pkg/models"
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
	token, err := as.accountService.Login(ctx, req.Email, req.Password)
	if err != nil {
		as.Logger.Error("Login error", zap.Error(err))
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}
	ginCtx.SetCookie(cookieName, token, time.Now().Add(time.Hour*24).Second(), "/", "", true, true)

	ginx.BuildSuccessResponse(ginCtx, http.StatusAccepted, gin.H{
		"token": token,
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

func (as *AccountSerializer) GetUserFromCookie(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	cookie, err := ginCtx.Cookie(cookieName)
	if err != nil {
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	acc, err := as.accountService.GetUserFromCookie(ctx, cookie)

	if err != nil {
		ginx.BuildErrorResponse(ginCtx, err, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginx.BuildSuccessResponse(ginCtx, http.StatusOK, acc)
}
