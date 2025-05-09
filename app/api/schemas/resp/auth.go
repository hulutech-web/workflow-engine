package resp

type AccountLoginResp struct {
	Token string               `json:"token" structs:"token"`
	Info  AccountLoginInfoResp `json:"info" structs:"info"`
}

type AccountLoginInfoResp struct {
	Username  string `json:"username" structs:"username"`
	Nickname  string `json:"nickname" structs:"nickname"`
	Avatar    string `json:"avatar" structs:"avatar"`
	Email     string `json:"email" structs:"email"`
	IsDisable uint8  `json:"is_disable" structs:"is_disable"`
	RoleId    uint8  `json:"role" structs:"role"`
}
