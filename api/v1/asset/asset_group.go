package asset

import (
	"anew-server/api/v1/system"
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/dto/service"
	system2 "anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetAssetGroups(c *gin.Context) {
	// 绑定参数
	var req request.AssetGroupReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	// 创建服务
	s := service.New()
	groups, err := s.GetAssetGroups(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	// 处理hosts_id
	var newList []response.AssetGroupListResp
	for _, g := range groups {
		var ng response.AssetGroupListResp
		ng.Name = g.Name
		ng.Id = g.Id
		ng.Desc = g.Desc
		ng.Creator = g.Creator
		for _, h := range g.Hosts {
			ng.HostsId = append(ng.HostsId, h.Id)
		}
		newList = append(newList, ng)
	}
	// 返回分页数据
	var resp response.PageData
	// 设置分页参数
	resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.DataList = newList
	response.SuccessWithData(resp)
}

func CreateAssetGroup(c *gin.Context) {
	user := system.GetCurrentUserFromCache(c)
	// 绑定参数
	var req request.CreateAssetGroupReq
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
	// 记录当前创建人信息
	req.Creator = user.(system2.SysUser).Name
	// 创建服务
	s := service.New()
	err = s.CreateAssetGroup(&req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()

}

func UpdateAssetGroupByID(c *gin.Context) {
	// 绑定参数
	var req request.UpdateAssetGroupReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}
	deptId := utils.Str2Uint(c.Param("groupId"))
	if deptId == 0 {
		response.FailWithMsg("分组编号不正确")
		return
	}
	// 创建服务
	s := service.New()
	// 更新数据
	err = s.UpdateAssetGroupById(deptId, req)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}

// 批量删除分组
func BatchDeleteAssetGroupByIds(c *gin.Context) {
	var req request.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithCode(response.ParmError)
		return
	}

	// 创建服务
	s := service.New()
	// 删除数据
	err = s.DeleteAssetGroupByIds(req.Ids)
	if err != nil {
		response.FailWithMsg(err.Error())
		return
	}
	response.Success()
}