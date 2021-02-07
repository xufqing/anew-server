package service

import (
	"anew-server/dto/request"
	"anew-server/models/asset"
	"anew-server/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func (s *MysqlService) GetHosts(req *request.HostReq) ([]asset.AssetHost, error) {
	var err error
	list := make([]asset.AssetHost, 0)
	query := s.db.Table(new(asset.AssetHost).TableName())
	group_id := strings.TrimSpace(req.GroupID)
	if group_id != "" {
		query = query.Raw("select * from tb_asset_host a where id in (select asset_host_id from relation_group_host where asset_group_id = ? ) and a.deleted_at is null", group_id)
	}
	host_name := strings.TrimSpace(req.HostName)
	if host_name != "" {
		query = query.Where("host_name LIKE ?", fmt.Sprintf("%%%s%%", host_name))
	}
	ip_address := strings.TrimSpace(req.IpAddress)
	if ip_address != "" {
		query = query.Where("ip_address LIKE ?", fmt.Sprintf("%%%s%%", ip_address))
	}
	os_version := strings.TrimSpace(req.OSVersion)
	if os_version != "" {
		query = query.Where("os_version LIKE ?", fmt.Sprintf("%%%s%%", os_version))
	}
	host_type := strings.TrimSpace(req.AuthType)
	if host_type != "" {
		query = query.Where("host_type LIKE ?", fmt.Sprintf("%%%s%%", host_type))
	}
	auth_type := strings.TrimSpace(req.AuthType)
	if auth_type != "" {
		query = query.Where("auth_type LIKE ?", fmt.Sprintf("%%%s%%", auth_type))
	}
	if group_id != "" {
		// 不使用分页
		err = query.Scan(&list).Error
		req.PageInfo.Total = int64((len(list)))

	} else {
		err = query.Find(&list).Count(&req.PageInfo.Total).Error
		if err == nil {
			if req.PageInfo.All {
				// 不使用分页
				err = query.Find(&list).Error
			} else {
				// 获取分页参数
				limit, offset := req.GetLimit()
				err = query.Limit(limit).Offset(offset).Find(&list).Error
			}
		}
	}

	return list, err
}

// 创建
func (s *MysqlService) CreateHost(req *request.CreateHostReq) (err error) {
	var host asset.AssetHost
	utils.Struct2StructByJson(req, &host)
	// 创建数据
	err = s.db.Create(&host).Error
	return
}

// 更新
func (s *MysqlService) UpdateHostById(id uint, req gin.H) (err error) {
	var oldHost asset.AssetHost
	query := s.db.Table(oldHost.TableName()).Where("id = ?", id).First(&oldHost)
	if query.Error == gorm.ErrRecordNotFound {
		return errors.New("记录不存在")
	}

	// 比对增量字段
	var m asset.AssetHost
	utils.CompareDifferenceStructByJson(oldHost, req, &m)
	// 更新指定列
	err = query.Updates(m).Error
	return
}

// 批量删除
func (s *MysqlService) DeleteHostByIds(ids []uint) (err error) {

	return s.db.Where("id IN (?)", ids).Delete(asset.AssetHost{}).Error
}

func (s *MysqlService) GetHostById(id uint) (asset.AssetHost, error) {
	var host asset.AssetHost
	err := s.db.Table(host.TableName()).Where("id = ?", id).First(&host).Error
	return host, err
}
