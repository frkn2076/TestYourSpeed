package service

import (
	"github.com/ddo/go-fast"
	"math"
)

var singleFastCom *fast.Fast

func init() {
	singleFastCom = fast.New()
}

func TestFast() (float64, float64, error) {

	if err := singleFastCom.Init(); err != nil {
		return 0, 0, err
	}

	urls, err := singleFastCom.GetUrls()

	if err != nil {
		return 0, 0, err
	}

	KbpsChan := make(chan float64)

	var result float64
	go func() {
		for Kbps := range KbpsChan {
			result = Kbps / 1000
		}
	}()

	if err = singleFastCom.Measure(urls, KbpsChan); err != nil {
		return 0, 0, err
	}

	result = math.Floor(result*100) / 100
	return result, 0, nil
}
