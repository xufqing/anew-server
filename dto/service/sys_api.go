package service

import (
	"anew-server/common"
	"anew-server/dto/request"
	"anew-server/models"
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

	if req.PageInfo.All {
		// 不使用分页
		err = query.Find(&list).Error
	} else {
		// 获取分页参数
		limit, offset := req.GetLimit()
		err = query.Limit(limit).Offset(offset).Find(&list).Error
	}

	return list, err
}

// 创建接口
func (s *MysqlService) CreateApi(req *request.CreateApiReq) (err error) {


	return
}

// 更新接口
func (s *MysqlService) UpdateApiById(id uint, req gin.H) (err error) {

	return
}

// 批量删除接口
func (s *MysqlService) DeleteApiByIds(ids []uint) (err error) {
	return
}


// 获取全部API
func (s *MysqlService) GetAllApi() []models.SysApi {
	apis := make([]models.SysApi, 0)
	// 查询所有菜单
	err := s.tx.Find(&apis).Error
	common.Log.Warn("[getAllApi]", err)
	return apis
}