package req

type TenantAddReq struct {
	Name        string `json:"name" form:"name" validate:"required,max=125" label:"租户名称"`
	Address     string `json:"address" form:"address" validate:"required,max=255" label:"租户地址"`
	Phone       string `json:"phone" form:"phone" validate:"required,phone" label:"租户手机"`
	Email       string `json:"email" form:"email" validate:"required,email" label:"租户邮箱"`
	Domain      string `json:"domain" form:"domain" validate:"required,domain" label:"租户域名"`
	Logo        string `json:"logo" form:"logo" validate:"omitempty,url" label:"租户Logo"`
	Description string `json:"description" form:"description" validate:"max=255" label:"租户描述"`
	IsDisable   uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	ExpiredAt   int64  `json:"expired_at" form:"expired_at" validate:"omitempty,gte=0" label:"租户过期时间"`
}

type TenantEditReq struct {
	ID          uint   `json:"id" form:"id" validate:"required,gt=0" label:"租户ID"`
	Name        string `json:"name" form:"name" validate:"required,max=125" label:"租户名称"`
	Address     string `json:"address" form:"address" validate:"required,max=255" label:"租户地址"`
	Phone       string `json:"phone" form:"phone" validate:"required,phone" label:"租户手机"`
	Email       string `json:"email" form:"email" validate:"required,email" label:"租户邮箱"`
	Domain      string `json:"domain" form:"domain" validate:"required,domain" label:"租户域名"`
	Logo        string `json:"logo" form:"logo" validate:"omitempty,url" label:"租户Logo"`
	Description string `json:"description" form:"description" validate:"max=255" label:"租户描述"`
	IsDisable   uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	ExpiredAt   int64  `json:"expired_at" form:"expired_at" validate:"omitempty,gte=0" label:"租户过期时间"`
}

type TenantQueryReq struct {
	Name      string `json:"name" form:"name" validate:"max=125" label:"租户名称"`
	IsDisable uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
}

type UserAddReq struct {
	Username     string `json:"username" form:"username" validate:"required,min=4,max=32" label:"用户名"`
	Phone        string `json:"phone" form:"phone" validate:"required,phone" label:"手机号"`
	Password     string `json:"password" form:"password" validate:"required,min=6,max=32" label:"密码"`
	Email        string `json:"email" form:"email" validate:"email" label:"邮箱"`
	Avatar       string `json:"avatar" form:"avatar"`
	Nickname     string `json:"nickname" form:"nickname"`
	RoleId       uint   `json:"role_id" form:"role_id" validate:"gte=0" label:"角色ID"`
	IsMultipoint uint8  `json:"is_multipoint" form:"is_multipoint" validate:"oneof=0 1" label:"是否多点登录"`
	IsDisable    uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	TenantId     uint   `json:"tenant_id" form:"tenant_id" validate:"gte=0" label:"租户ID"`
}

type UserEditReq struct {
	ID           uint   `json:"id" form:"id" validate:"required,gte=1" label:"用户ID"`
	Username     string `json:"username" form:"username" validate:"required,min=4,max=32" label:"用户名"`
	Phone        string `json:"phone" form:"phone" validate:"required,phone" label:"手机号"`
	Email        string `json:"email" form:"email" validate:"email" label:"邮箱"`
	Avatar       string `json:"avatar" form:"avatar"`
	Nickname     string `json:"nickname" form:"nickname"`
	RoleId       uint   `json:"role_id" form:"role_id" validate:"gte=0" label:"角色ID"`
	IsMultipoint uint8  `json:"is_multipoint" form:"is_multipoint" validate:"oneof=0 1" label:"是否多点登录"`
	IsDisable    uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	TenantId     uint   `json:"tenant_id" form:"tenant_id" validate:"gte=0" label:"租户ID"`
}

type UserUpdateReq struct {
	ID              uint   `json:"id" form:"id" validate:"required,gte=1" label:"用户ID"`
	Avatar          string `json:"avatar" form:"avatar"`
	Password        string `json:"password" form:"password" validate:"min=6,max=32" label:"密码"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"eq=password" label:"确认密码"`
}

type UserQueryReq struct {
	Username  string `json:"username" form:"username" label:"用户名"`
	Phone     string `json:"phone" form:"phone" label:"手机号"`
	Email     string `json:"email" form:"email" label:"邮箱"`
	RoleId    uint   `json:"role_id" form:"role_id" validate:"gte=0" label:"角色ID"`
	IsDisable int8   `json:"is_disable" form:"is_disable" validate:"oneof=0 1 -1" default:"-1" label:"是否禁用"`
	TenantId  uint   `json:"tenant_id" form:"tenant_id" validate:"gte=0" label:"租户ID"`
}

