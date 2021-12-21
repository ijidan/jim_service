package pkg

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ServiceRegister struct {
	Client         *clientv3.Client
	AppName        string
	ServiceName    string
	ServiceAddress string
	Ttl            int64
	LeaseId        clientv3.LeaseID
	KeepAliveCh    <-chan *clientv3.LeaseKeepAliveResponse
	CloseCh        <-chan struct{}
}

func (r *ServiceRegister) RegisterService(serviceName string, serviceAddress string, ttl int64) {
	r.ServiceName=serviceName
	r.ServiceAddress=serviceAddress
	r.Ttl=ttl
	LeaseRsp, err1 := r.Client.Grant(context.Background(), r.Ttl)
	if err1 != nil {
		panic(err1)
	}
	r.LeaseId = LeaseRsp.ID
	var err2 error
	r.KeepAliveCh, err2 = r.Client.KeepAlive(context.Background(), r.LeaseId)
	if err2 != nil {
		panic(err2)
	}
	key := ComputeKey(r.AppName,r.ServiceName)
	_, err3 := r.Client.KV.Put(context.Background(), key, serviceAddress, clientv3.WithLease(LeaseRsp.ID))
	if err3 != nil {
		panic(err3)
	}
}

func (r *ServiceRegister) UnRegisterService() {
	key := ComputeKey(r.AppName,r.ServiceName)
	_, err := r.Client.KV.Delete(context.Background(), key)
	if err != nil {
		panic(err)
	}
}

func (r *ServiceRegister) KeepAlive() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case res := <-r.KeepAliveCh:
			if res != nil {
				r.RegisterService()
			}
		case <-ticker.C:
			if len(r.KeepAliveCh) > 0 {
				r.RegisterService()
			}
		}
	}
}

type ServiceDiscovery struct {
	Client         *clientv3.Client
	AppName        string
	ServiceName    string
}

func (d *ServiceDiscovery) GetServiceList(appName string,serviceName string)[]string  {
	key:=ComputeKey(appName,serviceName)
	rsp,err:=d.Client.Get(context.Background(),key)
	if err!=nil{
		panic(err)
	}
	if rsp.Count==0{
		return nil
	}
	logrus.Fatal(rsp)
	return nil
}

func  ComputeKey(appName string,serviceName string) string {
	target := fmt.Sprintf("/etcdv3://%s/grpc/%s", appName, serviceName)
	return target
}

func NewClientV3(endPoints []string, timeOut time.Duration) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: timeOut*time.Second,
	})
	if err != nil {
		panic(err)
	}
	return cli
}

func NewServiceRegister(client *clientv3.Client, appName string) *ServiceRegister {
	serviceRegister := &ServiceRegister{
		Client:         client,
		AppName:        appName,
		ServiceName:    "",
		ServiceAddress: "",
		Ttl:            0,
		LeaseId:        0,
		KeepAliveCh:    make(<-chan *clientv3.LeaseKeepAliveResponse),
		CloseCh: make(<-chan struct{}),
	}
	return serviceRegister
}


