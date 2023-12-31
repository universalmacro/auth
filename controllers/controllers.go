package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/auth/controllers/models"
	"github.com/universalmacro/auth/services"
	"github.com/universalmacro/common/fault"
)

type AuthApi struct{}

func (AuthApi) CreateSession(ctx *gin.Context) {
	var createSessionRequest models.CreateSessionRequest
	ctx.ShouldBindJSON(&createSessionRequest)
	token, err := authService.CreateSession(createSessionRequest.Email, createSessionRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "")
		return
	}
	ctx.JSON(http.StatusCreated, &models.Session{
		Token: token,
	})
}

func (AuthApi) GetMe(ctx *gin.Context) {
	account := GetAccount(ctx)
	if account == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, models.Account{
		Email: account.Email(),
		Role:  string(account.Role()),
	})
}

func (AuthApi) CreateAccount(ctx *gin.Context) {
	account := GetAccount(ctx)
	if account == nil {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return
	}
	if account.Role() != "ROOT" {
		fault.GinHandler(ctx, fault.ErrPermissionDenied)
		return
	}
	var createAccountRequest models.CreateAccountRequest
	ctx.ShouldBindJSON(&createAccountRequest)
	newAccount, err := authService.CreateAccount(createAccountRequest.Email, createAccountRequest.Password)
	if err != nil {
		fault.GinHandler(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, models.Account{
		Email: newAccount.Email(),
		Role:  string(newAccount.Role()),
	})
}

func GetAccount(ctx *gin.Context) *services.Account {
	accountInterface, ok := ctx.Get("account")
	if !ok {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return nil
	}
	account, ok := accountInterface.(services.Account)
	if !ok {
		fault.GinHandler(ctx, fault.ErrUnauthorized)
		return nil
	}
	return &account
}
