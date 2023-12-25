package routes

import (
	"go-gin-boilerplate/api/controllers"

	"github.com/gin-gonic/gin"
)

func Logs(routerGroup *gin.RouterGroup) {
	controller := new(controllers.LogsController)

	routerGroup.POST("", controller.Index)
}
