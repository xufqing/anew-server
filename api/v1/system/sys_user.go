package system

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	"anew-server/models"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)


// 获取当前请求用户信息
func GetCurrentUser(c *gin.Context) (models.SysUser, []uint){
	user, exists := c.Get("user")
	var newUser models.SysUser
	if !exists {
		return newUser, []uint{}
	}
	u, _ := user.(models.SysUser)
	// 创建服务
	s := service.New(c)
	newUser, _ = s.GetUserById(u.Id)
	// 返回roles的id格式化
	roleIds := make([]uint, 0)
	for _, role := range newUser.Roles {
		roleIds = append(roleIds, role.Id)
	}
	return newUser, roleIds
}

// 创建用户
func CreateUser(c *gin.Context) {
	user,_ := GetCurrentUser(c)
	// 绑定参数
	var req request.CreateUserReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 参数校验
	err = common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
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
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	users, err := s.GetUsers(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	var respStruct []response.UserListResp
	for _,user := range users{
		// 把新增的key和title赋值
		var item response.UserListResp
		utils.Struct2StructByJson(user, &item)
		newRole := make([]response.UserRolesResp,0)
		for _,r := range item.Roles {
			r.Key = fmt.Sprintf("%d",r.Id)
			r.Title = r.Name
			newRole = append(newRole,r)
		}
		item.Roles = newRole
		respStruct = append(respStruct,item)
	}
	// 返回分页数据
	var resp response.PageData
	// 设置分页参数
	resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.DataList = respStruct
	response.SuccessWithData(resp)
}

// 更新用户
func UpdateUserById(c *gin.Context) {
	// 绑定参数
	var req request.UpdateUserReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	// 参数校验
	err = common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 获取url path中的userId
	userId := utils.Str2Uint(c.Param("userId"))
	if userId == 0 {
		response.FailWithMsg("用户编号不正确")
		return
	}
	// 创建服务
	s := service.New(c)
	// 更新数据
	err = s.UpdateUserById(userId, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 修改密码
func ChangePwd(c *gin.Context) {
	var msg string
	// 请求json绑定
	var req request.ChangePwdReq
	_ = c.ShouldBindJSON(&req)
	// 参数校验
	err := common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 获取当前用户
	user,_ := GetCurrentUser(c)
	query := common.Mysql.Where("username = ?", user.Username).First(&user)
	// 查询用户
	err = query.Error
	if err != nil {
		msg = err.Error()
	} else {
		// 校验密码
		if ok := utils.ComparePwd(req.OldPassword, user.Password); !ok {
			msg = "原密码错误"
		} else {
			// 更新密码
			err = query.Update("password", utils.GenPwd(req.NewPassword)).Error
			if err != nil {
				msg = err.Error()
			}
		}
	}
	if msg != "" {
		response.FailWithMsg(msg)
		return
	}
	response.Success()
}

// 批量删除用户
func DeleteUserByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New(c)
	// 删除数据
	err = s.DeleteUserByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}
