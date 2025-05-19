package req

type TenantAddReq struct {
	Name        string `json:"name" form:"name" validate:"required,max=125" label:"租户名称"`
	Address     string `json:"address" form:"address" validate:"required,max=255" label:"租户地址"`
	Phone       string `json:"phone" form:"phone" validate:"required,phone" label:"租户手机"`
	Email       string `json:"email" form:"email" validate:"required,email" label:"租户邮箱"`
	Domain      string `json:"domain" form:"domain" validate:"required,domain" label:"租户域名"`
	Logo        string `json:"logo" form:"logo" validate:"omitempty,url" label:"租户Logo"`
	Description string `json:"description" form:"description" validate:"max=255" label:"租户描述"`
	Menus       []uint `json:"menus" form:"menus" label:"菜单ID"`
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
	Menus       []uint `json:"menus" form:"menus" label:"菜单ID"`
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
	Pid        uint         `json:"pid" form:"pid" validate:"required,gt=0" label:"父菜单ID"`
	Title      string       `json:"title" form:"title" validate:"required,max=125" label:"菜单名称(中文)"`
	Name       string       `json:"name" form:"name" validate:"required,max=125" label:"菜单名称(英文)"`
	Path       string       `json:"path" form:"path" validate:"required,max=255" label:"菜单路径"`
	Component  string       `json:"component" form:"component" validate:"required,max=255" label:"菜单组件"`
	Icon       string       `json:"icon" form:"icon" validate:"max=255" label:"菜单图标"`
	MenuType   string       `json:"menu_type" form:"menu_type" validate:"required,oneof=menu action" label:"菜单类型"`
	Cacheable  bool         `json:"cacheable" form:"cacheable" validate:"oneof=0 1" label:"是否缓存"`
	RenderMenu bool         `json:"render_menu" form:"render_menu" validate:"oneof=0 1" label:"是否渲染菜单"`
	Permission string       `json:"permission" form:"permission" validate:"max=255" label:"权限标识"`
	Sort       uint16       `json:"sort" form:"sort" label:"菜单排序"`
	Target     string       `json:"target" form:"target" validate:"max=255" label:"打开方式"`
	Badge      string       `json:"badge" form:"badge" validate:"max=255" label:"徽标"`
	Button     []MenuButton `json:"button" form:"button" label:"按钮"`
}
type MenuEditReq struct {
	ID         uint         `json:"id" validate:"required,gt=0" form:"id"`
	Pid        uint         `json:"pid" form:"pid" validate:"required,gt=0" label:"父菜单ID"`
	Title      string       `json:"title" form:"title" validate:"required,max=125" label:"菜单名称(中文)"`
	Name       string       `json:"name" form:"name" validate:"required,max=125" label:"菜单名称(英文)"`
	Path       string       `json:"path" form:"path" validate:"required,max=255" label:"菜单路径"`
	Component  string       `json:"component" form:"component" validate:"required,max=255" label:"菜单组件"`
	Icon       string       `json:"icon" form:"icon" validate:"max=255" label:"菜单图标"`
	MenuType   string       `json:"menu_type" form:"menu_type" validate:"required,oneof=menu action" label:"菜单类型"`
	Cacheable  bool         `json:"cacheable" form:"cacheable" validate:"oneof=0 1" label:"是否缓存"`
	RenderMenu bool         `json:"render_menu" form:"render_menu" validate:"oneof=0 1" label:"是否渲染菜单"`
	Permission string       `json:"permission" form:"permission" validate:"max=255" label:"权限标识"`
	Sort       uint16       `json:"sort" form:"sort" label:"菜单排序"`
	Target     string       `json:"target" form:"target" validate:"max=255" label:"打开方式"`
	Badge      string       `json:"badge" form:"badge" validate:"max=255" label:"徽标"`
	Button     []MenuButton `json:"button" form:"button" label:"按钮"`
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
	Menus     []uint `json:"menus" form:"menus" label:"菜单ID"`
}

type RoleEditReq struct {
	ID        uint   `json:"id" form:"id" validate:"required,gt=0" label:"角色ID"`
	Name      string `json:"name" form:"name" validate:"required,max=125" label:"角色名称"`
	Remark    string `json:"remark" form:"remark" validate:"max=255" label:"角色描述"`
	IsDisable uint8  `json:"is_disable" form:"is_disable" validate:"oneof=0 1" label:"是否禁用"`
	Sort      uint16 `json:"sort" form:"sort" label:"角色排序"`
	IsAdmin   uint8  `json:"is_admin" form:"is_admin" validate:"oneof=0 1" label:"是否为管理员"`
	Menus     []uint `json:"menus" form:"menus" label:"菜单ID"`
}

type RoleQueryReq struct {
	Name      string `json:"name" form:"name" label:"角色名称"`
	IsDisable int8   `json:"is_disable" form:"is_disable" validate:"oneof=0 1 -1" default:"-1" label:"是否禁用"`
}

type CaptchaReq struct {
	Angle      int    `json:"angle" form:"angle" validate:"required" label:"角度"`
	CaptchaKey string `json:"captcha_key" form:"captcha_key" validate:"required" label:"验证码key"`
}
