package main

import (
	"feynman-podcast/internal/pkg/config"
	"feynman-podcast/internal/pkg/router"
	"feynman-podcast/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.New()
	config.ReadYaml()
	config.ReadEnv()
	config.Log()

	gin.SetMode(config.GinMode)
	r := gin.Default()

	// apply router
	router.MainRouter(r, service.NewClient(config))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(config.Port)
}
