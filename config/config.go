package config

import (
	"os"

	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

type Config struct {
	Debug   bool          `yaml:"debug,omitempty" default:"false" `
	Logger  LoggerConfig  `yaml:"logger,omitempty"`
	VCode   VCodeConfig   `yaml:"vcode"`
	YueMiao YueMiaoConfig `yaml:"yuemiao,omitempty"`
	JinNiu  JinNiuConfig  `yaml:"jinniu,omitempty"`
}

type LoggerConfig struct {
	Level string `yaml:"level,omitempty" default:"debug"`
	// json or text
	Format string `yaml:"format,omitempty" default:"json"`
	// file
	Output string `yaml:"output,omitempty" default:""`
}

type VCodeConfig struct {
	AppCode   string `yaml:"appCode"`
	AppKey    string `yaml:"appKey"`
	AppSecret string `yaml:"appSecret"`
}

type YueMiaoConfig struct {
	TK       string `yaml:"tk,omitempty"`
	Verbose  bool   `yaml:"verbose,omitempty" default:"false"`
	Province string `yaml:"province" default:""`
	City     string `yaml:"city" default:""`
	District string `yaml:"district" default:""`
	// 0:二价  1: 四价  3:九价
	Vaccines string `yaml:"vaccines" default:"3"`
	// 可配置多个，按照英文逗号分割
	DepartmentName string `yaml:"department_name" default:""`
	StartTime      string `yaml:"start_time"`

	// v2版本配置
	SeckillId string `yaml:"seckill_id"`
	LinkmanId string `yaml:"linkman_id"`
	IdCardNo  string `yaml:"id_card_no"`
}

type JinNiuConfig struct {
	StartTime      string `yaml:"start_time"`
	Cookie         string `yaml:"cookie,omitempty"`
	HisID          string `yaml:"his_id" default:"2153"`
	PlatformSource string `yaml:"platform_source" default:"1"`
	PlatformID     string `yaml:"platform_id" default:"2153"`
	Verbose        bool   `yaml:"verbose,omitempty" default:"false"`
}

var C *Config = &Config{}

func initLogger(debug bool, level, output string) {
	var conf zap.Config
	if debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	var zapLevel = zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zap.L().Panic("set logger level fail",
			zap.Strings("only", []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}),
			zap.Error(err),
		)
	}

	conf.Level = zapLevel
	conf.Encoding = "json"

	if output != "" {
		conf.OutputPaths = []string{output}
		conf.ErrorOutputPaths = []string{output}
	}

	logger, _ := conf.Build()

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
}

func Init(cfgFile string) {
	_ = os.Setenv("CONFIGOR_ENV_PREFIX", "-")

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	if cfgFile != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgFile); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}

	initLogger(C.Debug, C.Logger.Level, C.Logger.Output)
	zap.L().Debug("loaded config")
}
