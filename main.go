package main

import (
	"fmt"
	"go-gin-boilerplate/api/config"
	"go-gin-boilerplate/api/logging"
	"go-gin-boilerplate/api/lrucache"
	"go-gin-boilerplate/api/middlewares"
	"go-gin-boilerplate/api/routes"
	"go-gin-boilerplate/api/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.CreateDirectories([]string{"backend", "logs"})
	logging.InitializeLogger()
	lrucache.InitializeLruCache()
	lrucache.Cache.Add("config", config.GetConfig())
}

func main() {
	var conf config.Config

	if cachedConfig, ok := lrucache.Cache.Get("config"); ok {
		conf = cachedConfig.(config.Config)
	} else {
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(middlewares.CorsMiddleware())

	routerGroup := router.Group("")
	routes.Logs(routerGroup.Group("/logs"))

	router.Run(fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port))
}
