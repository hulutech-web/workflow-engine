package service

import (
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/models"
)

type UserService interface {
	FindByUsername(username string, tenantId string) (*models.User, error)
	List(page *req.PageReq)
}
