package pkg

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func ComputeKey(appName string, serviceName string) string {
	target := fmt.Sprintf("/%s/%s", appName, serviceName)
	return target
}

func NewClientV3(endPoints []string, timeOut uint64) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Duration(timeOut) * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return cli
}
