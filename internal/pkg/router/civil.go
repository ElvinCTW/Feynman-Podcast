package router

import (
	"feynman-podcast/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

func Civil(r *gin.Engine, client *service.Client) {
	//r.GET("/civil/question/parse", func(c *gin.Context) {
	//	if err := client.CrawlerClient.ConvertPDF(); err != nil {
	//		c.String(http.StatusBadRequest, err.Error())
	//		return
	//	} else {
	//		c.String(http.StatusOK, http.StatusText(http.StatusOK))
	//	}
	//})

	r.GET("/civil/question/latest", func(c *gin.Context) {

	})

	r.GET("civil/question/hottest", func(c *gin.Context) {

	})
}
