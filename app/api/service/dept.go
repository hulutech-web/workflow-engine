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
	Store(models.Dept) (*models.Dept, error)
	Update(models.Dept) (*models.Dept, error)
	Show(id int) *models.Dept
	Destroy(id int) error
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

func (d deptService) Store(dept models.Dept) (*models.Dept, error) {
	tx := d.db.Model(&models.Dept{}).Create(&dept)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dept, nil
}

func (d deptService) Update(dept models.Dept) (*models.Dept, error) {
	return nil, nil

}

func (d deptService) Show(id int) *models.Dept {
	return nil

}

func (d deptService) Destroy(id int) error {
	return nil

}
func (d deptService) BindManager(manager_id int, dept_id int) (response.Response, error) {
	return response.Response{}, nil

}
func (d deptService) BindDirector(director_id int, dept_id int) (response.Response, error) {
	return response.Response{}, nil
}

func NewDeptService(db *gorm.DB) DeptService {
	return &deptService{db: db}
}
