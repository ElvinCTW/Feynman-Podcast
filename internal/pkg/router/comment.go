package router

import (
	"feynman-podcast/internal/pkg/model/comment"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Comment(r *gin.Engine, client *service.Client) {
	r.POST("/comment/answer/:aid", func(c *gin.Context) {
		data := new(comment.Data)
		answerId := c.Param("aid")

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if err := client.CreateComment(answerId, data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})
}
