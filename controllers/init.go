package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/universalmacro/auth/services"
	"github.com/universalmacro/common/server"
)

var authService = services.GetAuthService()
var router *gin.Engine

func Init(addr ...string) {
	router = gin.Default()
	router.Use(authService.Auth())
	router.Use(server.CorsMiddleware())
	var authApi AuthApi
	server.MetricsMiddleware(router)
	router.GET("/me", authApi.GetMe)
	router.POST("/sessions", authApi.CreateSession)
	router.POST("/accounts", authApi.CreateAccount)
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "0.0.2",
		})
	})
	router.Run(addr...)
}
