package etcd

import (
	"context"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClient interface {
	PUT(key, value string) error
	Get(key string) ([]GetResponse, error)
	Close() error
}

func New(endpoints []string) (EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		cli: cli,
	}, nil
}

type Client struct {
	cli *clientv3.Client
}

func (c *Client) PUT(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := c.cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			return ErrCtxCanceled
		case context.DeadlineExceeded:
			return ErrCtxDeadlineExceeded
		case rpctypes.ErrEmptyKey:
			return ErrEmptyKey
		default:
			return ErrDefault
		}
	}
	return nil
}

type GetResponse struct {
	Key, Value []byte
}

func (c *Client) Get(key string) ([]GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	r, err := c.cli.Get(ctx, key)
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			return nil, ErrCtxCanceled
		case rpctypes.ErrEmptyKey:
			return nil, ErrEmptyKey
		default:
			return nil, ErrDefault
		}
	}

	res := make([]GetResponse, 0)
	for _, v := range r.Kvs {
		res = append(res, GetResponse{
			Key:   v.Key,
			Value: v.Value,
		})
	}
	return res, nil
}

func (cc *Client) Close() error {
	return cc.cli.Close()
}