type MenuAddReq struct {
	Pid           uint         `json:"pid" form:"pid" validate:"gte=0" label:"父菜单ID"`
	Name          string       `json:"name" form:"name" validate:"required,max=125" label:"菜单名称"`
	Path          string       `json:"path" form:"path" validate:"required,max=255" label:"菜单路径"`
	Redirect      string       `json:"redirect" form:"redirect" validate:"max=255" label:"重定向地址"`
	ComponentPath string       `json:"componentPath" form:"componentPath" validate:"max=255" label:"组件路径"`
	Title         string       `json:"title" form:"title" label:"页面标题"`                       // 页面标题
	Icon          string       `json:"icon" form:"icon" label:"图标"`                             // 图标
	RequiresAuth  bool         `json:"requiresAuth" form:"requiresAuth" label:"是否需要登录权限"` // 是否需要登录权限
	KeepAlive     bool         `json:"keepAlive" form:"keepAlive" label:"是否开启页面缓存"`       // 是否开启页面缓存
	Hide          bool         `json:"hide" form:"hide" label:"是否隐藏"`                         // 不显示在菜单中
	Sort          uint         `json:"sort" form:"sort" label:"菜单排序"`                         // 菜单排序
	Href          string       `json:"href" form:"href" label:"嵌套外链"`                         // 嵌套外链
	ActiveMenu    string       `json:"activeMenu" form:"activeMenu" label:"当前路由高亮菜单"`     // 当前路由高亮菜单
	WithoutTab    bool         `json:"withoutTab" form:"withoutTab" label:"是否添加到Tab"`        // 不添加到Tab
	PinTab        bool         `json:"pinTab" form:"pinTab" label:"固定Tab"`                      // 固定Tab
	MenuType      string       `json:"menuType" form:"menuType" label:"菜单类型"`                 // dir or page
	Button        []MenuButton `json:"button" form:"button" label:"按钮"`
}
type MenuEditReq struct {
	ID            uint         `json:"id" form:"id"`
	Pid           uint         `json:"pid" form:"pid" validate:"gte=0" label:"父菜单ID"`
	Name          string       `json:"name" form:"name" validate:"required,max=125" label:"菜单名称"`
	Path          string       `json:"path" form:"path" validate:"required,max=255" label:"菜单路径"`
	Redirect      string       `json:"redirect" form:"redirect" validate:"max=255" label:"重定向地址"`
	ComponentPath string       `json:"componentPath" form:"componentPath" validate:"max=255" label:"组件路径"`
	Title         string       `json:"title" form:"title" label:"页面标题"`                       // 页面标题
	Icon          string       `json:"icon" form:"icon" label:"图标"`                             // 图标
	RequiresAuth  bool         `json:"requiresAuth" form:"requiresAuth" label:"是否需要登录权限"` // 是否需要登录权限
	KeepAlive     bool         `json:"keepAlive" form:"keepAlive" label:"是否开启页面缓存"`       // 是否开启页面缓存
	Hide          bool         `json:"hide" form:"hide" label:"是否隐藏"`                         // 不显示在菜单中
	Sort          uint         `json:"sort" form:"sort" label:"菜单排序"`                         // 菜单排序
	Href          string       `json:"href" form:"href" label:"嵌套外链"`                         // 嵌套外链
	ActiveMenu    string       `json:"activeMenu" form:"activeMenu" label:"当前路由高亮菜单"`     // 当前路由高亮菜单
	WithoutTab    bool         `json:"withoutTab" form:"withoutTab" label:"是否添加到Tab"`        // 不添加到Tab
	PinTab        bool         `json:"pinTab" form:"pinTab" label:"固定Tab"`                      // 固定Tab
	MenuType      string       `json:"menuType" form:"menuType" label:"菜单类型"`                 // dir or page
	Button        []MenuButton `json:"button" form:"button" label:"按钮"`
}
type MenuButton struct {
	ID    uint   `json:"id" form:"id" label:"按钮ID"`
	Title string `json:"title" form:"title" label:"按钮名称"`
	Name  string `json:"name" form:"name" label:"按钮名称"`
	Sort  uint   `json:"sort" form:"sort" label:"按钮排序"`
}

type RoleAddReq struct {
	Name      string `json:"name" form:"name" validate:"required,max=125" label:"角色名称"`
	Remark    string `json:"remark" form:"remark" validate:"max=255" label:"角色描述"`
	IsDisable uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	Sort      uint16 `json:"sort" form:"sort" label:"角色排序"`
	IsAdmin   uint8  `json:"is_admin" form:"is_admin" validate:"oneof=0 1" label:"是否为管理员"`
	MenuIds   string `json:"menu_ids" form:"menu_ids" label:"菜单ID"`
}

type RoleEditReq struct {
	ID        uint   `json:"id" form:"id" validate:"required,gt=0" label:"角色ID"`
	Name      string `json:"name" form:"name" validate:"required,max=125" label:"角色名称"`
	Remark    string `json:"remark" form:"remark" validate:"max=255" label:"角色描述"`
	IsDisable uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	Sort      uint16 `json:"sort" form:"sort" label:"角色排序"`
	IsAdmin   uint8  `json:"is_admin" form:"is_admin" validate:"oneof=0 1" label:"是否为管理员"`
	MenuIds   string `json:"menu_ids" form:"menu_ids" label:"菜单ID"`
}

type RoleQueryReq struct {
	Name      string `json:"name" form:"name" label:"角色名称"`
	IsDisable int8   `json:"is_disable" form:"is_disable" validate:"oneof=0 1 -1" default:"-1" label:"是否禁用"`
}
