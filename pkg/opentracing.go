package pkg

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func AddSpan(c context.Context,spanName string,funcName string,req interface{},rsp interface{})  {
	span, _ := opentracing.StartSpanFromContext(c, spanName)
	defer func() {
		span.SetTag("func",funcName)
		span.SetTag("request", req)
		span.SetTag("reply", rsp)
		span.Finish()
	}()
}