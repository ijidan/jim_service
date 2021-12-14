package pkg

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"jim/config"
	"sync"
)

var (
	onceRd     sync.Once
	instanceRd redis.Conn
)

func GetRdInstance(config *config.Config) redis.Conn {
	onceRd.Do(func() {
		var err error
		addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
		instanceRd, err = redis.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}
	})
	return instanceRd
}
