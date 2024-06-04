package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dqfan2012/learngin/internal/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/home", handlers.HomeHandler)
	router.GET("/status", handlers.StatusHandler)
	router.GET("/user", handlers.GetUserHandler)
}
