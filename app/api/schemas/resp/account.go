package resp

type AccountLoginResp struct {
	Token    string   `json:"token" structs:"token"`
	UserInfo UserResp `json:"userInfo" structs:"userInfo"`
}
