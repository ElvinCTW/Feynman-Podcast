package router

import (
	"feynman-podcast/internal/pkg/model/question"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Question(r *gin.Engine, client *service.Client) {
	r.GET("/question/:qid", func(c *gin.Context) {
		questionId := c.Param("qid")

		if q := client.GetQuestion(questionId); q == nil {
			c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
			return
		} else {
			c.JSON(http.StatusOK, q)
		}
	})

	r.POST("/question", func(c *gin.Context) {
		data := new(question.Question)

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if id, err := client.CreateQuestion(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, *id)
		}
	})
}
