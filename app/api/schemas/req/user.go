package req

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
