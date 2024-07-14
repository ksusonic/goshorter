package app

import "go.uber.org/zap"

func newLogger() *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.DisableCaller = true
	zapConfig.Level.SetLevel(zap.DebugLevel)

	return zap.Must(zapConfig.Build())
}
