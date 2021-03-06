package router

import (
	"app/TestYourSpeed/controller"
	"app/TestYourSpeed/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ServiceLog())

	manager := new(controller.ManagerController)
	router.GET("/testspeed/:testerId", manager.Test)

	return router
}
