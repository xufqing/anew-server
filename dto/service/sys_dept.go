package service

import (
	"anew-server/dto/request"
	"anew-server/dto/response"
	"anew-server/models"
	"anew-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"sort"
	"fmt"
	"errors"
)

// 获取所有部门信息
func (s *MysqlService) GetDepts() []models.SysDept {
	//tree := make([]response.DeptTreeResp  , 0)
	depts := s.getAllDept()
	// 生成菜单树
	//tree = GenDeptTree(nil, depts)
	return depts
}

// 生成菜单树
func GenDeptTree(parent *response.DeptTreeResp, depts []models.SysDept) []response.DeptTreeResp {
	tree := make(response.DeptTreeResppList, 0)
	var resp []response.DeptTreeResp
	utils.Struct2StructByJson(depts, &resp)
	// parentId默认为0, 表示根菜单
	var parentId uint
	if parent != nil {
		parentId = parent.Id
	}
	for _, dept := range resp {
		// 增加key值
		dept.Key = fmt.Sprintf("%d",dept.Id)
		dept.Value = fmt.Sprintf("%d",dept.Id)
		dept.Title = dept.Name
		// 父菜单编号一致
		if dept.ParentId == parentId {
			// 递归获取子菜单
			dept.Children = GenDeptTree(&dept, depts)
			// 加入菜单树
			tree = append(tree, dept)
		}
	}
	// 排序
	sort.Sort(tree)
	return tree
}

// 获取部门信息，非树列表
func (s *MysqlService) getAllDept() []models.SysDept {
	depts := make([]models.SysDept, 0)
	s.tx.Order("sort").Find(&depts)
	return depts
}

// 创建部门
func (s *MysqlService) CreateDept(req *request.CreateDeptReq) (err error) {
	var dept models.SysDept
	utils.Struct2StructByJson(req, &dept)
	// 创建数据
	err = s.tx.Create(&dept).Error
	return
}

// 更新部门
func (s *MysqlService) UpdateDeptById(id uint, req gin.H) (err error) {
	var oldDept models.SysDept
	query := s.tx.Table(oldDept.TableName()).Where("id = ?", id).First(&oldDept)
	if query.RecordNotFound() {
		return errors.New("记录不存在")
	}

	// 比对增量字段
	m := make(gin.H, 0)
	utils.CompareDifferenceStructByJson(oldDept, req, &m)

	// 更新指定列
	err = query.Updates(m).Error
	return
}

// 批量删除菜单
func (s *MysqlService) DeleteDeptByIds(ids []uint) (err error) {
	var dept models.SysDept
	// 先解除父级关联
	err = s.tx.Table(dept.TableName()).Where("parent_id IN (?)", ids).Update("parent_id",0).Error
	if err != nil{
		return err
	}
	// 再删除
	err = s.tx.Where("id IN (?)", ids).Delete(models.SysDept{}).Error
	if err != nil{
		return err
	}
	return
}
