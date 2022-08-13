package main

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	zookeeper_demo "github.com/zstone12/zookeeper-demo"
)

var wg sync.WaitGroup

func main() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	r, err := zookeeper_demo.NewZookeeperResolver([]string{"127.0.0.1:2181"}, 40*time.Second)
	cli.Use(sd.Discovery(r))
	for i := 0; i < 10; i++ {
		status, body, err := cli.Get(context.Background(), nil, "http://hertz.test.demo/ping", config.WithSD(true))
		if err != nil {
			hlog.Fatal(err)
		}
		hlog.Infof("code=%d,body=%s", status, string(body))
	}
}
