package middleware

import (
	"app/TestYourSpeed/logger"
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ServiceLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// RequestLog
		body, _ := ioutil.ReadAll(c.Request.Body)
		ipAddress := c.ClientIP()
		if ipAddress == "::1" { //::1 means localhost
			ipAddress = "localhost"
		}
		logger.ServiceLog("Request : ", " IP: ", ipAddress, c.Request.RequestURI, " Body: ", string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		//ResponseLog
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()
		responseBody := bodyLogWriter.body.String()
		logger.ServiceLog("Response: ", " IP: ", ipAddress, c.Request.RequestURI, " Output: ", string(responseBody))
	}
}
