package main

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level       zapcore.Level `default:"debug"`
	LogEncoding string        `required:"true"`
}

func main() {
	appCfg := Config{}
	envconfig.MustProcess("MYAPP", &appCfg)

	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(appCfg.Level),
		DisableCaller:    true,
		Development:      true,
		Encoding:         appCfg.LogEncoding,
		OutputPaths:      []string{"stdout", "file_log.log"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	}

	logger := zap.Must(cfg.Build()).Sugar()

	logger.Info("Started")
	logger.Debug("Debug mode enabled")

	logger.With("count", 1).Info("One banana")
	logger.Desugar().With(zap.Int("count", 1)).Info("Two banana, no sugar")

}

// To show:
// 1. init, config
// 2. fields (sugar, desugar)
