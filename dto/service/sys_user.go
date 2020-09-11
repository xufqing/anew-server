package service

import (
	"anew-server/common"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/models"
	"anew-server/utils"
)

// 用户是否存在
func CheckUser(username string) int {
	var users models.SysUser
	common.Mysql.Select("id").Where("username = ?", username).First(&users)
	if users.Id > 0 {
		return response.ERROR_USERNAME_USED
	}
	return response.SUCCSE
}

// 创建用户
func CreateUser(req *request.CreateUserReq) int {
	var user models.SysUser
	utils.Struct2StructByJson(req, &user)
	// 将密码转为密文
	user.Password = utils.GenPwd(req.Password)
	err := common.Mysql.Create(&user)
	if err != nil {
		return response.FAILURE
	}
	return response.SUCCSE
}

// 用户列表
func GetUsers(pageSize int, pageNum int) []models.SysUser {
	var users []models.SysUser
	err := common.Mysql.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}
