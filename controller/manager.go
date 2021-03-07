package controller

import (
	"app/TestYourSpeed/logger"
	"app/TestYourSpeed/service"
	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

func (u *ManagerController) Test(c *gin.Context) {
	testerID := c.Param("testerID")
	var dwnld float64
	var upld float64
	if testerID == "1" {
		dwnld, upld, _ = service.TestFast()
	} else if testerID == "2" {
		dwnld, upld, _ = service.TestSpeedTest()
	} else {
		logger.ErrorLog("Unexpected testerID")
		c.AbortWithStatus(400)
	}
	logger.InfoLog("Download: ", dwnld, " Upload: ", upld)
	c.JSON(200, gin.H{
		"Download": dwnld,
		"Upload":   upld,
	})
}
