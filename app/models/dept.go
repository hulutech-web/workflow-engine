package models

import (
	"strings"
)

type Dept struct {
	Model
	DeptName   string  `gorm:"column:dept_name;not null;default:''" json:"dept_name"`
	PID        uint    `gorm:"column:pid;not null;default:0" json:"pid"`
	DirectorID int     `gorm:"column:director_id;not null;default:0" json:"director_id"` // 部门主管
	ManagerID  int     `gorm:"column:manager_id;not null;default:0" json:"manager_id"`   // 部门经理
	Rank       int     `gorm:"column:rank;not null;default:1" json:"rank"`
	Html       string  `gorm:"column:html;null;default:''" json:"html"`
	Level      int     `gorm:"column:level;null;default:0" json:"level"`
	Director   *Emp    `gorm:"foreignkey:DirectorID" swaggerignore:"true"` // 关联主管
	Manager    *Emp    `gorm:"foreignkey:ManagerID" swaggerignore:"true"`  // 关联经理
	Children   []*Dept `gorm:"foreignkey:PID" swaggerignore:"true"`
}

func (d *Dept) BuildDeptTree(models []Dept, pid uint) []*Dept {
	// 构建快速查找的map
	deptMap := make(map[uint][]Dept)
	for _, dept := range models {
		deptMap[dept.PID] = append(deptMap[dept.PID], dept)
	}

	// 递归构建树形结构
	var buildTree func(pid uint, level int) []*Dept
	buildTree = func(pid uint, level int) []*Dept {
		var nodes []*Dept
		for _, dept := range deptMap[pid] {
			node := dept // 创建副本避免引用问题
			node.Html = strings.Repeat("--", level)
			node.Level = level

			// 递归构建子节点
			children := buildTree(node.ID, level+1)
			if len(children) > 0 {
				node.Children = children
			}

			nodes = append(nodes, &node)
		}
		return nodes
	}

	return buildTree(pid, 0)
}

// FlattenDeptsWithLevel 扁平化部门列表并添加层级信息
func (d *Dept) FlattenDeptsWithLevel(models []Dept, pid uint, level int) []Dept {
	var result []Dept
	for _, dept := range models {
		if dept.PID == pid {
			// 创建副本避免修改原始数据
			newDept := dept
			newDept.Html = strings.Repeat("--", level)
			newDept.Level = level

			result = append(result, newDept)
			result = append(result, d.FlattenDeptsWithLevel(models, newDept.ID, level+1)...)
		}
	}
	return result
}
