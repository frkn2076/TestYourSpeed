package fast

import (
	"github.com/ddo/go-fast"
)

var singleFastCom *fast.Fast

func init() {
	singleFastCom = fast.New()
}

func Test() float64 {

	if err := singleFastCom.Init(); err != nil {
		panic(err)
	}

	urls, err := singleFastCom.GetUrls()

	if err != nil {
		panic(err)
	}

	KbpsChan := make(chan float64)

	var result float64
	go func() {
		for Kbps := range KbpsChan {
			result = Kbps / 1000
		}
	}()

	if err = singleFastCom.Measure(urls, KbpsChan); err != nil {
		panic(err)
	}

	return result
}
