package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
	rpc_objects "webPractice/rpc"
)

func main() {
	worker := new(rpc_objects.Args)

	//注册到默认的服务中心
	rpc.Register(worker)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Starting rpc listener failed:", err)
	}
	go http.Serve(listener, nil)
	time.Sleep(time.Second * 1000)
}
