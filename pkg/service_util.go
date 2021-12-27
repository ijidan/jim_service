package pkg

import (
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
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
func GetServerConfigMap()  string{
	serverConfigMap:=map[string]interface{}{"loadBalancingPolicy":roundrobin.Name}
	serverConfigBytes,_:=json.Marshal(serverConfigMap)
	serverConfigJson:=string(serverConfigBytes)
	return serverConfigJson
}
func GetRpcConn(serviceName string) (*grpc.ClientConn,error) {
	target:=fmt.Sprintf("%s:///%s",JScheme,serviceName)
	serverConfigJson:=GetServerConfigMap()
	conn, err := grpc.Dial(target,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(10),
		grpc.WithDefaultServiceConfig(serverConfigJson))
	return conn,err
}
