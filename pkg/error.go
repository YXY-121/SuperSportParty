package pkg

const (
	ErrorParams             = 400
	Success                 = 200
	Error                   = 500
	ErrorUserHaveExist      = 10000
	ErrorPhoneType          = 10001
	ErrorTokenNoExist       = 10002
	ErrorTokenInvalid       = 10003
	ErrorCreateFail         = 10004
	ErrorDeleteFail         = 10005
	ErrorUpdateFail         = 10006
	ErrorListFail           = 10007
	ErrorPageFail           = 10008
	ErrorListOrderFail      = 10009
	ErrorListEvaluationFail = 10010
	ErrorSaveImageFail      = 10011
)

var MsgMap = map[int]string{
	ErrorParams:             "参数错误",
	Success:                 "ok",
	Error:                   "fail",
	ErrorUserHaveExist:      "用户已存在",
	ErrorPhoneType:          "手机格式错误",
	ErrorTokenNoExist:       "token 不存在",
	ErrorTokenInvalid:       "token 错误",
	ErrorCreateFail:         "创建失败",
	ErrorDeleteFail:         "删除失败",
	ErrorUpdateFail:         "修改失败",
	ErrorListFail:           "查询失败",
	ErrorPageFail:           "分页失败",
	ErrorListEvaluationFail: "获取评价失败",
	ErrorSaveImageFail:      "保存图片失败",
}

func GetMessage(code int) string {
	return MsgMap[code]
}
