package routes

import (
	"github.com/dqfan2012/learngin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/status", handlers.StatusHandler)
	router.GET("/user", handlers.GetUserHandler)
}
