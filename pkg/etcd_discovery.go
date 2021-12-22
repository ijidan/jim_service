package pkg

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ServiceDiscovery struct {
	Client      *clientv3.Client
	AppName     string
}

func (d *ServiceDiscovery) GetServerList(serviceName string) []string {
	key := ComputeKey(d.AppName, serviceName)
	rsp, err := d.Client.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}
	if rsp.Count == 0 {
		return nil
	}
	var serverList []string
	for _,v:=range rsp.Kvs{
		value:=string(v.Value)
		serverList=append(serverList,value)
	}
	return serverList
}

func NewServiceDiscovery(client *clientv3.Client, appName string) *ServiceDiscovery  {
	serviceDiscovery:=&ServiceDiscovery{
		Client: client,
		AppName: appName,
	}
	return serviceDiscovery
}


