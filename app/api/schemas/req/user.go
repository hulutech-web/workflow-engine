package req

type UserAddReq struct {
	Username     string `json:"username" form:"username" validate:"required,min=4,max=32"`
	Phone        string `json:"phone" form:"phone" validate:"required,min=11,max=11"`
	Password     string `json:"password" form:"password" validate:"required,min=6,max=32"`
	Email        string
	Avatar       string
	Nickname     string
	RoleId       uint
	IsMultipoint uint8
	IsDisable    uint8
	TenantId     uint
}
