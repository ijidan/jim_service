package service

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ServiceRegister struct {
	Client              *clientv3.Client
	AppName             string
	Service             BasicService
	ServiceLeaseIdMap   map[BasicService]clientv3.LeaseID
	ServiceKeepAliveMap map[BasicService]<-chan *clientv3.LeaseKeepAliveResponse
	ServiceCloseMap     map[BasicService]chan bool
}

func (r *ServiceRegister) RegisterService(srvList ...BasicService) {
	for _, srv := range srvList {
		LeaseRsp, err1 := r.Client.Grant(context.Background(), srv.GetTTL())
		if err1 != nil {
			panic(err1)
		}
		keepAliveCh, err2 := r.Client.KeepAlive(context.Background(), LeaseRsp.ID)
		if err2 != nil {
			panic(err2)
		}
		_, err3 := r.Client.KV.Put(context.Background(), srv.GetName(), srv.GetAddress(), clientv3.WithLease(LeaseRsp.ID))
		if err3 != nil {
			panic(err3)
		}
		r.ServiceLeaseIdMap[srv] = LeaseRsp.ID
		r.ServiceKeepAliveMap[srv] = keepAliveCh

	}
	for _, srv := range srvList {
		go r.CloseSrv(srv)
	}
	go r.KeepAlive()
}

func (r *ServiceRegister) UnRegisterService(srv BasicService) {
	leaseId := r.ServiceLeaseIdMap[srv]
	var err error
	_, err = r.Client.Revoke(context.Background(), leaseId)
	if err != nil {
		return
	}
	_, err = r.Client.KV.Delete(context.Background(), srv.GetName())
	if err != nil {
		panic(err)
	}
}

func (r *ServiceRegister) Stop(srv BasicService) {
	closeCh := r.ServiceCloseMap[srv]
	closeCh <- true
}

func (r *ServiceRegister) CloseSrv(srv BasicService) {
	closeCh := r.ServiceCloseMap[srv]
	for {
		select {
		case res := <-closeCh:
			if res {
				r.UnRegisterService(srv)
			}
		}
	}
}

func (r *ServiceRegister) KeepAlive() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			for _, _ch := range r.ServiceKeepAliveMap {
				_=<-_ch
			}
		}
	}
}

func NewServiceRegister(client *clientv3.Client, appName string) *ServiceRegister {
	serviceRegister := &ServiceRegister{
		Client:              client,
		AppName:             appName,
		ServiceLeaseIdMap:   make(map[BasicService]clientv3.LeaseID),
		ServiceKeepAliveMap: make(map[BasicService]<-chan *clientv3.LeaseKeepAliveResponse),
		ServiceCloseMap:     make(map[BasicService]chan bool),
	}
	return serviceRegister
}
