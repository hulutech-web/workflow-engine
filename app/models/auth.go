package models

type AuthTenant struct {
	Model
	Name        string `gorm:"not null;default:'';comment:'名称''"`
	Address     string `gorm:"not null;default:'';comment:'地址''"`
	Phone       string `gorm:"not null;default:'';comment:'电话''"`
	Email       string `gorm:"not null;default:'';comment:'邮箱''"`
	Domain      string `gorm:"not null;default:'';comment:'域名''"`
	Logo        string `gorm:"not null;default:'';comment:'logo''"`
	Description string `gorm:"not null;default:'';comment:'描述''"`
	IsStop      uint8  `gorm:"not null;default:0;comment:'状态''"`
	ExpiredAt   int64  `gorm:"not null;default:0;comment:'过期时间''"`
	SoftDelete
}

type AuthMenu struct {
	Model
	Pid       uint   `gorm:"not null;default:0;comment:'上级菜单'"`
	MenuType  string `gorm:"not null;default:'';comment:'权限类型: M=目录，C=菜单，A=按钮''"`
	MenuName  string `gorm:"not null;default:'';comment:'菜单名称'"`
	MenuIcon  string `gorm:"not null;default:'';comment:'菜单图标'"`
	MenuSort  uint16 `gorm:"not null;default:0;comment:'菜单排序'"`
	Perms     string `gorm:"not null;default:'';comment:'权限标识'"`
	Paths     string `gorm:"not null;default:'';comment:'路由地址'"`
	Component string `gorm:"not null;default:'';comment:'前端组件'"`
	Selected  string `gorm:"not null;default:'';comment:'选中路径'"`
	Params    string `gorm:"not null;default:'';comment:'路由参数'"`
	IsCache   uint8  `gorm:"not null;default:0;comment:'是否缓存: 0=否, 1=是''"`
	IsShow    uint8  `gorm:"not null;default:1;comment:'是否显示: 0=否, 1=是'"`
	IsDisable uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	SoftDelete
}

type AuthRole struct {
	Model
	Name      string `gorm:"not null;default:'';comment:'角色名称''"`
	Remark    string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort      uint16 `gorm:"not null;default:0;comment:'角色排序'"`
	TenantID  uint   `gorm:"not null;default:0;comment:'租户ID'"`
	IsAdmin   uint8  `gorm:"not null;default:0;comment:'是否管理员'"`
	SoftDelete
}
