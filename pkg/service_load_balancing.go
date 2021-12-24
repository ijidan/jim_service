package pkg

import (
	"google.golang.org/grpc/resolver"
	"time"
)

const (
	jScheme = "jim_service"
)

type JResolver struct {
	target           resolver.Target
	cc               resolver.ClientConn
	addressStore     map[string][]string
	serviceDiscovery *ServiceDiscovery
}

func (r *JResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (r *JResolver) Close()                                  {}

func (r *JResolver) UpdateAddressStore() {
	r.addressStore = r.serviceDiscovery.GetServiceServerList()
}
func (r *JResolver) Start() {
	addrMap := r.addressStore[r.target.Endpoint]
	addrList := make([]resolver.Address, len(addrMap))
	for i, s := range addrMap {
		addrList[i] = resolver.Address{Addr: s}
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
	serviceDiscovery *ServiceDiscovery
}

func (b *JResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &JResolver{
		target:           target,
		cc:               cc,
		addressStore:     map[string][]string{},
		serviceDiscovery: b.serviceDiscovery,
	}
	r.Start()
	return r, nil
}
func (b *JResolverBuilder) Scheme() string { return jScheme }

func NewJResolverBuilder(serviceDiscovery *ServiceDiscovery) *JResolverBuilder {
	builder := &JResolverBuilder{serviceDiscovery: serviceDiscovery}
	return builder
}

func LoadBalancingRegister(builder *JResolverBuilder) {
	resolver.Register(builder)
}
