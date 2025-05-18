package req

type ProcPass struct {
	ProcessID int    `json:"process_id"`
	Content   string `json:"content"`
}
type ProcUnPass struct {
	ProcID  int    `json:"proc_id"`
	Content string `json:"content"`
}
