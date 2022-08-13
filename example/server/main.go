package main

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	zookeeper_demo "github.com/zstone12/zookeeper-demo"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go func() {
		defer wg.Done()
		addr := "127.0.0.1:8888"
		r, err := zookeeper_demo.NewZookeeperRegistry([]string{"127.0.0.1:2181"}, 40*time.Second)
		if err != nil {
			panic(err)
		}
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "hertz.test.demo",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}))
		h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
		})
		h.Spin()
	}()
	wg.Wait()
}
