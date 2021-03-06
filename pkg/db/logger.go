package db

import (
	"context"

	dbLogger "github.com/titpetric/factory/logger"
	"go.uber.org/zap"

	"github.com/cortezaproject/corteza-server/pkg/logger"
)

type (
	zapLogger struct {
		logger *zap.Logger
	}
)

func NewZapLogger(logger *zap.Logger) *zapLogger {
	return &zapLogger{
		logger: logger,
	}
}

func (z *zapLogger) Log(ctx context.Context, msg string, fields ...dbLogger.Field) {
	// @todo when factory.DatabaseProfilerContext gets access to context from
	//       db functions, try to extract RequestID with middleware.GetReqID()

	zapFields := []zap.Field{}
	for _, v := range fields {
		zapFields = append(zapFields, zap.Any(v.Name(), v.Value()))
	}

	logger.
		AddRequestID(ctx, z.logger).
		Debug(msg, zapFields...)
}
