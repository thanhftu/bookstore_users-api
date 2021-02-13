package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhftu/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

// StartApplication initiate app
func StartApplication() {
	mapURL()
	logger.Info("application is initating...")
	router.Run(":8081")
}
