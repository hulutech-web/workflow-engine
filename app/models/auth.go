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
	Pid           uint   `gorm:"column:pid"`
	Name          string `gorm:"column:name"`
	Path          string `gorm:"column:path"`
	Redirect      string `gorm:"column:redirect"`
	ComponentPath string `gorm:"column:componentPath"`
	IsDisabled    bool   `gorm:"column:isDisabled;default:false"`
	Icon          string `gorm:"column:icon"`
	MenuType      string `json:"menuType" gorm:"column:menuType"`                       // dir or page
	Title         string `json:"title" gorm:"column:title"`                             // 页面标题
	RequiresAuth  bool   `json:"requiresAuth" gorm:"column:requiresAuth;default:false"` // 是否需要登录权限
	KeepAlive     bool   `json:"keepAlive" gorm:"column:keepAlive;default:false"`       // 是否开启页面缓存
	Hide          bool   `json:"hide" gorm:"column:hide;default:false"`                 // 不显示在菜单中
	Sort          uint   `json:"sort" gorm:"column:sort;default:0"`                     // 排序
	Href          string `json:"href" gorm:"column:href"`                               // 嵌套外链
	ActiveMenu    string `json:"activeMenu" gorm:"column:activeMenu"`                   // 当前路由高亮菜单
	WithoutTab    bool   `json:"withoutTab" gorm:"column:withoutTab;default:false"`     // 不添加到Tab
	PinTab        bool   `json:"pinTab" gorm:"column:pinTab;default:false"`             // 固定Tab
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
