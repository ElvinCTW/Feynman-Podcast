package router

import (
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	voiceAnswer = "voiceAnswer"
)

func VoiceAnswer(r *gin.Engine, client *service.Client) {
	r.GET("/voiceanswer/question/:qid", func(c *gin.Context) {
		questionId := c.Param("qid")

		if list := client.ListVoiceAnswer(questionId); len(*list) == 0 {
			c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
			return
		} else {
			c.JSON(http.StatusOK, list)
		}
	})

	r.POST("/voiceanswer/question/:qid", func(c *gin.Context) {
		questionId := c.Param("qid")
		userId := c.Request.Header.Get(Authorization)

		if form, err := c.MultipartForm(); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if err := client.CreateVoiceAnswer(questionId, userId, form.File[voiceAnswer]); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})

	r.PUT("/voiceanswer/:vid/like", func(c *gin.Context) {
		voiceAnswerId := c.Param("vid")

		userId := c.Request.Header.Get(Authorization)
		if err := client.UpdateVoiceAnswerLike(voiceAnswerId, userId); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}

	})
}
