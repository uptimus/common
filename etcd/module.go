package etcd

import (
	"go.etcd.io/etcd/clientv3"
)

type EtcdService struct {
	Client *clientv3.Client
}

func NewEtcdService(config *Config) (*EtcdService, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, err
	}

	return &EtcdService{client}, nil
}
