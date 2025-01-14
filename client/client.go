package client

import (
	"context"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
)

type EtcdClient struct {
	Client         *clientv3.Client
	Retries        uint64
	RequestTimeout time.Duration
	Context        context.Context
}

func (cli *EtcdClient) SetContext(ctx context.Context) *EtcdClient {
	return &EtcdClient{
		Client: cli.Client,
		Retries: cli.Retries,
		RequestTimeout: cli.RequestTimeout,
		Context: ctx,
	}
}

func (cli *EtcdClient) Close() {
	cli.Client.Close()
}

func shouldRetry(err error, retries uint64) bool {
	etcdErr, ok := err.(rpctypes.EtcdError)
	if !ok {
		return false
	}

	if etcdErr.Code() != codes.Unavailable || retries == 0 {
		return false
	}

	return true
}
