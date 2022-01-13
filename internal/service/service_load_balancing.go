package service

import (
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"strings"
	"time"
)

const (
	JScheme = "jim"
)

type JResolver struct {
	target           resolver.Target
	cc               resolver.ClientConn
	addressStore     map[string][]string
	serviceDiscovery *Discovery
}

func (r *JResolver) ResolveNow(o resolver.ResolveNowOptions) {
	logrus.Println("resolve now")
}
func (r *JResolver) Close()                                  {
	logrus.Println("resolver close")
}

func (r *JResolver) UpdateAddressStore() {
	r.addressStore = r.serviceDiscovery.GetServiceServerList()
}

func (r *JResolver) Start() {
	key:=strings.Trim(r.target.URL.Path,"/")
	addrMap := r.addressStore[key]
	addrList := make([]resolver.Address, len(addrMap))
	for i, v := range addrMap {
		addrList[i] = resolver.Address{Addr: v}
	}
	err := r.cc.UpdateState(resolver.State{Addresses: addrList})
	if err != nil {
		panic(err)
	}
	go r.Watch()
}

func (r *JResolver) Watch() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			r.UpdateAddressStore()
		}
	}
}

type JResolverBuilder struct {
	serviceDiscovery *Discovery
}

func (b *JResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &JResolver{
		target:           target,
		cc:               cc,
		addressStore:     b.serviceDiscovery.GetServiceServerList(),
		serviceDiscovery: b.serviceDiscovery,
	}
	r.Start()
	return r, nil
}
func (b *JResolverBuilder) Scheme() string { return JScheme }

func NewJResolverBuilder(client *clientv3.Client) *JResolverBuilder {
	serviceDiscovery:=NewServiceDiscovery(client)
	builder := &JResolverBuilder{serviceDiscovery: serviceDiscovery}
	return builder
}


