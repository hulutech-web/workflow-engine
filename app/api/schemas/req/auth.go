package req

type AuthLoginReq struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type AuthTokenReq struct {
	Token string `json:"token" form:"token"`
}
