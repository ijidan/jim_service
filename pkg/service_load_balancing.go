package pkg

import "google.golang.org/grpc/resolver"

const (
	jScheme      = "jim_service"
)

type JResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addressStore map[string][]string
}

func (*JResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*JResolver) Close()  {}
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
}

type JResolverBuilder struct{
	serviceName string
	addrList []string
}

func (b *JResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &JResolver{
		target: target,
		cc:     cc,
		addressStore: map[string][]string{
			b.serviceName: b.addrList,
		},
	}
	r.Start()
	return r, nil
}
func (b *JResolverBuilder) Scheme() string { return jScheme }

func (b *JResolverBuilder) Register(serviceName string,addrList []string)  {
	b.serviceName=serviceName
	b.addrList=addrList
}

func NewJResolverBuilder(serviceName string,addrList []string) *JResolverBuilder  {
	builder:=&JResolverBuilder{}
	builder.Register(serviceName,addrList)
	return builder
}

func LoadBalancingRegister(builder *JResolverBuilder)  {
	resolver.Register(builder)
}
