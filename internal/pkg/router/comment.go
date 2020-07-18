package router

import (
	"feynman-podcast/internal/pkg/model/question"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Comment(r *gin.Engine, client *service.Client) {
	r.POST("/comment/voiceanswer/:vid", func(c *gin.Context) {
		data := new(question.Comment)
		voiceAnswerId := c.Param("vid")

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if err := client.CreateComment(voiceAnswerId, data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})
}
