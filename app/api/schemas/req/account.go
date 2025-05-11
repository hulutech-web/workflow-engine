package req

type AccountLoginReq struct {
	TenantId uint   `json:"tenantId" validate:"required" validate:"gt=0" label:"租户"`
	Username string `json:"username" validate:"required,min=4,max=32" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=32" label:"密码"`
}

type AccountTokenReq struct {
	Token string `json:"token" form:"token" validate:"required" label:"Token"`
}
