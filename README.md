## TestYourSpeed

#### A small application that compares Ookla's speedtest and Netflix's fast
* https://www.speedtest.net
* https://fast.com

You can call http://localhost:8080/testspeed/1 for fast and http://localhost:8080/testspeed/2 for speedtest while running application on your local machine.


#### You can change running port and filelog paths from .env files according to the environment you want to run and can add new environment variables on these files.
* **go run main.go** &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; => DEV
* **ENV=UAT go run main.go** &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;=> UAT
* **ENV=PROD go run main.go** &nbsp; &nbsp; &nbsp; => PROD

#### Or you can decide which envrironment to choose by launch.json of .vscode as,
```json
  "env": {
    "ENV":"UAT"
  }
```
