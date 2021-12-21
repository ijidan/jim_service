package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"io"
	"jim_service/config"
	"jim_service/pkg"
	"path/filepath"
	"runtime"
)

var (
	Root      string
	Config    *config.Config
	Logger    *logrus.Logger
	Db        *gorm.DB
	Rd        redis.Conn
	Response  *pkg.Response
	Tracer    opentracing.Tracer
	Closer    io.Closer
	RequestId string
	ClientV3 *clientv3.Client
	ServiceRegister *pkg.ServiceRegister
	ServiceDiscovery *pkg.ServiceDiscovery
)

func Close() {
	sqlDB, _ := Db.DB()
	_ = sqlDB.Close()
	_ = Rd.Close()
	_ = Closer.Close()
}
func init() {
	_, file, _, _ := runtime.Caller(0)
	Root = filepath.Dir(filepath.Dir(file))
	Config = config.GetConfigInstance(Root)
	Logger = pkg.GetLoggerInstance(Config, Root)
	Db = pkg.GetDbInstance(Config)
	Rd = pkg.GetRdInstance(Config)
	Response = pkg.GetResponseInstance()
	Tracer, Closer = pkg.NewJaeger(Config, "jim_service")
	RequestId = "X-Request-Id"
	ClientV3=pkg.NewClientV3(Config.Etcd.EndPoints,Config.Etcd.DialTimeout)
	ServiceRegister=pkg.NewServiceRegister(ClientV3,Config.App.Name)
}
