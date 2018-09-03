package manager

import "github.com/vlorc/hprose-gateway-types"

type etcdRegiser struct {
	manager *etcdManager
	key     string
}

func (r *etcdRegiser) Update(service *types.Service) error {
	return r.manager.update(r.key, service)
}

func (r *etcdRegiser) Close() error {
	return r.manager.remove(r.key)
}
