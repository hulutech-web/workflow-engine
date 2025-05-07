package workflow

import (
	"errors"
	"time"
)

var (
	ErrProcessNotFound   = errors.New("process not found")
	ErrStepNotFound      = errors.New("step not found")
	ErrNotAuthorized     = errors.New("not authorized to approve")
	ErrInvalidTransition = errors.New("invalid status transition")
)

type Engine struct {
	repo Repository
}

type Repository interface {
	FindProcess(id uint) (*Process, error)
	SaveProcess(process *Process) error
}

func NewEngine(repo Repository) *Engine {
	return &Engine{repo: repo}
}

func (e *Engine) StartProcess(process *Process) error {
	process.Status = "active"
	process.CurrentStep = 0
	return e.repo.SaveProcess(process)
}

func (e *Engine) ApproveStep(processID uint, stepID uint, userID uint, comment string) error {
	process, err := e.repo.FindProcess(processID)
	if err != nil {
		return err
	}

	if process.Status != "active" {
		return ErrInvalidTransition
	}

	if int(stepID) != process.CurrentStep {
		return ErrStepNotFound
	}

	step := process.Steps[process.CurrentStep]

	// 检查用户是否有权限审批
	canApprove := false
	for _, approver := range step.Approvers {
		if approver.UserID == userID {
			canApprove = true
			break
		}
	}

	if !canApprove {
		return ErrNotAuthorized
	}

	// 更新审批状态
	for i, approver := range step.Approvers {
		if approver.UserID == userID {
			step.Approvers[i].Approved = true
			step.Approvers[i].Comment = comment
			step.Approvers[i].Timestamp = time.Now()
			break
		}
	}

	// 检查是否所有审批人都已审批
	allApproved := true
	for _, approver := range step.Approvers {
		if !approver.Approved {
			allApproved = false
			break
		}
	}

	if allApproved {
		process.CurrentStep++
		if process.CurrentStep >= len(process.Steps) {
			process.Status = "completed"
		}
	}

	return e.repo.SaveProcess(process)
}
