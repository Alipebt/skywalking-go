package grpc

import (
	"context"
	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
	"google.golang.org/grpc/metadata"
)

var theMethodKey string = "theMethodKey"

type ClientUnaryInterceptor struct {
}

func (h *ClientUnaryInterceptor) BeforeInvoke(invocation operator.Invocation) error {

	// 获取参数
	ctx := invocation.Args()[0].(context.Context)
	method := invocation.Args()[1].(string)

	clientconn := invocation.CallerInstance().(*nativeClientConn)
	remoteAddr := clientconn.Target()
	if remoteAddr == "127.0.0.1:11800" {
		return nil
	}

	s, err := tracing.CreateExitSpan(formatOperationName(method, ""), remoteAddr, func(headerKey, headerValue string) error {
		ctx = metadata.AppendToOutgoingContext(ctx, headerKey, headerValue)
		invocation.ChangeArg(0, ctx)

		return nil
	},
		tracing.WithLayer(tracing.SpanLayerRPCFramework),
		tracing.WithTag(tracing.TagURL, method),
		tracing.WithTag("RPCType", "Unary"),
		tracing.WithComponent(23),
	)

	if err != nil {
		return err
	}

	invocation.SetContext(s)

	return nil
}

func (h *ClientUnaryInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {

	if invocation.GetContext() != nil {

		span := invocation.GetContext().(tracing.Span)

		if err, ok := result[0].(error); ok && err != nil {
			span.Error(err.Error())
		}

		span.End()
	}
	return nil
}
