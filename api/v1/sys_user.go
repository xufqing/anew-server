package v1

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/utils"
	"github.com/gin-gonic/gin"
)

// 创建用户
func CreateUser(c *gin.Context) {
	// 请求校验
	var user request.CreateUserReq
	err := c.ShouldBindJSON(&user)
	if err != nil{
		response.Resp(c, 400,user)
		return
	}
	code := service.CheckUser(user.Username)
	if code == response.SUCCSE {
		code = service.CreateUser(&user)
		response.Resp(c, code,user)
	} else {
		response.Resp(c, code,user)
	}
}

// 获取用户列表
func GetUsers(c *gin.Context) {
	pageSize := utils.Str2Int(c.Query("pagesize"))
	pageNum := utils.Str2Int(c.Query("pagenum"))
	users := service.GetUsers(pageSize, pageNum)
	// 转为Resp Struct, 隐藏部分字段
	var respUsers []response.UserListResp
	utils.Struct2StructByJson(users, &respUsers)
	response.Resp(c, response.SUCCSE, respUsers)
}
