package dao

import (
	"anew-server/api/request"
	"anew-server/api/response"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sort"
	"strings"
)

// 获取所有字典信息
func (s *MysqlService) GetDicts(req *request.DictListReq) []system.SysDict {
	Dicts := make([]system.SysDict, 0)
	db := common.Mysql
	typeKey := strings.TrimSpace(req.TypeKey)
	if typeKey != "" {
		var dist system.SysDict
		db = db.Order("sort").Preload("Dicts").Where("`dict_key` LIKE ?", fmt.Sprintf("%%%s%%", typeKey)).First(&dist)
		return dist.Dicts
	}
	key := strings.TrimSpace(req.DictKey)
	if key != "" {
		db = db.Where("dict_key LIKE ?", fmt.Sprintf("%%%s%%", key))
	}
	value := strings.TrimSpace(req.DictValue)
	if value != "" {
		db = db.Where("dict_value LIKE ?", fmt.Sprintf("%%%s%%", value))
	}
	creator := strings.TrimSpace(req.Creator)
	if creator != "" {
		db = db.Where("creator LIKE ?", fmt.Sprintf("%%%s%%", creator))
	}
	db.Order("sort").Find(&Dicts)
	return Dicts
}

// 生成字典树
func GenDictTree(parent *response.DictTreeResp, Dicts []system.SysDict) []response.DictTreeResp {
	tree := make(response.DictTreeRespList, 0)
	var resp []response.DictTreeResp
	utils.Struct2StructByJson(Dicts, &resp)
	// parentId默认为0, 表示根菜单
	var parentId uint
	if parent != nil {
		parentId = parent.Id
	}
	for _, Dict := range resp {
		// 父菜单编号一致
		if Dict.ParentId == parentId {
			// 递归获取子菜单
			Dict.Children = GenDictTree(&Dict, Dicts)
			// 加入菜单树
			tree = append(tree, Dict)
		}
	}
	sort.Sort(tree)
	return tree
}

// 字典MAP
func GenDictMap(parent *response.DictTreeResp, Dicts []system.SysDict) map[string]response.DictTreeRespList {
	dictMap := map[string]response.DictTreeRespList{}
	var resp []response.DictTreeResp
	utils.Struct2StructByJson(Dicts, &resp)
	var parentId uint
	if parent != nil {
		parentId = parent.Id
	}
	for _, Dict := range resp {
		if Dict.ParentId == parentId {
			dictMap[Dict.DictKey] = GenDictTree(&Dict, Dicts)

		}
	}
	return dictMap
}

// 创建字典
func (s *MysqlService) CreateDict(req *request.CreateDictReq) (err error) {
	var Dict system.SysDict
	utils.Struct2StructByJson(req, &Dict)
	// 创建数据
	err = s.db.Create(&Dict).Error
	return
}

// 更新字典
func (s *MysqlService) UpdateDictById(id uint, req request.UpdateDictReq) (err error) {
	if id == req.ParentId {
		return errors.New("不能自关联")
	}
	var oldDict system.SysDict
	query := s.db.Table(oldDict.TableName()).Where("id = ?", id).First(&oldDict)
	if query.Error == gorm.ErrRecordNotFound {
		return errors.New("记录不存在")
	}
	// 比对增量字段,使用map确保gorm可更新零值
	var m map[string]interface{}
	utils.CompareDifferenceStructByJson(oldDict, req, &m)
	// 更新指定列
	err = query.Updates(m).Error
	return
}

// 批量删除字典
func (s *MysqlService) DeleteDictByIds(ids []uint) (err error) {
	var Dict system.SysDict
	// 先解除父级关联
	err = s.db.Table(Dict.TableName()).Where("parent_id IN (?)", ids).Update("parent_id", 0).Error
	if err != nil {
		return err
	}
	// 再删除
	err = s.db.Where("id IN (?)", ids).Delete(&Dict).Error
	if err != nil {
		return err
	}
	return
}
