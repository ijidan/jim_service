package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"jim/config"
	"jim/pkg"
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
	Tracer, Closer = pkg.NewJaeger(Config, "jim_api")
	RequestId = "X-Request-Id"
}
