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
	"fmt"
	"log"
	"net/http"

	pb "test/plugins/scenarios/grpc/api"

	_ "github.com/apache/skywalking-go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999")
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	defer conn.Close()
	client := pb.NewSendMsgClient(conn)

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("ok"))
	})

	http.HandleFunc("/consumer", func(writer http.ResponseWriter, request *http.Request) {
		resp, _ := client.SendMsg(context.Background(), &pb.Request{requestMsg: "massages"})
		fmt.Printf(resp.GetResponseMsg())
	})

	_ = http.ListenAndServe(":8080", nil)
}
