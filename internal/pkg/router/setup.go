package router

import (
	"feynman-podcast/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func MainRouter(r *gin.Engine, client *service.Client) {
	User(r, client)
}
