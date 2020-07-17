package router

import (
	"feynman-podcast/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, client *service.Client) {
	r.POST("/user", func(c *gin.Context) {

	})

	r.GET("/user/:uid", func(c *gin.Context) {
		// c.JSON(http.StatusOK, )
	})
}
