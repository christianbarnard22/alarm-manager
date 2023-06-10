package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

type contextKey string

var (
	loggerKey contextKey = "alarm-manager"
)

func GetLogger(ctx context.Context) *logrus.Entry {
	if logger, ok := ctx.Value(loggerKey).(*logrus.Entry); ok {
		return logger
	}
	return logrus.NewEntry(logrus.StandardLogger())
}

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
