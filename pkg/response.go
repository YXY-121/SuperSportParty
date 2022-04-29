package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Gin struct {
// 	C *gin.Context
// }
type response struct {
	RetCode int         `json:"ret_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, errCode int, data interface{}) {
	c.JSON(http.StatusOK, response{
		RetCode: errCode,
		ErrMsg:  GetMessage(errCode),
		Data:    data,
	})
}
