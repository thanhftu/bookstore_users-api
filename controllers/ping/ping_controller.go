package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// Ping Test the status of app
	c.String(http.StatusOK, "pong")
}
