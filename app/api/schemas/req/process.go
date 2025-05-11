package req

type ProcessReq struct {
	FlowID int `json:"flow_id" validate:"required" validate:"gt=0" label:"流程id"`
}

type ProReq struct {
	FlowID uint   `json:"flow_id"`
	Left   string `json:"left"`
	Top    string `json:"top"`
}
