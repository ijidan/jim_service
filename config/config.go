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
	Pprof struct {
		Host string `yaml:"host"`
		Port uint `yaml:"port"`
	}
	Websocket struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
		Log  string `yaml:"log"`
	}
	Rpc struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
		Ttl  int64  `yaml:"ttl"`
		Log  string `yaml:"log"`
	}
	Ps struct {
		Host string `yaml:"host"`
		Port uint   `yaml:"port"`
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
		Secret       string `yaml:"secret"`
		DefaultToken string `yaml:"default_token"`
	}
	Pager struct {
		PageSize uint `yaml:"page_size"`
	}
	Etcd struct {
		Host    []string `yaml:"host"`
		Timeout uint64   `yaml:"timeout"`
	}
	Smms struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	Email struct {
		Smtp     string `yaml:"smtp"`
		Port     uint64 `yaml:"port"`
		Ssl      bool   `yaml:"ssl"`
		Account  string `yaml:"account"`
		Password string `yaml:"password"`
	}
	PubSub struct {
		Brokers []string `yaml:"brokers"`
	}
	Manager struct {
		Email map[string]string `yaml:"email"`
	}
	Gateway struct {
		Id []string `yaml:"id"`
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
