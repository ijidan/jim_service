package pkg

import (
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"jim_service/config"
	"path/filepath"
	"runtime"
)

var (
	Root      string
	Conf    *config.Config
	Logger    *logrus.Logger
	Db        *gorm.DB
	Rd        redis.Conn
	Rsp  *Response
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
	Conf = config.GetConfigInstance(Root)
	Logger = GetLoggerInstance(Conf, Root)
	Db = GetDbInstance(Conf)
	Rd = GetRdInstance(Conf)
	Rsp = GetResponseInstance()
	Tracer, Closer = NewJaeger(Conf, "jim_service")
	RequestId = "X-Request-Id"
}
