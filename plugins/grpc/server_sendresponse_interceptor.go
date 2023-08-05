package grpc

import (
	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
)

type ServerSendResponseInterceptor struct {
}

func (h *ServerSendResponseInterceptor) BeforeInvoke(invocation operator.Invocation) error {

	cs := invocation.Args()[1].(*nativeStream)
	method := cs.Method()

	s, err := tracing.CreateLocalSpan(formatOperationName(method, "/server/Response/sendResponse"),
		tracing.WithLayer(tracing.SpanLayerRPCFramework),
		tracing.WithTag(tracing.TagURL, method),
		tracing.WithComponent(23),
	)

	if err != nil {
		return err
	}

	invocation.SetContext(s)

	return nil
}

func (h *ServerSendResponseInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {

	if invocation.GetContext() != nil {

		span := invocation.GetContext().(tracing.Span)

		if err, ok := result[0].(error); ok && err != nil {
			span.Error(err.Error())
		}

		span.End()
	}

	return nil
}
