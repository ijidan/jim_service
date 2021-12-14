package call

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	status2 "google.golang.org/grpc/status"
	"io"
	"jim/global"
	"sync"
	"time"
)

// BasicCall 基本调用类
type BasicCall struct {
	ctx    context.Context
	conn   *grpc.ClientConn
	tracer opentracing.Tracer
	closer io.Closer
	span   opentracing.Span
}

//构建ctx
func (c *BasicCall) buildCtx() {
	//上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	c.ctx = ctx
}

//构建链接
func (c *BasicCall) buildConnection() {
	address := fmt.Sprintf("%s:%d", global.Config.Rpc.Host, global.Config.Rpc.Port)
	conn := buildConnection(address)
	//defer conn.Close()
	c.conn = conn
}

// CloseConn 关闭连接
func (c *BasicCall) CloseConn() {
	_ = c.conn.Close()
}

//结束时操作
func (c *BasicCall) endHandler() {
	//c.closeConn()
}

// GetConn 获取连接
func (c *BasicCall) GetConn() *grpc.ClientConn {
	return c.conn
}

// GetCtx 获取上下文
func (c *BasicCall) GetCtx() context.Context {
	return c.ctx
}

// Close 关闭
func (c *BasicCall) Close() {
	_ = c.conn.Close()
}

// CheckTimeout 检测超时
func (c *BasicCall) CheckTimeout(err error) {
	c.doCheckErrorCode(err, codes.DeadlineExceeded, "gRpc timeout!")
}

// CheckCancel 检测取消
func (c *BasicCall) CheckCancel(err error) {
	c.doCheckErrorCode(err, codes.Canceled, "gRpc canceled")
}

// CheckError 检测错误
func (c *BasicCall) CheckError(err error) {
	if status, ok := status2.FromError(err); ok {
		message := status.Message()
		log.Error(message)
	}
}

//检测错误
func (c *BasicCall) doCheckErrorCode(err error, code codes.Code, message string) {
	if status, ok := status2.FromError(err); ok {
		statusCode := status.Code()
		if statusCode == code {
			log.Error(message)
		}
	}
}

var connMap = map[string]*grpc.ClientConn{}
var lock sync.Mutex

func buildConnection(address string) *grpc.ClientConn {
	lock.Lock()
	defer lock.Unlock()
	conn := connMap[address]
	if conn == nil {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			grpclog.Fatalln("gRPC connection error：" + err.Error())
		}
		connMap[address] = conn
	}
	return connMap[address]
}
