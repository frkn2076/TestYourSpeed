package controller

import (
	"app/TestYourSpeed/service"
	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

func (u *ManagerController) Test(c *gin.Context) {
	testerId := c.Param("testerId")
	var dwnld float64
	var upld float64
	if testerId == "1" {
		dwnld, upld, _ = service.TestFast()
	} else if testerId == "2" {
		dwnld, upld, _ = service.TestSpeedTest()
	} else {
		c.AbortWithStatus(400)
	}
	c.JSON(200, gin.H{
		"Download": dwnld,
		"Upload": upld,
	})
}
