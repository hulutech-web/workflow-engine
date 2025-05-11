package req

type AccountLoginReq struct {
	TenantId uint   `json:"tenantId" validate:"required" validate:"gt=0" label:"租户"`
	Username string `json:"username" validate:"required,min=4,max=32" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=32" label:"密码"`
}

type AccountTokenReq struct {
	Token string `json:"token" form:"token" validate:"required" label:"Token"`
}

type AccountRegisterReq struct {
	TenantId uint   `json:"tenantId" validate:"required,gt=0" label:"租户"`
	Username string `json:"username" validate:"required,min=4,max=32" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=32" label:"密码"`
	Confirm  string `json:"confirm" validate:"required,eqfield=Password" label:"确认密码"`
	Email    string `json:"email" validate:"required,email" label:"邮箱"`
	Phone    string `json:"phone" validate:"required,phone" label:"手机号"`
}
