package setting

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func NewSetting() (*Setting, error) {
	vp := viper.New()

	configPath := os.Getenv("CONFIG_DIR")
	if configPath == "" {
		configPath = "configs/"
	}

	entries, err := os.ReadDir(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read configs directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			tempViper := viper.New()
			tempViper.SetConfigFile(filepath.Join("configs", entry.Name()))
			tempViper.SetConfigType("yaml")

			if err := tempViper.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("failed to read config file %s: %w", entry.Name(), err)
			}

			if err := vp.MergeConfigMap(tempViper.AllSettings()); err != nil {
				return nil, fmt.Errorf("failed to merge config file %s: %w", entry.Name(), err)
			}
		}
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadConfigField(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
