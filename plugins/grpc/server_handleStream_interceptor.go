package grpc

import (
	"google.golang.org/grpc/metadata"

	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
)

type ServerHandleStreamInterceptor struct {
}

func (h *ServerHandleStreamInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	stream := invocation.Args()[1].(*nativeStream)
	method := stream.Method()
	ctx := stream.Context()
	md, _ := metadata.FromIncomingContext(ctx)

	s, err := tracing.CreateEntrySpan(formatOperationName(method, ""), func(headerKey string) (string, error) {

		Value := ""
		vals := md.Get(headerKey)
		if len(vals) > 0 {
			Value = vals[0]
		}

		return Value, nil
	}, tracing.WithLayer(tracing.SpanLayerRPCFramework),
		tracing.WithTag(tracing.TagURL, method),
		tracing.WithComponent(23),
	)

	if err != nil {
		return err
	}

	invocation.SetContext(s)
	return nil
}

func (h *ServerHandleStreamInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {

	if invocation.GetContext() == nil {
		return nil
	}

	invocation.GetContext().(tracing.Span).End()
	return nil
}
