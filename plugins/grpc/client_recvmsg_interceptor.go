package grpc

import (
	"google.golang.org/grpc/peer"

	"github.com/apache/skywalking-go/plugins/core/operator"
	"github.com/apache/skywalking-go/plugins/core/tracing"
)

type ClientRecvMsgInterceptor struct {
}

func (h *ClientRecvMsgInterceptor) BeforeInvoke(invocation operator.Invocation) error {
	cs := invocation.CallerInstance().(*nativeclientStream)
	ctx := cs.Context()
	method := cs.callHdr.Method

	peerKey, ok := peer.FromContext(ctx)
	remoteAddr := ""
	if ok {
		remoteAddr = peerKey.Addr.String()
	}

	if remoteAddr == "127.0.0.1:11800" {
		return nil
	}

	s, err := tracing.CreateLocalSpan(formatOperationName(method, "/client/Response/recvMsg"),
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

func (h *ClientRecvMsgInterceptor) AfterInvoke(invocation operator.Invocation, result ...interface{}) error {

	if invocation.GetContext() != nil {

		span := invocation.GetContext().(tracing.Span)

		if err, ok := result[0].(error); ok && err != nil {
			span.Error(err.Error())
		}

		span.End()
	}

	return nil
}
