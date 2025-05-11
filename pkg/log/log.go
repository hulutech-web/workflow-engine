package log

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hulutech-web/workflow-engine/app/api/schemas/req"
	"github.com/hulutech-web/workflow-engine/app/models"
	"gorm.io/gorm"
	"time"
)

type LogCollector struct {
	LogChan chan models.Log
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func NewLogCollector(db *gorm.DB, bufferSize int) *LogCollector {
	lc := &LogCollector{
		LogChan: make(chan models.Log, bufferSize),
	}
	c := NewLogConsumer(db, lc.LogChan)
	c.Start(context.Background())
	return lc
}

func (lc *LogCollector) Log(source string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		start := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		duration := time.Since(start).Milliseconds()

		userClaims := req.GetAuth(c)

		operationType := getOperationType(c)

		log := models.Log{
			UserID: userClaims.UserId,
			Action: operationType,
			Detail: source,
			IP:     c.ClientIP(),
			Agent:  c.Request.UserAgent(),
			Status: getStatus(c.Writer.Status()),
			Uri:    c.Request.URL.Path,
			Error:  blw.body.String(),
			Time:   duration,
		}
		lc.LogChan <- log
	}
}

func getOperationType(c *gin.Context) string {
	switch c.Request.Method {
	case "GET":
		return "数据查询"
	case "POST":
		return "新增数据"
	case "PUT":
		return "更新数据"
	case "DELETE":
		return "删除数据"
	case "PATCH":
		return "部分更新数据"
	default:
		return "其他操作"
	}
}

func getStatus(statusCode int) string {
	if statusCode >= 200 && statusCode < 400 {
		return "成功"
	}
	return "失败"
}
