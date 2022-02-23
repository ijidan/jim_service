package interceptor

type LimiterInterceptor struct{}

func (i *LimiterInterceptor) Limit() bool {
	return false
}

func NewLimiterInterceptor() *LimiterInterceptor  {
	return &LimiterInterceptor{}
}