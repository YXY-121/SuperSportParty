package pkg

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}
type response struct {
	RetCode int         `json:"ret_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpcode int, errCode int, data interface{}) {
	g.C.JSON(httpcode, response{
		RetCode: errCode,
		ErrMsg:  GetMessage(errCode),
		Data:    data,
	})
}
