package router

import (
	"feynman-podcast/internal/pkg/model/question"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VoiceAnswer(r *gin.Engine, client *service.Client) {
	r.POST("/voiceanswer", func(c *gin.Context) {
		data := new(question.VoiceAnswer)

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if err := client.CreateVoiceAnswer(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})

	r.POST("/voiceanswer/:vid/comment", func(c *gin.Context) {
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
