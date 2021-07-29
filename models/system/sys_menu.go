package system

import "anew-server/models"

// 系统菜单表
type SysMenu struct {
	models.Model
	Name     string    `gorm:"comment:'菜单名称';size:64" json:"name"`
	Icon     string    `gorm:"comment:'菜单图标';size:64" json:"icon"`
	Path     string    `gorm:"comment:'菜单访问路径';size:64" json:"path"`
	Sort     int       `gorm:"type:int(3);comment:'菜单顺序(同级菜单, 从0开始, 越小显示越靠前)'" json:"sort"`
	ParentId uint      `gorm:"default:0;comment:'父菜单编号(编号为0时表示根菜单)'" json:"parent_id"`
	Creator  string    `gorm:"comment:'创建人';size:64" json:"creator"`
	Children []SysMenu `gorm:"-" json:"children"`                          // 子菜单集合
	Roles    []SysRole `gorm:"many2many:relation_role_menu;" json:"roles"` // 角色菜单多对多关系
}

func (m SysMenu) TableName() string {
	return m.Model.TableName("sys_menu")
}

// 获取选中列表
func GetCheckedMenuIds(list []uint, allMenu []SysMenu) []uint {
	checked := make([]uint, 0)
	for _, c := range list {
		// 获取子流水线
		parent := SysMenu{
			ParentId: c,
		}
		children := parent.GetChildrenIds(allMenu)
		// 判断子流水线是否全部在create中
		count := 0
		for _, child := range children {
			// 避免环包调用, 不再调用utils
			// if utils.ContainsUint(list, child) {
			// 	count++
			// }
			contains := false
			for _, v := range list {
				if v == child {
					contains = true
				}
			}
			if contains {
				count++
			}
		}
		if len(children) == count {
			// 全部选中
			checked = append(checked, c)
		}
	}
	return checked
}

// 查找子菜单编号
func (m SysMenu) GetChildrenIds(allMenu []SysMenu) []uint {
	childrenIds := make([]uint, 0)
	for _, menu := range allMenu {
		if menu.ParentId == m.ParentId {
			childrenIds = append(childrenIds, menu.Id)
		}
	}
	return childrenIds
}
