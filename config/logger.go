package config

import (
	"go.uber.org/zap"
)

type loggerConf struct {
	Mode   string
	Level  int8
	Output string
}

func (a *app) initLogger() {
	Logger, _ = zap.NewDevelopment()
	defer Logger.Sync()
	Logger.Info("instantiation", zap.String("type", "logger"), zap.String("source", "zap"), zap.String("status", "done"))
}
