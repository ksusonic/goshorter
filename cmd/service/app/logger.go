package app

import (
	"bytes"
	"io"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger() *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.DisableCaller = true
	zapConfig.Level.SetLevel(zap.DebugLevel)

	return zap.Must(zapConfig.Build())
}

func newGinZapConfig() *ginzap.Config {
	return &ginzap.Config{
		UTC:        true,
		TimeFormat: time.RFC3339,
		SkipPaths:  []string{ping},
		Context: func(c *gin.Context) []zapcore.Field {
			fields := make([]zapcore.Field, 0, len(c.Request.Header)+1)

			// log headers
			for name, value := range c.Request.Header {
				if len(value) == 0 {
					continue
				}

				fields = append(fields, zap.String(name, value[0]))
			}

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))

			return fields
		},
		DefaultLevel: zap.InfoLevel,
	}
}
