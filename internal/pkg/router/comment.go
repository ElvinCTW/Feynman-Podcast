package router

import (
	"feynman-podcast/internal/pkg/model/comment"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Comment(r *gin.Engine, client *service.Client) {
	r.GET("/comment/answer/:aid", func(c *gin.Context) {
		answerId := c.Param("aid")

		if list := client.ListComment(answerId); len(*list) == 0 {
			c.String(http.StatusNoContent, http.StatusText(http.StatusNoContent))
			return
		} else {
			c.JSON(http.StatusOK, list)
		}
	})

	r.POST("/comment/answer/:aid", func(c *gin.Context) {
		data := new(comment.Data)
		answerId := c.Param("aid")
		userId := c.Request.Header.Get(Authorization)

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if err := client.CreateComment(answerId, userId, data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})

	r.PUT("/comment/:cid/like", func(c *gin.Context) {
		commentId := c.Param("cid")

		userId := c.Request.Header.Get(Authorization)
		if err := client.UpdateCommentLike(commentId, userId); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.String(http.StatusOK, http.StatusText(http.StatusOK))
		}
	})

	r.DELETE("/comment/:cid", func(c *gin.Context) {
		commentId := c.Param("cid")

		userId := c.Request.Header.Get(Authorization)
		client.DeleteComment(commentId, userId)
		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	})
}
