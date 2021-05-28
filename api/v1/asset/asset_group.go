package asset

import (
	request2 "anew-server/api/request"
	response2 "anew-server/api/response"
	"anew-server/api/v1/system"
	service2 "anew-server/dao"
	system2 "anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetAssetGroups(c *gin.Context) {
	// 绑定参数
	var req request2.AssetGroupReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	// 创建服务
	s := service2.New()
	groups, err := s.GetAssetGroups(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	// 处理hosts_id
	var newList []response2.AssetGroupListResp
	for _, g := range groups {
		var ng response2.AssetGroupListResp
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
	var resp response2.PageData
	// 设置分页参数
	resp.PageInfo = req.PageInfo
	// 设置数据列表
	resp.DataList = newList
	response2.SuccessWithData(resp)
}

func CreateAssetGroup(c *gin.Context) {
	user := system.GetCurrentUserFromCache(c)
	// 绑定参数
	var req request2.CreateAssetGroupReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	// 参数校验
	err = common.NewValidatorError(common.Validate.Struct(req), req.FieldTrans())
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	// 记录当前创建人信息
	req.Creator = user.(system2.SysUser).Name
	// 创建服务
	s := service2.New()
	err = s.CreateAssetGroup(&req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()

}

func UpdateAssetGroupByID(c *gin.Context) {
	// 绑定参数
	var req request2.UpdateAssetGroupReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}
	deptId := utils.Str2Uint(c.Param("groupId"))
	if deptId == 0 {
		response2.FailWithMsg("分组编号不正确")
		return
	}
	// 创建服务
	s := service2.New()
	// 更新数据
	err = s.UpdateAssetGroupById(deptId, req)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}

// 批量删除分组
func BatchDeleteAssetGroupByIds(c *gin.Context) {
	var req request2.IdsReq
	err := c.Bind(&req)
	if err != nil {
		response2.FailWithCode(response2.ParmError)
		return
	}

	// 创建服务
	s := service2.New()
	// 删除数据
	err = s.DeleteAssetGroupByIds(req.Ids)
	if err != nil {
		response2.FailWithMsg(err.Error())
		return
	}
	response2.Success()
}