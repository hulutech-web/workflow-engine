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
	IsDisable   uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	ExpiredAt   int64  `gorm:"not null;default:0;comment:'过期时间''"`
	SoftDelete
}

type AuthMenu struct {
	Model
	Pid        uint   `gorm:"column:pid"`
	Title      string `gorm:"column:title"`
	Name       string `gorm:"column:name"`
	Path       string `gorm:"column:path"`
	Component  string `gorm:"column:component"`
	Icon       string `gorm:"column:icon"`
	MenuType   string `json:"menu_type" gorm:"column:menu_type"` // page, action
	Cacheable  bool   `json:"cacheable" gorm:"column:cacheable"`
	RenderMenu bool   `json:"renderM_mnu" gorm:"column:render_menu"`
	Permission string `json:"permission" gorm:"column:permission"`
	Sort       uint16 `json:"sort" gorm:"column:sort"`
	Target     string `json:"target" gorm:"column:target"`
	Badge      string `json:"badge" gorm:"column:badge"`
}

type AuthRole struct {
	Model
	Name      string `gorm:"not null;default:'';comment:'角色名称''"`
	Remark    string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort      uint16 `gorm:"not null;default:0;comment:'角色排序'"`
	TenantId  uint   `gorm:"not null;default:0;comment:'租户ID'"`
	IsAdmin   uint8  `gorm:"not null;default:0;comment:'是否管理员'"`
	SoftDelete
}

type AuthPerm struct {
	ID     string `gorm:"not null;default:'';primary_key;comment:'权限标识'"`
	Type   string `gorm:"not null;default:'';comment:'权限类型: role=角色, tenant=租户';index"`
	TypeId uint   `gorm:"not null;default:0;comment:'权限类型ID'"`
	MenuId uint   `gorm:"not null;default:0;comment:'菜单ID'"`
}
