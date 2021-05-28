package service

import (
	"anew-server/dao/request"
	"anew-server/models/asset"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// 获取所有组
func (s *MysqlService) GetAssetGroups(req *request.AssetGroupReq) ([]asset.AssetGroup, error) {
	var err error
	list := make([]asset.AssetGroup, 0)
	db := common.Mysql
	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	creator := strings.TrimSpace(req.Creator)
	if creator != "" {
		db = db.Where("creator LIKE ?", fmt.Sprintf("%%%s%%", creator))
	}

	// 查询条数
	err = db.Find(&list).Count(&req.PageInfo.Total).Error
	if err == nil {
		if req.PageInfo.All {
			// 不使用分页
			err = db.Preload("Hosts").Find(&list).Error
		} else {
			// 获取分页参数
			limit, offset := req.GetLimit()
			//err = db.Preload("Hosts",
			//	func(db *gorm.DB) *gorm.DB { return db.Select("id") }).
			//	Limit(limit).Offset(offset).Find(&list).Error
			err = db.Preload("Hosts").Limit(limit).Offset(offset).Find(&list).Error
		}
	}
	if req.NotNull {
		// 过滤出非空的主机组
		newList := make([]asset.AssetGroup, 0)
		for _, i := range list {
			if len(i.Hosts) > 0 {
				newList = append(newList, i)
			}
		}
		return newList, err
	}
	return list, err
}

// 创建分组
func (s *MysqlService) CreateAssetGroup(req *request.CreateAssetGroupReq) (err error) {
	var group asset.AssetGroup
	hostsId := req.Hosts
	req.Hosts = nil
	utils.Struct2StructByJson(req, &group)
	// 创建关联
	if hostsId != nil {
		var hosts []asset.AssetHost
		err = s.db.Where("id in (?)", hostsId).Find(&hosts).Error
		if err != nil {
			return
		}
		group.Hosts = hosts
	}
	// 创建数据
	err = s.db.Create(&group).Error
	return
}

// 更新分组
func (s *MysqlService) UpdateAssetGroupById(id uint, req request.UpdateAssetGroupReq) (err error) {
	var oldGroup asset.AssetGroup
	query := s.db.Model(oldGroup).Where("id = ?", id).First(&oldGroup)
	if query.Error == gorm.ErrRecordNotFound {
		return errors.New("记录不存在")
	}
	// 比对增量字段,使用map确保gorm可更新零值
	var m map[string]interface{}
	utils.CompareDifferenceStructByJson(oldGroup, req, &m)
	delete(m, "hosts")
	// 更新指定列
	err = query.Updates(m).Error
	// 更新关联
	if req.Hosts != nil {
		var hosts []asset.AssetHost
		err = s.db.Where("id in (?)", req.Hosts).Find(&hosts).Error
		if err != nil {
			return
		}
		var group asset.AssetGroup
		err = s.db.Where("id = ?", id).First(&group).Error
		err = s.db.Model(&group).Association("Hosts").Replace(&hosts)
	}
	return
}

// 批量删除分组
func (s *MysqlService) DeleteAssetGroupByIds(ids []uint) (err error) {
	var groups []asset.AssetGroup
	// 查询符合条件的分组, 以及关联的主机
	err = s.db.Preload("Hosts").Where("id IN (?)", ids).Find(&groups).Error
	if err != nil {
		return
	}
	newIds := make([]uint, 0)
	for _, g := range groups {
		if len(g.Hosts) > 0 {
			return errors.New(fmt.Sprintf("分组[%s]仍有%d台关联主机, 请先移除关联主机再删除分组", g.Name, len(g.Hosts)))
		}
		newIds = append(newIds, g.Id)
	}
	if len(newIds) > 0 {
		// 执行删除
		err = s.db.Where("id IN (?)", newIds).Delete(asset.AssetGroup{}).Error
	}
	return
}
