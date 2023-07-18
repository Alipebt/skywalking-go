package grpc

import (
	"context"

	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type ClientUnaryInterceptor struct {
}

func (h *ClientUnaryInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	// 获取参数
	ctx := invocation.Args()[0].(context.Context)
	method := invocation.Args()[1].(string)

	// 获取路径
	md, _ := metadata.FromOutgoingContext(ctx)
	path := ""
	vals := md.Get(":path")
	if len(vals) > 0 {
		path = vals[0]
	}

	// 获取peer（）
	remotePeer, _ := getRemotePeer(ctx)

	// 创建span
	s, err := tracing.CreateExitSpan(method, remotePeer, func(headerKey, headerValue string) error {
		// 将新的元数据附加到上下文中
		ctx = context.WithValue(ctx, headerKey, headerValue)
		return nil
	},
		tracing.WithLayer(tracing.SpanLayerRPCFramework),
		tracing.WithTag(tracing.TagURL, path),
		tracing.WithComponent(5016),
	)

	if err != nil {
		return err
	}

	invocation.SetContext(s)
	return nil
}

func (h *ClientUnaryInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	invocation.GetContext()
	span := invocation.GetContext().(tracing.Span)

	if err, ok := result[0].(error); ok && err != nil {
		span.Error(err.Error())
	}

	span.End()
	return nil
}

func getRemotePeer(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		// handle error
	}
	addr := pr.Addr.String()
	if !ok {
		// handle error
	}
	return addr, nil
}
