package v1

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/utils"
	"github.com/gin-gonic/gin"
)


// 获取当前请求用户信息
func GetCurrentUser(c *gin.Context) models.SysUser {
	user, exists := c.Get("user")
	var newUser models.SysUser
	if !exists {
		return newUser
	}
	u, _ := user.(models.SysUser)
	// 创建服务
	s := service.New(c)
	newUser, _ = s.GetUserById(u.Id)
	return newUser
}

// 创建用户
func CreateUser(c *gin.Context) {
	user := GetCurrentUser(c)
	// 绑定参数
	var req request.CreateUserReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}

	// 记录当前创建人信息
	req.Creator = user.Name
	// 创建服务
	s := service.New(c)
	err = s.CreateUser(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 获取用户列表
func GetUsers(c *gin.Context) {
	// 绑定参数
	var req request.UserListReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMsg("参数绑定失败, 请检查数据类型")
		return
	}

	// 创建服务
	s := service.New(c)
	users, err := s.GetUsers(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 转为ResponseStruct, 隐藏部分字段
	var respStruct []response.UserListResp
	utils.Struct2StructByJson(users, &respStruct)
	response.SuccessWithData(respStruct)
}

// 更新用户
func UpdateUserById(c *gin.Context) {
	// 绑定参数
	var req gin.H
	var pwd request.ChangePwdReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMsg("参数绑定失败, 请检查数据类型")
		return
	}

	// 将部分参数转为pwd, 如果值不为空, 可能会用到
	utils.Struct2StructByJson(req, &pwd)
	// 获取url path中的userId
	userId := utils.Str2Uint(c.Param("userId"))
	if userId == 0 {
		response.FailWithMsg("用户编号不正确")
		return
	}
	// 创建服务
	s := service.New(c)
	// 更新数据
	err = s.UpdateUserById(userId, pwd.NewPassword, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 批量删除用户
func DeleteUserByIds(c *gin.Context) {
	var req request.Req
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMsg("参数绑定失败, 请检查数据类型")
		return
	}

	// 创建服务
	s := service.New(c)
	// 删除数据
	err = s.DeleteUserByIds(req.GetUintIds())
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
