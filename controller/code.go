package controller

type MyCode int64

const (
	CodeSuccess       MyCode = 1000
	CodeInvalidParams MyCode = 1001
	InternalError     MyCode = 1002
	ResultNotExist    MyCode = 1003
)

var msgFlags = map[MyCode]string{
	CodeSuccess:       "返回成功",
	CodeInvalidParams: "参数错误",
	InternalError:     "内部错误",
	ResultNotExist:    "结果不存在",
}

func (mc MyCode) Msg() string {
	if msg, ok := msgFlags[mc]; ok {
		return msg
	}
	return "服务繁忙"
}
