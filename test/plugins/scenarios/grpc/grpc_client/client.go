/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"log"
	"net/http"

	pb "test/plugins/scenarios/grpc/api"

	_ "github.com/apache/skywalking-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})

	http.HandleFunc("/consumer", func(writer http.ResponseWriter, request *http.Request) {
		resp, err := client.SimpleRPC(context.Background(), &pb.HelloRequest{Name: "simpleRPC"})
		if err != nil {
			writer.WriteHeader(500)
			_, _ = writer.Write([]byte(err.Error()))
			log.Println(err)
		}

		log.Println(resp.GetReply())

		_, _ = writer.Write([]byte(resp.String()))
	})

	_ = http.ListenAndServe(":8080", nil)
}
