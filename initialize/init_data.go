package initialize

import (
	"anew-server/common"
	"anew-server/dto/service"
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
			Name:    "管理员",
			Keyword: "admin",
			Desc:    "系统管理员",
			Status:  &status,
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:    "访客",
			Keyword: "guest",
			Desc:    "外来访问人员",
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
			Name:      "Index",
			Title:     "首页",
			Icon:      "dashboard",
			Path:      "index",
			Component: "/Index/index",
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
			Title:      "系统设置根目录",
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
				roles[0],
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
			ParentId:  3,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[0],
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
			ParentId:  3,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[0],
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
			ParentId:  3,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[0],
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
			ParentId:  3,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[0],
			},
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Name:      "operlog",
			Title:     "操作日志",
			Icon:      "example",
			Path:      "log",
			Component: "/system/operlog",
			Sort:      4,
			Status:    &status,
			Visible:   &visible,
			ParentId:  3,
			Creator:   creator,
			Roles: []models.SysRole{
				roles[0],
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
	users := []models.SysUser{
		{
			Username: "admin",
			Password: utils.GenPwd("123456"),
			Mobile:   "18888888888",
			Avatar:   avatar,
			Name:     "管理员",
			RoleId:   1,
			Creator:  creator,
		},
		{
			Username: "guest",
			Password: utils.GenPwd("123456"),
			Mobile:   "15888888888",
			Avatar:   avatar,
			Name:     "来宾",
			RoleId:   2,
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

	// 4. 初始化接口
	apis := []models.SysApi{
		{
			Model: models.Model{
				Id: 1,
			},
			Method:   "POST",
			Path:     "/auth/login",
			Category: "auth",
			Desc:     "用户登录",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Method:   "POST",
			Path:     "/auth/logout",
			Category: "auth",
			Desc:     "用户登出",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 3,
			},
			Method:   "POST",
			Path:     "/auth/refresh_token",
			Category: "auth",
			Desc:     "刷新JWT令牌",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 4,
			},
			Method:   "POST",
			Path:     "/v1/user/info",
			Category: "user",
			Desc:     "获取当前登录用户信息",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 5,
			},
			Method:   "GET",
			Path:     "/v1/user/list",
			Category: "user",
			Desc:     "获取用户列表",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 6,
			},
			Method:   "PUT",
			Path:     "/v1/user/changePwd",
			Category: "user",
			Desc:     "修改用户登录密码",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 7,
			},
			Method:   "POST",
			Path:     "/v1/user/create",
			Category: "user",
			Desc:     "创建用户",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Method:   "PATCH",
			Path:     "/v1/user/update/:userId",
			Category: "user",
			Desc:     "更新用户",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 9,
			},
			Method:   "DELETE",
			Path:     "/v1/user/delete",
			Category: "user",
			Desc:     "批量删除用户",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 10,
			},
			Method:   "GET",
			Path:     "/v1/menu/tree",
			Category: "menu",
			Desc:     "获取权限菜单",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 11,
			},
			Method:   "GET",
			Path:     "/v1/menu/list",
			Category: "menu",
			Desc:     "获取菜单列表",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 12,
			},
			Method:   "POST",
			Path:     "/v1/menu/create",
			Category: "menu",
			Desc:     "创建菜单",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 13,
			},
			Method:   "PATCH",
			Path:     "/v1/menu/update/:menuId",
			Category: "menu",
			Desc:     "更新菜单",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 14,
			},
			Method:   "DELETE",
			Path:     "/v1/menu/delete",
			Category: "menu",
			Desc:     "批量删除菜单",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 15,
			},
			Method:   "GET",
			Path:     "/v1/role/list",
			Category: "role",
			Desc:     "获取角色列表",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 16,
			},
			Method:   "POST",
			Path:     "/v1/role/create",
			Category: "role",
			Desc:     "创建角色",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 17,
			},
			Method:   "PATCH",
			Path:     "/v1/role/update/:roleId",
			Category: "role",
			Desc:     "更新角色",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 18,
			},
			Method:   "DELETE",
			Path:     "/v1/role/delete",
			Category: "role",
			Desc:     "批量删除角色",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 19,
			},
			Method:   "GET",
			Path:     "/v1/api/list",
			Category: "api",
			Desc:     "获取接口列表",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 20,
			},
			Method:   "POST",
			Path:     "/v1/api/create",
			Category: "api",
			Desc:     "创建接口",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 21,
			},
			Method:   "PATCH",
			Path:     "/v1/api/update/:roleId",
			Category: "api",
			Desc:     "更新接口",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 22,
			},
			Method:   "DELETE",
			Path:     "/v1/api/delete",
			Category: "api",
			Desc:     "批量删除接口",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 23,
			},
			Method:   "GET",
			Path:     "/v1/menu/all/:roleId",
			Category: "menu",
			Desc:     "查询指定角色的菜单树",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 24,
			},
			Method:   "GET",
			Path:     "/v1/api/all/category/:roleId",
			Category: "api",
			Desc:     "查询指定角色的接口(以分类分组)",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 25,
			},
			Method:   "PATCH",
			Path:     "/v1/role/menus/update/:roleId",
			Category: "role",
			Desc:     "更新角色的权限菜单",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 26,
			},
			Method:   "PATCH",
			Path:     "/v1/role/apis/update/:roleId",
			Category: "role",
			Desc:     "更新角色的权限接口",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 27,
			},
			Method:   "GET",
			Path:     "/v1/operlog/list",
			Category: "operation-log",
			Desc:     "获取操作日志列表",
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 28,
			},
			Method:   "DELETE",
			Path:     "/v1/operlog/delete",
			Category: "operation-log",
			Desc:     "批量删除操作日志",
			Creator:  creator,
		},
	}
	for _, api := range apis {
		oldApi := models.SysApi{}
		notFound := common.Mysql.Where("id = ?", api.Id).First(&oldApi).RecordNotFound()
		if notFound {
			common.Mysql.Create(&api)
			// 创建服务
			s := service.New(nil)
			// 管理员拥有所有API权限role[1]
			s.CreateRoleCasbin(models.SysRoleCasbin{
				Keyword: roles[0].Keyword,
				Path:    api.Path,
				Method:  api.Method,
			})
			// 其他人暂时只有登录/获取用户信息的权限
			if api.Id < 5 || api.Id == 10 {
				s.CreateRoleCasbin(models.SysRoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}
		}
	}
}