package service

import (
	"anew-server/dto/request"
	"anew-server/models"
	"anew-server/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)


func (s *MysqlService) GetApis(req *request.ApiListReq) ([]models.SysApi, error) {
	var err error
	list := make([]models.SysApi, 0)
	query := s.tx.Table(new(models.SysApi).TableName())
	method := strings.TrimSpace(req.Method)
	if method != "" {
		query = query.Where("method LIKE ?", fmt.Sprintf("%%%s%%", method))
	}
	path := strings.TrimSpace(req.Path)
	if path != "" {
		query = query.Where("path LIKE ?", fmt.Sprintf("%%%s%%", path))
	}
	category := strings.TrimSpace(req.Category)
	if category != "" {
		query = query.Where("category LIKE ?", fmt.Sprintf("%%%s%%", category))
	}
	creator := strings.TrimSpace(req.Creator)
	if creator != "" {
		query = query.Where("creator LIKE ?", fmt.Sprintf("%%%s%%", creator))
	}
	permission := strings.TrimSpace(req.Permission)
	if permission != "" {
		query = query.Where("permission LIKE ?", fmt.Sprintf("%%%s%%", permission))
	}

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

// 创建接口
func (s *MysqlService) CreateApi(req *request.CreateApiReq) (err error) {
	var api models.SysApi
	utils.Struct2StructByJson(req, &api)
	// 创建数据
	err = s.tx.Create(&api).Error
	return
}

// 更新接口
func (s *MysqlService) UpdateApiById(id uint, req gin.H) (err error) {
	var oldApi models.SysApi
	query := s.tx.Table(oldApi.TableName()).Where("id = ?", id).First(&oldApi)
	if query.RecordNotFound() {
		return errors.New("记录不存在")
	}

	// 比对增量字段
	m := make(gin.H, 0)
	utils.CompareDifferenceStructByJson(oldApi, req, &m)
	// 更新指定列
	err = query.Updates(m).Error
	return
}

// 批量删除接口
func (s *MysqlService) DeleteApiByIds(ids []uint) (err error) {

	return s.tx.Where("id IN (?)", ids).Delete(models.SysApi{}).Error
}