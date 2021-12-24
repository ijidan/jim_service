package pkg

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type ServiceDiscovery struct {
	Client               *clientv3.Client
	AppName              string
	ServiceServerListMap map[string][]string
	lock                 sync.Mutex
}

func (d *ServiceDiscovery) InitServerList() []string {
	key:=ComputePrefixKey(d.AppName)
	rsp, err := d.Client.Get(context.Background(), key,clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	if rsp.Count == 0 {
		return nil
	}
	var serverList []string
	for _, kv := range rsp.Kvs {
		k:=string(kv.Key)
		v:=string(kv.Value)
		serviceName:=ExtractServiceName(d.AppName,k)
		d.ServiceServerListMap[serviceName]=append(d.ServiceServerListMap[serviceName],v)
	}
	return serverList
}
func (d *ServiceDiscovery) GetServiceServerList()map[string][]string  {
	return d.ServiceServerListMap
}

func (d *ServiceDiscovery) GetServerList(serviceName string) []string {
	return d.ServiceServerListMap[serviceName]
}

func (d *ServiceDiscovery) PutServer(serviceName string,server string) {
	d.lock.Lock()
	serverList:=d.ServiceServerListMap[serviceName]
	newServerList:=append(serverList,server)
	d.ServiceServerListMap[serviceName]=newServerList
	d.lock.Unlock()

}

func (d *ServiceDiscovery) DeleteServer(serviceName string,server string)  {
	d.lock.Lock()
	serverList:=d.ServiceServerListMap[serviceName]
	var newServerList []string
	for _,v:=range serverList{
		if v!=server{
			newServerList=append(newServerList,v)
		}
	}
	d.ServiceServerListMap[serviceName]=newServerList
	d.lock.Unlock()
}

func (d *ServiceDiscovery) Watch() {
	prefixKey := fmt.Sprintf("%s/", d.AppName)
	watchCh := d.Client.Watch(context.Background(), prefixKey, clientv3.WithPrefix())
	for {
		select {
		case rsp := <-watchCh:
			for _, event := range rsp.Events {
				key:=string(event.Kv.Key)
				value:=string(event.Kv.Value)
				serviceName:=ExtractServiceName(d.AppName,key)
				switch event.Type {
				case mvccpb.PUT:
					d.PutServer(serviceName,value)
					break
				case mvccpb.DELETE:
					d.DeleteServer(serviceName,value)
					break
				}
			}
		}
	}
}

func NewServiceDiscovery(client *clientv3.Client, appName string) *ServiceDiscovery {
	serviceDiscovery := &ServiceDiscovery{
		Client:  client,
		AppName: appName,
		ServiceServerListMap: map[string][]string{},
		lock: sync.Mutex{},
	}
	serviceDiscovery.InitServerList()
	go serviceDiscovery.Watch()
	return serviceDiscovery
}
