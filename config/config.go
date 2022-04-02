package config

import (
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

type Config struct {
	Debug bool `yaml:"debug,omitempty" default:"false"`

	Pprof      PprofConfig  `yaml:"pprof,omitempty"`
	Logger     LoggerConfig `yaml:"logger,omitempty"`
	TK         string       `yaml:"tk"`
	MemberID   string       `yaml:"member_id"`
	IDCard     string       `yaml:"id_card"`
	RegionCode string       `yaml:"region_code" default:"5101"`
	Cookie     string       `yaml:"cookie"`
	VaccineID  string       `yaml:"vaccine_id"`
	Total      int          `yaml:"total" default:"5"`
	Step       int          `yaml:"step" default:"200"`
}

type LoggerConfig struct {
	Level string `yaml:"level,omitempty" default:"debug"`
	// json or text
	Format string `yaml:"format,omitempty" default:"json"`
	// file
	Output string `yaml:"output,omitempty" default:""`
}

type PprofConfig struct {
	Enabled bool   `yaml:"enabled" default:"false"`
	Addr    string `yaml:"addr" default:":32999"`
}

var C *Config

func Init(cfgFile string) {
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

	zap.L().Debug("loaded config")
}

func init() {
	C = &Config{}
}
