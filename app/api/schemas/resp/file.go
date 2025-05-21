package resp

type FileResp struct {
	ID       uint   `json:"id" structs:"id"`
	Cid      uint   `json:"cid" structs:"cid"`
	UserId   uint   `json:"user_id" structs:"user_id"`
	Type     string `json:"type" structs:"type"`
	Name     string `json:"name" structs:"name"`
	Uri      string `json:"uri" structs:"uri"`
	Ext      string `json:"ext" structs:"ext"`
	Size     int64  `json:"size" structs:"size"`
	TenantId uint   `json:"tenant_id" structs:"tenant_id"`
}
