package environments

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	env := os.Getenv("ENV")
	if env == "UAT" {
		godotenv.Load("environment/UAT.env")
	} else if env == "PROD" {
		godotenv.Load("environment/PROD.env")
	} else {
		godotenv.Load("environment/DEV.env")
	}
}
