package grpc

import (
	"github.com/apache/skywalking-go/plugins/core/operator"
	// "github.com/apache/skywalking-go/plugins/core/tracing"
	// "google.golang.org/grpc"
)

type GetPeerInterceptor struct {
}

func (h *GetPeerInterceptor) BeforeInvoke(invocation operator.Invocation) error {

	// span := tracing.ActiveSpan()
	// 获取实例
	// a := invocation.CallerInstance().(*csAttempt)
	// csAttempt中的t接口字段有获取peer的方法
	// remotePeer := a.t.RemoteAddr()
	// 设置peer
	// span.SetPeer(remotePeer)

	return nil
}
