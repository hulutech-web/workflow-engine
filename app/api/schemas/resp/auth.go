package resp

type AuthLoginResp struct {
	Token string            `json:"token" structs:"token"`
	Info  AuthLoginInfoResp `json:"info" structs:"info"`
}

type AuthLoginInfoResp struct {
	Username string `json:"username" structs:"username"`
	Nickname string `json:"nickname" structs:"nickname"`
	Avatar   string `json:"avatar" structs:"avatar"`
	Email    string `json:"email" structs:"email"`
	Status   uint8  `json:"status" structs:"status"`
	RoleId   uint8  `json:"role" structs:"role"`
}
