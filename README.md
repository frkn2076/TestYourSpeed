## TestYourSpeed

#### A small application that compares download and upload speed tests of Ookla's speedtest and Netflix's fast
* https://www.speedtest.net
* https://fast.com

You can call http://localhost:8080/test/1 for fast and http://localhost:8080/test/2 for speedtest while running application on your local machine.


#### You can change running port and filelog paths from .env files according to the environment you want to run and can add new environment variables on these files.
* **go run main.go**&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;=> DEV
* **ENV=UAT go run main.go**&nbsp;&nbsp; &nbsp; &nbsp;=> UAT
* **ENV=PROD go run main.go**&nbsp; &nbsp;=> PROD
* **go test**&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;=> To run unit tests

#### Or you can decide which envrironment to choose by launch.json of .vscode as,
```json
  "env": {
    "ENV":"UAT"
  }
```
&nbsp;

<p align="center">
  <img src="https://github.com/frkn2076/TestYourSpeed/blob/main/image.png">
</p>
&nbsp;

* Used gin as a web framework
* Used testify for assertions
