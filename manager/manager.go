package manager

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/vlorc/hprose-gateway-types"
	"github.com/vlorc/hprose-gateway-etcd/client"
	"sync"
)

type etcdManager struct {
	client  func() *clientv3.Client
	leaseId func() clientv3.LeaseID
	backend context.Context
	cancel  context.CancelFunc
	scheme  string
	pool    sync.Map
}

func NewManager(cli func() *clientv3.Client, parent context.Context, scheme string, ttl int64) types.NamedManger {
	ctx, cancel := context.WithCancel(parent)
	return &etcdManager{
		client:  cli,
		backend: ctx,
		cancel:  cancel,
		scheme:  scheme,
		leaseId: client.NewLazyLease(client.Grant(cli, ctx, ttl)),
	}
}

func (m *etcdManager) Register(name, uuid string) types.NamedRegister {
	return m.register(m.formatKey(name, uuid))
}

func (m *etcdManager) formatKey(name, uuid string) string {
	if "" != uuid {
		return "/" + m.scheme + "/" + name + "/" + uuid
	}
	return "/" + m.scheme + "/" + name
}

func (m *etcdManager) formatValue(data interface{}) string {
	value, _ := json.MarshalIndent(data, "", "    ")
	return string(value)
}

func (m *etcdManager) register(key string) types.NamedRegister {
	return &etcdRegiser{
		manager: m,
		key:     key,
	}
}

func (m *etcdManager) update(key string, val interface{}) error {
	value := m.formatValue(val)
	if _, err := m.client().Put(m.backend, key, value, clientv3.WithLease(m.leaseId())); err != nil {
		return err
	}
	m.pool.Store(key, true)
	return nil
}

func (m *etcdManager) remove(key string) error {
	if _, err := m.client().Delete(m.backend, key); err != nil {
		return err
	}
	m.pool.Delete(key)
	return nil
}

func (m *etcdManager) Close() error {
	_, err := m.client().Revoke(m.backend, m.leaseId())
	return err
}

func (m *etcdManager) Keys() (result []string) {
	m.pool.Range(func(key, _ interface{}) bool {
		result = append(result, key.(string))
		return true
	})
	return result
}
