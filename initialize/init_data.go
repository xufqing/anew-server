package initialize

import (
	"anew-server/common"
	"anew-server/models"
	"anew-server/utils"
)

// 初始化数据
func InitData() {
	// 1. 初始化角色
	creator := "系统自动创建"
	status := true
	visible := true
	roles := []models.SysRole{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:    "访客",
			Keyword: "guest",
			Desc:    "外来访问人员",
			Status:  &status,
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:    "管理员",
			Keyword: "admin",
			Desc:    "系统管理员",
			Status:  &status,
			Creator: creator,
		},
	}
	for _, role := range roles {
		oldRole := models.SysRole{}
		notFound := common.Mysql.Where("id = ?", role.Id).First(&oldRole).RecordNotFound()
		if notFound {
			common.Mysql.Create(&role)
		}
	}

	// 2. 初始化菜单
	noBreadcrumb := false
	menus := []models.SysMenu{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:       "Root", // 对于想让子菜单显示在上层不显示的父级菜单不设置名字
			Title:      "根",
			Icon:       "dashboard",
			Path:       "/",
			Component:  "", // 如果包含子菜单, Component为空
			Sort:       0,
			Status:     &status,
			Visible:    &visible,
			Breadcrumb: &noBreadcrumb, // 面包屑不可见
			ParentId:   0,
			Creator:    creator,
			Roles:      roles,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:      "index",
			Title:     "首页",
			Icon:      "dashboard",
			Path:      "index",
			Component: "/index/index",
			Sort:      0,
			Status:    &status,
			Visible:   &visible,
			ParentId:  1,
			Creator:   creator,
			Roles:     roles,
		},
		{
			Model: models.Model{
				Id: 3,
			},
			Name:       "systemRoot",
			Title:      "系统设置",
			Icon:       "component",
			Path:       "/system",
			Component:  "",
			Sort:       1,
			Status:     &status,
			Visible:    &visible,
			Breadcrumb: &noBreadcrumb, // 面包屑不可见
			ParentId:   0,
			Creator:    creator,
			Roles: []models.SysRole{
				roles[1],
			},
		},
		{
			Model: models.Model{
				Id: 4,
			},
			Name:      "menu",
			Title:     "菜单管理",
			Icon:      "tree-table",
			Path:      "menu", // 子菜单不用全路径, 自动继承
			Component: "/system/menu",
			Sort:      0,
			Status:    &status,
			Visible:   &visible,
			ParentId:  2,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[1],
			},
		},
		{
			Model: models.Model{
				Id: 5,
			},
			Name:      "role",
			Title:     "角色管理",
			Icon:      "peoples",
			Path:      "role",
			Component: "/system/role",
			Sort:      1,
			Status:    &status,
			Visible:   &visible,
			ParentId:  2,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[1],
			},
		},
		{
			Model: models.Model{
				Id: 6,
			},
			Name:      "user",
			Title:     "用户管理",
			Icon:      "user",
			Path:      "user",
			Component: "/system/user",
			Sort:      2,
			Status:    &status,
			Visible:   &visible,
			ParentId:  2,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[1],
			},
		},
		{
			Model: models.Model{
				Id: 7,
			},
			Name:      "api",
			Title:     "接口管理",
			Icon:      "tree",
			Path:      "api",
			Component: "/system/api",
			Sort:      3,
			Status:    &status,
			Visible:   &visible,
			ParentId:  2,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[1],
			},
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Name:      "operation-log",
			Title:     "操作日志",
			Icon:      "example",
			Path:      "optlog",
			Component: "/system/optlog",
			Sort:      5,
			Status:    &status,
			Visible:   &visible,
			ParentId:  2,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[1],
			},
		},
	}
	for _, menu := range menus {
		oldMenu := models.SysMenu{}
		notFound := common.Mysql.Where("id = ?", menu.Id).First(&oldMenu).RecordNotFound()
		if notFound {
			common.Mysql.Create(&menu)
		}
	}

	// 3. 初始化用户
	// 默认头像
	avatar := "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	users := [2]models.SysUser{
		{
			Username: "admin",
			Password: utils.GenPwd("123456"),
			Mobile:   "18888888888",
			Avatar:   avatar,
			Name:     "管理员",
			RoleId:   2,
			Creator:  creator,
		},
		{
			Username: "guest",
			Password: utils.GenPwd("123456"),
			Mobile:   "15888888888",
			Avatar:   avatar,
			Name:     "来宾",
			RoleId:   1,
			Creator:  creator,
		},
	}
	for _, user := range users {
		oldUser := models.SysUser{}
		notFound := common.Mysql.Where("username = ?", user.Username).First(&oldUser).RecordNotFound()
		if notFound {
			common.Mysql.Create(&user)
		}
	}
}
