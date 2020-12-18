package router

//
//import (
//	"feynman-podcast/internal/pkg/service"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//const (
//	answer = "answer"
//)
//
//func Answer(r *gin.Engine, client *service.Client) {
//	r.GET("/answer/question/:qid", func(c *gin.Context) {
//		questionId := c.Param("qid")
//
//		if list := client.ListAnswer(questionId); len(*list) == 0 {
//			c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
//			return
//		} else {
//			c.JSON(http.StatusOK, list)
//		}
//	})
//
//	r.POST("/answer/question/:qid", func(c *gin.Context) {
//		questionId := c.Param("qid")
//		userId := c.Request.Header.Get(Authorization)
//
//		if form, err := c.MultipartForm(); err != nil {
//			c.String(http.StatusBadRequest, err.Error())
//			return
//		} else if err := client.CreateAnswer(questionId, userId, form.File[answer]); err != nil {
//			c.String(http.StatusBadRequest, err.Error())
//			return
//		} else {
//			c.String(http.StatusOK, http.StatusText(http.StatusOK))
//		}
//	})
//
//	r.PUT("/answer/:aid/like", func(c *gin.Context) {
//		answerId := c.Param("aid")
//
//		userId := c.Request.Header.Get(Authorization)
//		if err := client.UpdateAnswerLike(answerId, userId); err != nil {
//			c.String(http.StatusBadRequest, err.Error())
//			return
//		} else {
//			c.String(http.StatusOK, http.StatusText(http.StatusOK))
//		}
//
//	})
//}
