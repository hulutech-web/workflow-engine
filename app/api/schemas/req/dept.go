package req

type BindManagerReq struct {
	ManagerID int `json:"manager_id" form:"manager_id"`
	DeptID    int `json:"dept_id" form:"dept_id"`
}

type BindDirectorReq struct {
	DirectorID int `json:"director_id" form:"director_id"`
	DeptID     int `json:"dept_id" form:"dept_id"`
}
