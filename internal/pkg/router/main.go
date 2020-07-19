package router

import (
	"feynman-podcast/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

const (
	Authorization = "Authorization"
)

func MainRouter(r *gin.Engine, client *service.Client) {
	User(r, client)

	Question(r, client)

	Answer(r, client)

	Comment(r, client)

	Civil(r, client)
}
