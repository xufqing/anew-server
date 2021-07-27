package dao

import (
	"anew-server/api/request"
	"anew-server/api/response"
	"anew-server/models/system"
	"anew-server/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

func (s *MysqlService) GetApis() ([]system.SysApi, error) {
	var err error
	apis := make([]system.SysApi, 0)
	s.db.Find(&apis)
	return apis, err
}

// 创建接口
func (s *MysqlService) CreateApi(req *request.CreateApiReq) (err error) {
	var api system.SysApi
	utils.Struct2StructByJson(req, &api)
	// 创建数据
	err = s.db.Create(&api).Error
	return
}

// 更新接口
func (s *MysqlService) UpdateApiById(id uint, req request.UpdateApiReq) (err error) {
	if id == req.ParentId {
		return errors.New("不能自关联")
	}
	var oldApi system.SysApi
	query := s.db.Table(oldApi.TableName()).Where("id = ?", id).First(&oldApi)
	if query.Error == gorm.ErrRecordNotFound {
		return errors.New("记录不存在")
	}

	// 比对增量字段
	var m system.SysApi
	utils.CompareDifferenceStructByJson(oldApi, req, &m)
	// 更新指定列
	err = query.Updates(m).Error
	return
}

// 批量删除接口
func (s *MysqlService) DeleteApiByIds(ids []uint) (err error) {
	var api system.SysApi
	// 先解除父级关联
	err = s.db.Table(api.TableName()).Where("parent_id IN (?)", ids).Update("parent_id", 0).Error
	if err != nil {
		return err
	}
	// 再删除
	err = s.db.Where("id IN (?)", ids).Delete(&api).Error
	if err != nil {
		return err
	}
	return
}

// 生成接口树
func GenApiTree(parent *response.ApiListResp, apis []system.SysApi) []response.ApiListResp {
	tree := make([]response.ApiListResp, 0)
	var resp []response.ApiListResp
	utils.Struct2StructByJson(apis, &resp)
	var parentId uint
	if parent != nil {
		parentId = parent.Id
	}
	for _, api := range resp {
		if api.ParentId == parentId {
			api.Children = GenApiTree(&api, apis)
			tree = append(tree, api)
		}
	}
	// 排序
	return tree
}