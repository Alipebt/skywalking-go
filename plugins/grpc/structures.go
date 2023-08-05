package grpc

import (
	"context"
)

//skywalking:native google.golang.org/grpc/internal/transport Stream
type nativeStream struct {
	ctx    context.Context
	method string
}

func (s *nativeStream) Method() string {
	return s.method
}

func (s *nativeStream) Context() context.Context {
	return s.ctx
}

//skywalking:native google.golang.org/grpc/internal/transport Stream
type nativeCallHdr struct {
	Method string
}

//skywalking:native google.golang.org/grpc ClientConn
type nativeClientConn struct {
	ctx context.Context
}

func (cc *nativeClientConn) Target() string {
	return ""
}

//skywalking:native google.golang.org/grpc clientStream
type nativeclientStream struct {
	cc      *nativeClientConn
	ctx     context.Context
	callHdr *nativeCallHdr
}

func (cs *nativeclientStream) Context() context.Context {
	return cs.ctx
}

//skywalking:native google.golang.org/grpc serverStream
type nativeserverStream struct {
	ctx context.Context
	s   *nativeStream
}

func (cs *nativeserverStream) Context() context.Context {
	return cs.ctx
}
