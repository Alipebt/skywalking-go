package grpc

import "strings"

func formatOperationName(service string, method string) string {
	service = service[1:]
	service = strings.ReplaceAll(service, "/", ".")
	operationName := service + method

	return operationName
}
