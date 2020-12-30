package service

import (
	"anew-server/dto/request"
	"anew-server/models/asset"
	"anew-server/pkg/common"
	"fmt"
	"strings"
)

// 获取操作日志
func (s *MysqlService) GetSshRecords(req *request.SshRecordListReq) ([]asset.SshRecord, error) {
	var err error
	list := make([]asset.SshRecord, 0)
	query := common.Mysql
	key := strings.TrimSpace(req.Key)
	if key != "" {
		query = query.Where("key LIKE ?", fmt.Sprintf("%%%s%%", key))
	}
	username := strings.TrimSpace(req.UserName)
	if username != "" {
		query = query.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	hostname := strings.TrimSpace(req.HostName)
	if hostname != "" {
		query = query.Where("host_name LIKE ?", fmt.Sprintf("%%%s%%", hostname))
	}
	ip_address := strings.TrimSpace(req.IpAddress)
	if ip_address != "" {
		query = query.Where("ip_address LIKE ?", fmt.Sprintf("%%%s%%", ip_address))
	}
	query = query.Order("id DESC")
	// 查询条数
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
	return list, err
}

// 批量删除记录
func (s *MysqlService) DeleteSshRecordByIds(ids []uint) (err error) {
	return s.db.Where("id IN (?)", ids).Delete(asset.SshRecord{}).Error
}