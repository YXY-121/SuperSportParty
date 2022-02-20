package pkg

const (
	ErrorParams        = 400
	Success            = 200
	Error              = 500
	ErrorUserHaveExist = 10000
	ErrorPhoneType     = 10001
	ErrorTokenNoExist  = 10002
	ErrorTokenInvalid  = 10003
)

var MsgMap = map[int]string{
	ErrorParams:        "参数错误",
	Success:            "ok",
	Error:              "fail",
	ErrorUserHaveExist: "用户已存在",
	ErrorPhoneType:     "手机格式错误",
	ErrorTokenNoExist:  "token 不存在",
	ErrorTokenInvalid:  "token 错误",
}

func GetMessage(code int) string {
	return MsgMap[code]
}
