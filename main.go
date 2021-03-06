package main

import (
	_ "app/TestYourSpeed/environment"
	"app/TestYourSpeed/router"
	"os"
)

func main() {
	r := router.SetupRouter()
	r.Run(os.Getenv("Port"))
}
