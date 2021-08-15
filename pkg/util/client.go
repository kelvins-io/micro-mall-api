package util

import (
	"context"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"gitee.com/kelvins-io/kelvins/util/client_conn"
	"gitee.com/kelvins-io/kelvins/util/middleware"
	"google.golang.org/grpc"
)

func GetGrpcClient(ctx context.Context, serverName string) (*grpc.ClientConn, error) {
	client, err := client_conn.NewConnClient(serverName)
	if err != nil {
		return nil, err
	}
	conf := &setting.RPCAuthSettingS{
		Token:             "c9VW6ForlmzdeDkZE2i8",
		TransportSecurity: false,
	}
	opts := middleware.GetRPCAuthDialOptions(conf)
	return client.GetConn(ctx, opts...)
}

func GetHttpEndpoints(serverName string) ([]string, error) {
	client, err := client_conn.NewConnClient(serverName)
	if err != nil {
		return nil, err
	}
	return client.GetEndpoints(context.Background())
}
