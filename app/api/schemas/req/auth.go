package req

type AccountLoginReq struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type AccountTokenReq struct {
	Token string `json:"token" form:"token" validate:"demo" label:"demo token"`
}
