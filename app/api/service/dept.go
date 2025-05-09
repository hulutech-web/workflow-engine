package service

import (
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"gorm.io/gorm"
)

/*
用来操作数据库查表
*/
type DeptService interface {
	Index() ([]models.Dept, error)
	List() ([]models.Dept, error)
	Store(models.Dept) (response.Response, error)
	Update(models.Dept) (response.Response, error)
	Show(id int) response.Response
	Destroy(id int) (response.Response, error)
	BindManager(manager_id int, dept_id int) (response.Response, error)
	BindDirector(director_id int, dept_id int) (response.Response, error)
}

type deptService struct {
	db *gorm.DB
}

func (d deptService) Index() ([]models.Dept, error) {
	var depts []models.Dept
	d.db.Model(&models.Dept{}).Find(&depts)
	return depts, nil
}

func (d deptService) List() ([]models.Dept, error) {
	return nil, nil
}

func (d deptService) Store(dept models.Dept) (response.Response, error) {
	return response.Response{}, nil
}

func (d deptService) Update(dept models.Dept) (response.Response, error) {
	return response.Response{}, nil

}

func (d deptService) Show() response.Response {
	return response.Response{}

}

func (d deptService) Destroy(id int) (response.Response, error) {
	return response.Response{}, nil

}
func (d deptService) BindManager(manager_id int, dept_id int) (response.Response, error) {
	return response.Response{}, nil

}
func (d deptService) BindDirector(director_id int, dept_id int) (response.Response, error) {
	return response.Response{}, nil
}
