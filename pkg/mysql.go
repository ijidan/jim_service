package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"jim/config"
	"sync"
	"time"
)

var (
	onceDb     sync.Once
	instanceDb *gorm.DB
)

func GetDbInstance(conf *config.Config) *gorm.DB {
	onceDb.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Database, conf.Mysql.Charset)
		instanceDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic(err)
		}
		sqlDB, _ := instanceDb.DB()
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return instanceDb
}
