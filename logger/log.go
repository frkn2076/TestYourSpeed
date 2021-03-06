package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var singleInfo *log.Logger
var singleError *log.Logger

func InfoLog(logText ...interface{}) {
	singleInfo.Println(logText)
}

func ErrorLog(logText ...interface{}) {
	singleError.Println(logText)
}

func init() {
	singleInfo = initLogger("InfoLog")
	singleError = initLogger("ErrorLog")
}

func initLogger(folderName string) *log.Logger {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	dt := time.Now()
	today := dt.Format("02-Jan-2006")

	folderPath := path.Join(wd, "logger", folderName)

	//check folder created before
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0700)
	}

	fileName := path.Join(folderPath, today)

	//check log file created before
	_, err = os.Stat(fileName)
	fileNotExist := os.IsNotExist(err)

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "* ", log.LstdFlags)
	if fileNotExist {
		logger.Println(fmt.Sprintf("%s has created", folderName))
	}
	return logger
}
