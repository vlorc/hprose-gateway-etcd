package etcd

import (
	"context"
	"github.com/vlorc/hprose-gateway-types"
	"github.com/vlorc/hprose-gateway-etcd/client"
	"github.com/vlorc/hprose-gateway-etcd/resolver"
	"github.com/vlorc/hprose-gateway-etcd/watcher"
	"github.com/vlorc/hprose-gateway-etcd/manager"
	"testing"
	"time"
)

var cli = client.NewLazyClient(client.NewClient("localhost:2379"))
var ctx, cancel = context.WithCancel(context.Background())

func Test_Resolver(t *testing.T) {
	res := resolver.NewResolver(cli, ctx, "rpc")
	go res.Watch("*", watcher.NewPrintWatcher(t.Logf))
}

func Test_Manage(t *testing.T) {
	time.Sleep(time.Second * 5)
	man := manager.NewManager(cli, context.Background(), "rpc", 5)
	user := man.Register("user", "1")
	user.Update(&types.Service{
		Id:       "1",
		Name:     "user",
		Version:  "1.0.0",
		Url:      "http://localhost:8080",
		Platform: "1",
		Meta: map[string]interface{}{
			"appid": 1,
			"key":   "123",
		},
	})
	time.Sleep(time.Second * 30)
	man.Close()
	time.Sleep(time.Second * 5)
	cancel()
}
