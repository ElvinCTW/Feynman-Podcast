package router

import (
	"feynman-podcast/internal/pkg/model/question"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Question(r *gin.Engine, client *service.Client) {
	r.POST("/question", func(c *gin.Context) {
		data := new(question.Data)

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if err := client.CreateQuestion(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})
}
