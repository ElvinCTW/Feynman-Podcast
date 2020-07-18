package router

import (
	"feynman-podcast/internal/pkg/model/user"
	"feynman-podcast/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, client *service.Client) {
	r.GET("/user", func(c *gin.Context) {
		data := new(user.Data)

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, client.GetUserWithPassword(data.Email, data.PassWord))
		}
	})

	r.GET("/user/:email", func(c *gin.Context) {
		email := c.Param("email")

		c.JSON(http.StatusOK, client.GetUser(email))
	})

	r.POST("/user", func(c *gin.Context) {
		data := new(user.Data)

		if err := c.BindJSON(data); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		} else if d := client.GetUser(data.Email); d != nil {
			c.String(http.StatusConflict, http.StatusText(http.StatusConflict))
		} else {
			c.JSON(http.StatusOK, client.CreateUser(data))
		}
	})
}
