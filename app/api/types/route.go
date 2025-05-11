package types

import (
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/pkg/log"
)

type ApiRouter struct {
	*gin.RouterGroup
	*log.LogCollector
}
