package service

import (
	"app/TestYourSpeed/logger"
	"github.com/kylegrantlucas/speedtest"
	"math"
)

var singleSpeedTest *speedtest.Client

func init() {
	var err error
	singleSpeedTest, err = speedtest.NewDefaultClient()
	if err != nil {
		logger.ErrorLog(err)
		panic(err)
	}
}

func TestSpeedTest() (float64, float64, error) {

	server, err := singleSpeedTest.GetServer("")
	if err != nil {
		logger.ErrorLog(err)
		return 0, 0, err
	}

	dmbps, err := singleSpeedTest.Download(server)
	if err != nil {
		logger.ErrorLog(err)
		return 0, 0, err
	}

	umbps, err := singleSpeedTest.Upload(server)
	if err != nil {
		logger.ErrorLog(err)
		return 0, 0, err
	}

	dwnld := math.Floor(dmbps*100) / 100
	upld := math.Floor(umbps*100) / 100
	return dwnld, upld, nil
}
