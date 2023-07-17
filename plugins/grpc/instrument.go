// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package grpc

import (
	"embed"
	"strings"

	"github.com/apache/skywalking-go/plugins/core/instrument"
)

//go:embed *
var fs embed.FS

//skywalking:nocopy
type Instrument struct {
}

func NewInstrument() *Instrument {
	return &Instrument{}
}

func (i *Instrument) Name() string {
	return "grpc"
}

func (i *Instrument) BasePackage() string {
	return "google.golang.org/grpc"
}

func (i *Instrument) VersionChecker(version string) bool {
	return strings.HasPrefix(version, "v1.")
}

func (i *Instrument) Points() []*instrument.Point {
	return []*instrument.Point{
		{
			PackagePath: "client",
			At: instrument.NewMethodEnhance("*ClientConn", "Invoke",
				instrument.WithArgType(0, "context.Context"),
				instrument.WithArgType(1, "string"),
				instrument.WithResultCount(1),
				instrument.WithResultType(0, "error")),
			Interceptor: "ClientUnaryInterceptor",
		},
		{
			PackagePath: "server",
			At: instrument.NewMethodEnhance("*Server", "handleStream",
				instrument.WithArgsCount(3),
				instrument.WithArgType(0, "transport.ServerTransport"),
				instrument.WithArgType(1, "*transport.Stream"),
				instrument.WithArgType(2, "*traceInfo")),
			Interceptor: "ServerStartInterceptor",
		},
		{
			PackagePath: "server",
			At: instrument.NewMethodEnhance("*Server", "processUnaryRPC",
				instrument.WithArgsCount(5),
				instrument.WithArgType(0, "transport.ServerTransport"),
				instrument.WithArgType(1, "*transport.Stream"),
				instrument.WithArgType(2, "*serviceInfo"),
				instrument.WithArgType(3, "*MethodDesc"),
				instrument.WithArgType(4, "*traceInfo"),
				instrument.WithResultCount(1),
				instrument.WithResultType(0, "error")),
			Interceptor: "ServerUnaryInterceptor",
		},
	}
}

func (i *Instrument) FS() *embed.FS {
	return &fs
}
