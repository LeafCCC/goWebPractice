package main

import (
	"fmt"
	"log"
	"net/rpc"
	rpc_objects "webPractice/rpc"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpc_objects.Args{7, 8}
	var reply int

	//此处调用服务器处的乘方法
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
