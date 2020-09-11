package response

const (
	SUCCSE              = 200
	FAILURE             = 400
	FORBIDDEN           = 403
	TOKEN_UNAUTHORIZED  = 1000
	ERROR_USERNAME_USED = 1001
	ERROR_LOGIN_WRONG   = 1002
)

var Custom_Msg = map[int]string{
	SUCCSE:              "操作完成",
	FAILURE:             "操作失败",
	FORBIDDEN:           "无权操作",
	TOKEN_UNAUTHORIZED:  "未经授权",
	ERROR_USERNAME_USED: "用户已存在",
	ERROR_LOGIN_WRONG:   "用户名或密码错误",
}
