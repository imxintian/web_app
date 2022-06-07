package controller

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserNotExist
	CodeUserNotFound
	CodeUserExist
	codeInvalidPassword
	CodeUserPasswordError
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "invalid param",
	CodeUserNotExist:      "user not exist",
	CodeUserExist:         "user exist",
	CodeUserNotFound:      "user not found",
	codeInvalidPassword:   "invalid password",
	CodeUserPasswordError: "user password error",
	CodeServerBusy:        "server busy",
}

func Msg(c ResCode) string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
