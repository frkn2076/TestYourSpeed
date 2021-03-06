package controller

import (
	"app/TestYourSpeed/service"
	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

func (u *ManagerController) Test(c *gin.Context) {
	testerId := c.Param("testerId")
	var res float64
	if testerId == "1" {
		res, _ = service.FastDownloadTest()
	} else if testerId == "2" {
		res, _ = service.SpeedTestDownloadTest()
	} else {
		c.AbortWithStatus(400)
	}
	c.JSON(200, res)
}
