package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

const EnvLocal = "local"
const EnvTest = "test"
const EnvStage = "stage"
const EnvProduction = "production"

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Ver  string `yaml:"ver"`
		Env  string `yaml:"env"`
	}
	Http struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Websocket struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Rpc struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Mysql struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Charset  string `yaml:"charset"`
		Database string `yaml:"database"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     uint   `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		db       int    `yaml:"db"`
	}
	Jaeger struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
	Pager struct {
		PageSize uint `yaml:"page_size"`
	}
}

var (
	onceConfig     sync.Once
	instanceConfig *Config
)

func GetConfigInstance(root string) *Config {
	onceConfig.Do(func() {
		instanceConfig = &Config{}
		v := viper.New()
		v.AddConfigPath(root)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
		})
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := v.Unmarshal(instanceConfig); err != nil {
			panic(err)
		}
	})
	return instanceConfig
}
