package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(r *gin.Engine) {
	HandleApi(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
