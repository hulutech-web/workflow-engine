package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/service"
	"github.com/hulutech-web/workflow-engine/app/api/types"
	"github.com/hulutech-web/workflow-engine/app/models"
	"github.com/hulutech-web/workflow-engine/pkg/plugin/response"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type flowlink struct {
	fx.In
	Srv service.FlowlinkService
}

func flowlinkRoutes(a flowlink, r *types.ApiRouter) {
	r.POST("/flowlink", a.Update)
}

func (r *flowlink) Update(ctx *gin.Context) {
	var flk models.Flow
	ctx.Bind(&flk)
	logrus.WithFields(logrus.Fields{
		"flk": flk,
	}).Info("绑定的flow")
	err := r.Srv.Update(ctx, flk)
	if err != nil {
		response.FailWithMsg(ctx, response.Failed, err.Error())
		return
	}
	response.OkWithData(ctx, "操作成功")
}
