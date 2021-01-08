package util

import (
	"context"
	"gitee.com/kelvins-io/kelvins/util/client_conn"
	"google.golang.org/grpc"
)


func GetGrpcClient(serverName string) (*grpc.ClientConn, error) {
	client, err := client_conn.NewConn(serverName)
	if err != nil {
		return nil, err
	}
	return client.GetConn(context.Background())
}
