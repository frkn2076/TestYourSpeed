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
	router.GET("/test/:testerID", manager.Test)
	router.NoRoute(func(c *gin.Context) {
		c.Data(404, "text/plain", []byte("Call \".../test/1\" for fast\nCall \".../test/2\" for testspeed"))
	})

	return router
}
