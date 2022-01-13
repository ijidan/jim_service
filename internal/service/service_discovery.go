package service

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type Discovery struct {
	Client               *clientv3.Client
	ServiceServerListMap map[string][]string
	lock                 sync.Mutex
}

func (d *Discovery) InitServerList() []string {
	key:=ComputePrefixKey()
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
		d.ServiceServerListMap[k]=append(d.ServiceServerListMap[k],v)
	}
	return serverList
}
func (d *Discovery) GetServiceServerList()map[string][]string  {
	return d.ServiceServerListMap
}

func (d *Discovery) GetServerList(serviceName string) []string {
	return d.ServiceServerListMap[serviceName]
}

func (d *Discovery) PutServer(serviceName string,server string) {
	d.lock.Lock()
	serverList:=d.ServiceServerListMap[serviceName]
	newServerList:=append(serverList,server)
	d.ServiceServerListMap[serviceName]=newServerList
	d.lock.Unlock()

}

func (d *Discovery) DeleteServer(serviceName string,server string)  {
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

func (d *Discovery) Watch() {
	key:=ComputePrefixKey()
	watchCh := d.Client.Watch(context.Background(), key, clientv3.WithPrefix())
	for {
		select {
		case rsp := <-watchCh:
			for _, event := range rsp.Events {
				key:=string(event.Kv.Key)
				value:=string(event.Kv.Value)
				switch event.Type {
				case mvccpb.PUT:
					d.PutServer(key,value)
					break
				case mvccpb.DELETE:
					d.DeleteServer(key,value)
					break
				}
			}
		}
	}
}

func NewServiceDiscovery(client *clientv3.Client) *Discovery {
	serviceDiscovery := &Discovery{
		Client:  client,
		ServiceServerListMap: map[string][]string{},
		lock: sync.Mutex{},
	}
	serviceDiscovery.InitServerList()
	go serviceDiscovery.Watch()
	return serviceDiscovery
}
