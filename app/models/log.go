package models

type Log struct {
	Model
	UserID uint   `json:"user_id" gorm:"not null;index"`
	Action string `json:"action" gorm:"not null;comment:动作"`
	Detail string `json:"detail" gorm:"not null;comment:详情"`
	IP     string `json:"ip" gorm:"not null;comment:IP"`
	Agent  string `json:"agent" gorm:"not null;comment:User-Agent"`
	Status string `json:"status" gorm:"not null;comment:状态"`
	Uri    string `json:"uri" gorm:"not null;comment:URI"`
	Error  string `json:"error" gorm:"comment:错误信息"`
	Time   int64  `json:"time" gorm:"not null;comment:操作耗时，单位为毫秒"`
}
