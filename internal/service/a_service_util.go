package service

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

func ComputePrefixKey() string  {
	target:="service_"
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

func GetRpcConn(serviceName string) (*grpc.ClientConn,error) {
	target:=fmt.Sprintf("%s:///%s",JScheme,serviceName)
	conn, err := grpc.Dial(target,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
	return conn,err
}
