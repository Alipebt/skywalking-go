package grpc

import (
	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
)

type ServerUnaryInterceptor struct {
}

func (h *ServerUnaryInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	span := tracing.ActiveSpan()
	span.Tag("RPCType", "Unary")

	return nil
}
func (h *ServerUnaryInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {

	return nil
}
