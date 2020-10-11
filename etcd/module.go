package etcd

import (
	"go.etcd.io/etcd/clientv3"

	"github.com/uptimus/common/etcd"
)

type EtcdService struct {
}

func NewEtcdService(config *etcd.Config) (*EtcdService, error) {

	_, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Endpoints,
		DialTimeout: config.DialTimeout,
	})
	if err != nil {
		return nil, err
	}

	return &EtcdService{}, nil
}
