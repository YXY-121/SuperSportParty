package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogToGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			logrus.Errorf(err.Error())
		}
		//很关键
		//把读过的字节流重新放到body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		logrus.Debugf("requestbody is [%v],requestUrl is [%v],requestHost is [%v],time is [%v]/n",
			c.Request.Body, c.Request.URL, c.Request.Host)
	}

}
