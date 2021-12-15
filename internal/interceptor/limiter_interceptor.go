package interceptor

type LimiterInterceptor struct{}

func (i *LimiterInterceptor) Limit() bool {
	return true
}

func NewLimiterInterceptor() *LimiterInterceptor  {
	return &LimiterInterceptor{}
}