package req

type ProcessReq struct {
	FlowID int `json:"flow_id" validate:"required" validate:"gt=0" label:"流程id"`
}
