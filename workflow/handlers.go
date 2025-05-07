package workflow

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/auth"
	"github.com/hulutech-web/workflow-engine/http"
)

type Handlers struct {
	engine *Engine
}

func NewHandlers(engine *Engine) *Handlers {
	return &Handlers{engine: engine}
}

func (h *Handlers) StartProcess(ctx http.Context) error {
	var input struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description"`
		Steps       []Step `json:"steps" validate:"required,min=1"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, _ := ctx.Get("user")
	process := &Process{
		Name:        input.Name,
		Description: input.Description,
		Steps:       input.Steps,
		Status:      "draft",
	}

	if err := h.engine.StartProcess(process); err != nil {
		return ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, process)
}

func (h *Handlers) ApproveStep(ctx http.Context) error {
	var input struct {
		ProcessID uint   `json:"process_id" validate:"required"`
		StepID    uint   `json:"step_id" validate:"required"`
		Comment   string `json:"comment"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := ctx.Get("user").(auth.User)

	if err := h.engine.ApproveStep(input.ProcessID, input.StepID, user.GetID(), input.Comment); err != nil {
		return ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, gin.H{"status": "approved"})
}

// 其他处理器方法...
