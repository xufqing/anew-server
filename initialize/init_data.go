package initialize

import (
	dao "anew-server/dao"
	"anew-server/models"
	"anew-server/models/system"
	"anew-server/pkg/common"
	"anew-server/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

// 初始化数据
func InitData() {
	// 1. 初始化角色
	creator := "系统创建"
	roles := []system.SysRole{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:    "管理员",
			Keyword: "admin",
			Desc:    "系统管理员",
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:    "访客",
			Keyword: "guest",
			Desc:    "外来访问人员",
			Creator: creator,
		},
	}
	for _, role := range roles {
		oldRole := system.SysRole{}
		err := common.Mysql.Where("id = ?", role.Id).First(&oldRole).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&role)
		}
	}

	// 2. 初始化菜单
	menus := []system.SysMenu{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:     "工作台",
			Icon:     "icon-yibiaopan",
			Path:     "workbench",
			Sort:     0,
			ParentId: 0,
			Creator:  creator,
			Roles:    roles,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:     "仪表盘",
			Icon:     "icon-yibiaopan",
			Path:     "dashboard",
			Sort:     0,
			ParentId: 1,
			Creator:  creator,
			Roles:    roles,
		},
		{
			Model: models.Model{
				Id: 3,
			},
			Name:     "我的项目",
			Icon:     "icon-yibiaopan",
			Path:     "myprojet",
			Sort:     0,
			ParentId: 1,
			Creator:  creator,
			Roles:    roles,
		},
		{
			Model: models.Model{
				Id: 12,
			},
			Name:     "资产管理",
			Icon:     "icon-zichan",
			Path:     "asset",
			Sort:     1,
			ParentId: 0,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 13,
			},
			Name:     "主机管理",
			Path:     "host",
			Sort:     1,
			ParentId: 12,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 14,
			},
			Name:     "主机分组",
			Path:     "group",
			Sort:     2,
			ParentId: 12,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 15,
			},
			Name:     "会话管理",
			Path:     "session",
			Sort:     3,
			ParentId: 12,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 4,
			},
			Name:     "系统设置",
			Icon:     "icon-shezhi",
			Path:     "system",
			Sort:     999,
			ParentId: 0,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 5,
			},
			Name:     "用户管理",
			Path:     "user",
			Sort:     10,
			ParentId: 4,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 6,
			},
			Name:     "部门管理",
			Path:     "dept",
			Sort:     11,
			ParentId: 4,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 7,
			},
			Name:     "菜单管理",
			Path:     "menu",
			Sort:     12,
			ParentId: 4,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Name:     "角色管理",
			Path:     "role",
			Sort:     13,
			ParentId: 4,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 9,
			},
			Name:     "接口管理",
			Path:     "api",
			Sort:     14,
			ParentId: 4,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 10,
			},
			Name:     "字典管理",
			Path:     "dict",
			Sort:     15,
			ParentId: 4,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 11,
			},
			Name:     "操作日志",
			Path:     "operlog",
			Sort:     16,
			ParentId: 4,
			Creator:  creator,
			Roles: []system.SysRole{
				roles[0],
			},
		},
	}
	for _, menu := range menus {
		oldMenu := system.SysMenu{}
		err := common.Mysql.Where("id = ?", menu.Id).First(&oldMenu).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&menu)
		}
	}

	// 3. 初始化用户
	// 默认头像
	avatar := "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	users := []system.SysUser{
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
		oldUser := system.SysUser{}
		err := common.Mysql.Where("username = ?", user.Username).First(&oldUser).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&user)
		}
	}
	// 初始化字典
	dicts := []system.SysDict{
		{
			Model: models.Model{
				Id: 1,
			},
			DictKey:     "env_type",
			DictValue:   "应用环境",
			Desc:    "应用环境",
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			DictKey:      "dev",
			DictValue:    "开发环境",
			ParentId: 1,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 3,
			},
			DictKey:      "test",
			DictValue:    "测试环境",
			ParentId: 1,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 4,
			},
			DictKey:      "prod",
			DictValue:    "生产环境",
			ParentId: 1,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 5,
			},
			DictKey:     "auth_type",
			DictValue:   "认证类型",
			Desc:    "主机认证类型",
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 6,
			},
			DictKey:      "key",
			DictValue:    "秘钥验证",
			ParentId: 5,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 7,
			},
			DictKey:      "password",
			DictValue:    "密码验证",
			ParentId: 5,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 8,
			},
			DictKey:     "host_type",
			DictValue:   "主机类型",
			Desc:    "主机分类",
			Creator: creator,
		},
		{
			Model: models.Model{
				Id: 9,
			},
			DictKey:      "vm",
			DictValue:    "虚拟机",
			ParentId: 8,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 10,
			},
			DictKey:      "phy",
			DictValue:    "物理机",
			ParentId: 8,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 11,
			},
			DictKey:      "aliyun",
			DictValue:    "阿里云",
			ParentId: 8,
			Creator:  creator,
		},
	}

	for _, dict := range dicts {
		oldDict := system.SysDict{}
		err := common.Mysql.Where("id = ?", dict.Id).First(&oldDict).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&dict)
		}
	}

	// 5. 初始化接口
	apis := []system.SysApi{
		{
			Model: models.Model{
				Id: 1,
			},
			Name:     "基本权限",
			Method:   "",
			Path:     "",
			PermsTag: "base:all",
			Desc:     "基本权限",
			ParentId: 0,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 2,
			},
			Name:     "用户登录",
			Method:   "POST",
			Path:     "/auth/login",
			PermsTag: "base:login",
			Desc:     "获取用户信息和token",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 3,
			},
			Name:     "用户登出",
			Method:   "POST",
			Path:     "/auth/logout",
			PermsTag: "base:logout",
			Desc:     "用户登出",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 4,
			},
			Name:     "刷新令牌",
			Method:   "POST",
			Path:     "/auth/refresh_token",
			PermsTag: "base:refresh:token",
			Desc:     "刷新JWT令牌",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 5,
			},
			Name:     "用户信息",
			Method:   "GET",
			Path:     "/v1/user/info",
			PermsTag: "base:info",
			Desc:     "用户信息",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 6,
			},
			Name:     "更新基本信息",
			Method:   "PATCH",
			Path:     "/v1/user/info/update/:userId",
			PermsTag: "base:update:info",
			Desc:     "更新信息",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 7,
			},
			Name:     "上传头像",
			Method:   "POST",
			Path:     "/v1/user/info/uploadImg",
			PermsTag: "base:upload:avatar",
			Desc:     "上传头像",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 8,
			},
			Name:     "修改密码",
			Method:   "PUT",
			Path:     "/v1/user/changePwd",
			PermsTag: "base:change:pwd",
			Desc:     "修改密码",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 9,
			},
			Name:     "当前菜单",
			Method:   "GET",
			Path:     "/v1/menu/tree",
			PermsTag: "base:menu:tree",
			Desc:     "当前菜单",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 10,
			},
			Name:     "用户管理",
			Method:   "",
			Path:     "",
			PermsTag: "user:all",
			Desc:     "用户管理",
			ParentId: 0,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 11,
			},
			Name:     "用户列表",
			Method:   "GET",
			Path:     "/v1/user/list",
			PermsTag: "user:list",
			Desc:     "用户列表",
			Creator:  creator,
			ParentId: 10,
		},
		{
			Model: models.Model{
				Id: 12,
			},
			Name:     "创建用户",
			Method:   "POST",
			Path:     "/v1/user/create",
			PermsTag: "user:create",
			Desc:     "创建用户",
			Creator:  creator,
			ParentId: 10,
		},
		{
			Model: models.Model{
				Id: 13,
			},
			Name:     "更新用户",
			Method:   "PATCH",
			Path:     "/v1/user/update/:userId",
			PermsTag: "user:update",
			Desc:     "更新用户",
			Creator:  creator,
			ParentId: 10,
		},
		{
			Model: models.Model{
				Id: 14,
			},
			Name:     "删除用户",
			Method:   "DELETE",
			Path:     "/v1/user/delete",
			PermsTag: "user:delete",
			Desc:     "删除用户",
			Creator:  creator,
			ParentId: 10,
		},
		{
			Model: models.Model{
				Id: 15,
			},
			Name:     "菜单管理",
			Method:   "",
			Path:     "",
			PermsTag: "menu:all",
			Desc:     "菜单管理",
			ParentId: 0,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 16,
			},
			Name:     "菜单列表",
			Method:   "GET",
			Path:     "/v1/menu/list",
			PermsTag: "menu:list",
			Desc:     "菜单列表",
			Creator:  creator,
			ParentId: 15,
		},
		{
			Model: models.Model{
				Id: 17,
			},
			Name:     "创建菜单",
			Method:   "POST",
			Path:     "/v1/menu/create",
			PermsTag: "menu:create",
			Desc:     "创建菜单",
			Creator:  creator,
			ParentId: 15,
		},
		{
			Model: models.Model{
				Id: 18,
			},
			Name:     "更新菜单",
			Method:   "PATCH",
			Path:     "/v1/menu/update/:menuId",
			PermsTag: "menu:update",
			Desc:     "更新菜单",
			Creator:  creator,
			ParentId: 15,
		},
		{
			Model: models.Model{
				Id: 19,
			},
			Name:     "删除菜单",
			Method:   "DELETE",
			Path:     "/v1/menu/delete",
			PermsTag: "menu:delete",
			Desc:     "删除菜单",
			Creator:  creator,
			ParentId: 15,
		},
		{
			Model: models.Model{
				Id: 20,
			},
			Name:     "角色管理",
			Method:   "",
			Path:     "",
			PermsTag: "role:all",
			Desc:     "角色管理",
			ParentId: 0,
			Creator:  creator,
		},
		{
			Model: models.Model{
				Id: 21,
			},
			Name:     "角色列表",
			Method:   "GET",
			Path:     "/v1/role/list",
			PermsTag: "role:list",
			Desc:     "角色列表",
			Creator:  creator,
			ParentId: 20,
		},
		{
			Model: models.Model{
				Id: 22,
			},
			Name:     "创建角色",
			Method:   "POST",
			Path:     "/v1/role/create",
			PermsTag: "role:create",
			Desc:     "创建角色",
			Creator:  creator,
			ParentId: 20,
		},
		{
			Model: models.Model{
				Id: 23,
			},
			Name:     "更新角色",
			Method:   "PATCH",
			Path:     "/v1/role/update/:roleId",
			PermsTag: "role:update",
			Desc:     "更新角色",
			Creator:  creator,
			ParentId: 20,
		},
		{
			Model: models.Model{
				Id: 24,
			},
			Name:     "删除角色",
			Method:   "DELETE",
			Path:     "/v1/role/delete",
			PermsTag: "role:delete",
			Desc:     "删除角色",
			Creator:  creator,
			ParentId: 20,
		},
		{
			Model: models.Model{
				Id: 25,
			},
			Name:     "修改权限",
			Method:   "PATCH",
			Path:     "/v1/role/perms/update/:roleId",
			PermsTag: "role:update:perms",
			Desc:     "更新权限",
			Creator:  creator,
			ParentId: 20,
		},
		{
			Model: models.Model{
				Id: 26,
			},
			Name:     "获取权限",
			Method:   "GET",
			Path:     "/v1/role/perms/:roleId",
			PermsTag: "role:list:perms",
			Desc:     "获取权限",
			Creator:  creator,
			ParentId: 20,
		},
		{
			Model: models.Model{
				Id: 27,
			},
			Name:     "接口管理",
			Method:   "",
			Path:     "",
			PermsTag: "api:all",
			Desc:     "接口管理",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 28,
			},
			Name:     "查询接口",
			Method:   "GET",
			Path:     "/v1/api/list",
			PermsTag: "api:list",
			Desc:     "获取接口",
			Creator:  creator,
			ParentId: 27,
		},
		{
			Model: models.Model{
				Id: 29,
			},
			Name:     "创建接口",
			Method:   "POST",
			Path:     "/v1/api/create",
			PermsTag: "api:create",
			Desc:     "创建接口",
			Creator:  creator,
			ParentId: 27,
		},
		{
			Model: models.Model{
				Id: 30,
			},
			Name:     "更新接口",
			Method:   "PATCH",
			Path:     "/v1/api/update/:apiId",
			PermsTag: "api:update",
			Desc:     "更新接口",
			Creator:  creator,
			ParentId: 27,
		},
		{
			Model: models.Model{
				Id: 31,
			},
			Name:     "删除接口",
			Method:   "DELETE",
			Path:     "/v1/api/delete",
			PermsTag: "api:delete",
			Desc:     "删除接口",
			Creator:  creator,
			ParentId: 27,
		},
		{
			Model: models.Model{
				Id: 32,
			},
			Name:     "日志管理",
			Method:   "",
			Path:     "",
			PermsTag: "operlog:all",
			Desc:     "日志管理",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 33,
			},
			Name:     "日志列表",
			Method:   "GET",
			Path:     "/v1/operlog/list",
			PermsTag: "operlog:list",
			Desc:     "日志列表",
			Creator:  creator,
			ParentId: 32,
		},
		{
			Model: models.Model{
				Id: 34,
			},
			Name:     "删除日志",
			Method:   "DELETE",
			Path:     "/v1/operlog/delete",
			PermsTag: "operlog:delete",
			Desc:     "删除日志",
			Creator:  creator,
			ParentId: 32,
		},
		{
			Model: models.Model{
				Id: 35,
			},
			Name:     "字典管理",
			Method:   "",
			Path:     "",
			PermsTag: "dict:all",
			Desc:     "字典管理",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 36,
			},
			Name:     "字典列表",
			Method:   "GET",
			Path:     "/v1/dict/list",
			PermsTag: "dict:list",
			Desc:     "字典列表",
			Creator:  creator,
			ParentId: 1,
		},
		{
			Model: models.Model{
				Id: 37,
			},
			Name:     "创建字典",
			Method:   "POST",
			Path:     "/v1/dict/create",
			PermsTag: "dict:create",
			Desc:     "创建字典",
			Creator:  creator,
			ParentId: 35,
		},
		{
			Model: models.Model{
				Id: 38,
			},
			Name:     "更新字典",
			Method:   "PATCH",
			Path:     "/v1/dict/update/:dictId",
			PermsTag: "dict:update",
			Desc:     "更新字典",
			Creator:  creator,
			ParentId: 35,
		},
		{
			Model: models.Model{
				Id: 39,
			},
			Name:     "删除字典",
			Method:   "DELETE",
			Path:     "/v1/dict/delete",
			PermsTag: "dict:delete",
			Desc:     "删除字典",
			Creator:  creator,
			ParentId: 35,
		},
		{
			Model: models.Model{
				Id: 40,
			},
			Name:     "主机管理",
			Method:   "",
			Path:     "",
			PermsTag: "host:all",
			Desc:     "主机管理",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 41,
			},
			Name:     "查询主机",
			Method:   "GET",
			Path:     "/v1/host/list",
			PermsTag: "host:list",
			Desc:     "主机列表",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 42,
			},
			Name:     "创建主机",
			Method:   "POST",
			Path:     "/v1/host/create",
			PermsTag: "host:create",
			Desc:     "创建主机",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 43,
			},
			Name:     "更新主机",
			Method:   "PATCH",
			Path:     "/v1/host/update/:hostId",
			PermsTag: "host:update",
			Desc:     "更新主机",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 44,
			},
			Name:     "删除主机",
			Method:   "DELETE",
			Path:     "/v1/host/delete",
			PermsTag: "host:delete",
			Desc:     "删除主机",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 45,
			},
			Name:     "连接SSH",
			Method:   "GET",
			Path:     "/v1/host/SSH",
			PermsTag: "host:ssh:",
			Desc:     "连接SSH",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 46,
			},
			Name:     "显示文件",
			Method:   "GET",
			Path:     "/v1/host/SSH/ls",
			PermsTag: "host:ssh:ls",
			Desc:     "显示文件",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 47,
			},
			Name:     "上传文件",
			Method:   "POST",
			Path:     "/v1/host/SSH/upload",
			PermsTag: "host:ssh:upload",
			Desc:     "上传文件",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 48,
			},
			Name:     "下载文件",
			Method:   "GET",
			Path:     "/v1/host/ssh/download",
			PermsTag: "host:ssh:download",
			Desc:     "下载文件",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 49,
			},
			Name:     "删除文件",
			Method:   "DELETE",
			Path:     "/v1/host/ssh/rm",
			PermsTag: "host:ssh:rm",
			Desc:     "删除文件",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 50,
			},
			Name:     "主机信息",
			Method:   "GET",
			Path:     "/v1/host/info/:hostId",
			PermsTag: "host:info",
			Desc:     "查询主机信息",
			Creator:  creator,
			ParentId: 40,
		},
		{
			Model: models.Model{
				Id: 51,
			},
			Name:     "会话管理",
			Method:   "",
			Path:     "",
			PermsTag: "session:all",
			Desc:     "会话管理",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 52,
			},
			Name:     "查询会话",
			Method:   "GET",
			Path:     "/v1/host/session/list",
			PermsTag: "session:list",
			Desc:     "查询会话",
			Creator:  creator,
			ParentId: 51,
		},
		{
			Model: models.Model{
				Id: 53,
			},
			Name:     "注销会话",
			Method:   "DELETE",
			Path:     "/v1/host/session/delete",
			PermsTag: "session:delete",
			Desc:     "注销会话",
			Creator:  creator,
			ParentId: 51,
		},
		{
			Model: models.Model{
				Id: 54,
			},
			Name:     "操作审计",
			Method:   "",
			Path:     "",
			PermsTag: "record:all",
			Desc:     "操作审计",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 55,
			},
			Name:     "查询记录",
			Method:   "GET",
			Path:     "/v1/host/record/list",
			PermsTag: "record:list",
			Desc:     "查询记录",
			Creator:  creator,
			ParentId: 54,
		},
		{
			Model: models.Model{
				Id: 56,
			},
			Name:     "删除记录",
			Method:   "DELETE",
			Path:     "/v1/host/record/delete",
			PermsTag: "record:delete",
			Desc:     "查询记录",
			Creator:  creator,
			ParentId: 54,
		},
		{
			Model: models.Model{
				Id: 57,
			},
			Name:     "播放录像",
			Method:   "GET",
			Path:     "/v1/host/record/play",
			PermsTag: "record:play",
			Desc:     "播放录像",
			Creator:  creator,
			ParentId: 54,
		},
		{
			Model: models.Model{
				Id: 58,
			},
			Name:     "主机分组",
			Method:   "",
			Path:     "",
			PermsTag: "host:group:all",
			Desc:     "主机分组",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 59,
			},
			Name:     "分组列表",
			Method:   "GET",
			Path:     "/v1/host/group/list",
			PermsTag: "host:group:list",
			Desc:     "主机分组列表",
			Creator:  creator,
			ParentId: 58,
		},
		{
			Model: models.Model{
				Id: 60,
			},
			Name:     "创建分组",
			Method:   "POST",
			Path:     "/v1/host/group/create",
			PermsTag: "host:group:create",
			Desc:     "创建主机分组",
			Creator:  creator,
			ParentId: 58,
		},
		{
			Model: models.Model{
				Id: 61,
			},
			Name:     "更新分组",
			Method:   "PATCH",
			Path:     "/v1/host/group/update/:groupId",
			PermsTag: "host:group:update",
			Desc:     "更新主机分组",
			Creator:  creator,
			ParentId: 58,
		},
		{
			Model: models.Model{
				Id: 62,
			},
			Name:     "删除分组",
			Method:   "DELETE",
			Path:     "/v1/host/group/delete",
			PermsTag: "host:group:delete",
			Desc:     "删除主机分组",
			Creator:  creator,
			ParentId: 58,
		},
		{
			Model: models.Model{
				Id: 63,
			},
			Name:     "部门管理",
			Method:   "",
			Path:     "",
			PermsTag: "dept:all",
			Desc:     "部门管理",
			Creator:  creator,
			ParentId: 0,
		},
		{
			Model: models.Model{
				Id: 64,
			},
			Name:     "部门列表",
			Method:   "GET",
			Path:     "/v1/dept/list",
			PermsTag: "dept:list",
			Desc:     "部门列表",
			Creator:  creator,
			ParentId: 63,
		},
		{
			Model: models.Model{
				Id: 65,
			},
			Name:     "创建部门",
			Method:   "POST",
			Path:     "/v1/dept/create",
			PermsTag: "dept:create",
			Desc:     "创建部门",
			Creator:  creator,
			ParentId: 63,
		},
		{
			Model: models.Model{
				Id: 66,
			},
			Name:     "更新部门",
			Method:   "PATCH",
			Path:     "/v1/dept/update/:deptId",
			PermsTag: "dept:update",
			Desc:     "更新部门",
			Creator:  creator,
			ParentId: 63,
		},
		{
			Model: models.Model{
				Id: 67,
			},
			Name:     "删除部门",
			Method:   "DELETE",
			Path:     "/v1/dept/delete",
			PermsTag: "dept:delete",
			Desc:     "删除部门",
			Creator:  creator,
			ParentId: 63,
		},
	}
	for _, api := range apis {
		oldApi := system.SysApi{}
		err := common.Mysql.Where("id = ?", api.Id).First(&oldApi).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.Mysql.Create(&api)
			s := dao.New()
			// 管理员拥有所有API权限role[0]
			//_, err = s.CreateRoleCasbin(system.SysRoleCasbin{
			//	Keyword: roles[0].Keyword,
			//	Path:    api.Path,
			//	Method:  api.Method,
			//})
			// 来宾权限

			if api.ParentId == 1 {
				_, err = s.CreateRoleCasbin(system.SysRoleCasbin{
					Keyword: roles[1].Keyword,
					Path:    api.Path,
					Method:  api.Method,
				})
			}

		}
	}
}
