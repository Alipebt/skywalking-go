package grpc

import (
	"context"

	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
	"google.golang.org/grpc/internal/transport"
	"google.golang.org/grpc/metadata"
)

type ServerStartInterceptor struct {
}

func (h *ServerStartInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	// 获取参数
	ctx := invocation.Args()[0].(context.Context)
	stream := invocation.Args()[1].(*transport.Stream)
	method := stream.Method()

	md, _ := metadata.FromOutgoingContext(ctx)
	
	path := ""
	vals := md.Get(":path")
	if len(vals) > 0 {
		path = vals[0]
	}
	// 创建span
	s, err := tracing.CreateEntrySpan(method, func(headerKey string) (string, error) {
		// 获取header的值
		Value := ""
		vals := md.Get(headerKey)
		if len(vals) > 0 {
			Value = vals[0]
		}

		return Value, nil
	}, tracing.WithLayer(tracing.SpanLayerRPCFramework),
		tracing.WithTag(tracing.TagURL, path),
		tracing.WithComponent(5016),
	)

	if err != nil {
		return err
	}

	invocation.SetContext(s)
	return nil
}

func (h *ServerStartInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {
	if invocation.GetContext() == nil {
		return nil
	}

	invocation.GetContext().(tracing.Span).End()
	return nil
}
