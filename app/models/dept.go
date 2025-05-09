package models

type Dept struct {
	Model
	DeptName   string `gorm:"column:dept_name;not null;default:''" json:"dept_name"`
	Pid        uint   `gorm:"column:pid;not null;default:0" json:"pid"`
	DirectorID int    `gorm:"column:director_id;not null;default:0" json:"director_id"` // 部门主管
	ManagerID  int    `gorm:"column:manager_id;not null;default:0" json:"manager_id"`   // 部门经理
	Rank       int    `gorm:"column:rank;not null;default:1" json:"rank"`
	Html       string `gorm:"column:html;null;default:''" json:"html"`
	Level      int    `gorm:"column:level;null;default:0" json:"level"`
	Director   *Emp   `gorm:"foreignkey:DirectorID"` // 关联主管
	Manager    *Emp   `gorm:"foreignkey:ManagerID"`  // 关联经理
}
