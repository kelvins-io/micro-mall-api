package util

import (
	"context"
	"gitee.com/kelvins-io/kelvins/util/client_conn"
	"google.golang.org/grpc"
)

func GetGrpcClient(serverName string) (*grpc.ClientConn, error) {
	client, err := client_conn.NewConnClient(serverName)
	if err != nil {
		return nil, err
	}
	return client.GetConn(context.Background())
}

func GetHttpEndpoints(serverName string) ([]string, error) {
	client, err := client_conn.NewConnClient(serverName)
	if err != nil {
		return nil, err
	}
	return client.GetEndpoints(context.Background())
}
