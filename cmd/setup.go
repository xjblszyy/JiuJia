package cmd

import (
	"net/http"

	"go.uber.org/zap"
	"jiujia/config"
	http2 "jiujia/pkg/http"
)

// SetupPprofServer 开启 pprof
func SetupPprofServer() error {
	cfg := config.C.Pprof
	if !cfg.Enabled {
		return nil
	}
	zap.L().Sugar().Infof("启动 pprof 监听 %s.", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, nil)

	if err != nil {
		zap.L().Error("开启 pprof 监听失败 %s", zap.Error(err))
	}
	return nil
}

func SetupHttpClient() error {
	http2.InitClient("", config.C.Debug)
	return nil
}

func SetupLogger() error {
	var conf zap.Config
	if config.C.Debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	var zapLevel = zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(config.C.Logger.Level)); err != nil {
		zap.L().Panic("set logger level fail",
			zap.Strings("only", []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}),
			zap.Error(err),
		)
	}

	conf.Level = zapLevel
	conf.Encoding = "json"

	if config.C.Logger.Output != "" {
		conf.OutputPaths = []string{config.C.Logger.Output}
		conf.ErrorOutputPaths = []string{config.C.Logger.Output}
	}

	logger, _ := conf.Build()

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)

	return nil
}
