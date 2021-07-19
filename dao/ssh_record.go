package dao

import (
	"anew-server/api/request"
	"anew-server/models/asset"
	"anew-server/pkg/common"
	"fmt"
	"strings"
)

// 获取操作日志
func (s *MysqlService) GetSSHRecords(host_id uint, req *request.SSHRecordReq) ([]asset.SSHRecord, error) {
	var err error
	list := make([]asset.SSHRecord, 0)
	query := common.Mysql
	username := strings.TrimSpace(req.UserName)
	if username != "" {
		query = query.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", username))
	}
	query = query.Order("id DESC")
	// 查询条数
	err = query.Where("host_id = ?", host_id).Find(&list).Count(&req.PageInfo.Total).Error
	if err == nil {
		if req.PageInfo.All {
			// 不使用分页
			err = query.Where("host_id = ?", host_id).Find(&list).Error
		} else {
			// 获取分页参数
			limit, offset := req.GetLimit()
			err = query.Limit(limit).Offset(offset).Find(&list).Error
		}
	}
	return list, err
}

func (s *MysqlService) GetSSHRecordByConnectID(connect_id string) (asset.SSHRecord, error) {
	var record asset.SSHRecord
	err := s.db.Where("connect_id = ?", connect_id).First(&record).Error
	return record, err

}

// 批量删除记录
func (s *MysqlService) DeleteSSHRecordByIds(ids []uint) (err error) {
	return s.db.Where("id IN (?)", ids).Delete(asset.SSHRecord{}).Error
}
