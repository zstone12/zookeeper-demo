## how to use?
1. get zookeeper
```shell
make 
```
2. run zookeeper
```shell
cd zookeeper/bin/
./zkServer.sh start
```
3. run demo
```shell
go run example/server/main.go
go run example/client/main.go
```