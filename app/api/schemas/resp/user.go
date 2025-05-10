package resp

type UserResp struct {
	Username     string `json:"username" structs:"username"`
	Phone        string `json:"phone" structs:"phone"`
	Password     string `json:"password" structs:"password"`
	Salt         string `json:"salt" structs:"salt"`
	Email        string `json:"email" structs:"email"`
	Avatar       string `json:"avatar" structs:"avatar"`
	Nickname     string `json:"nickname" structs:"nickname"`
	IsMultipoint uint8  `json:"is_multipoint" structs:"is_multipoint"`
	IsDisable    uint8  `json:"is_disable" structs:"is_disable"`
	Role         struct {
		ID   uint   `json:"id" structs:"id"`
		Name string `json:"name" structs:"name"`
	} `json:"role" structs:"role"`
	Tenant struct {
		ID   uint   `json:"id" structs:"id"`
		Name string `json:"name" structs:"name"`
	} `json:"tenant" structs:"tenant"`
}
