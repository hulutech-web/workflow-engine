package log

import (
	"context"
	"github.com/hulutech-web/workflow-engine/app/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LogConsumer struct {
	db      *gorm.DB
	logChan <-chan models.Log
}

func NewLogConsumer(db *gorm.DB, logChan <-chan models.Log) *LogConsumer {
	return &LogConsumer{
		db:      db,
		logChan: logChan,
	}
}

func (lc *LogConsumer) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case logEntry := <-lc.logChan:
				lc.saveLog(logEntry)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (lc *LogConsumer) saveLog(logEntry models.Log) {
	tx := lc.db.Create(&logEntry)
	if tx.Error != nil {
		// 记录错误日志
		zap.S().Error(tx.Error)
	}
}
