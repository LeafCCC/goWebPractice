package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var userRecord map[string]int

func main() {
	userRecord = make(map[string]int)

	fmt.Println("Starting the server ...")
	// 创建 listener 选择tcp协议 并设置好监听端口
	// 可选的类型为 "tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listener creating,", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error listener accepting,", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		l, err := conn.Read(buf)

		rec := string(buf[:l])

		recSplit := strings.Split(rec, " ")

		//提取姓名与发送的内容
		name, content := recSplit[1], recSplit[2]
		switch content {
		case "SH":
			fmt.Println("Server is shut down by the client", name[:len(name)-1], "...")
			os.Exit(0)
		case "Q":
			fmt.Println("User", name, "leaves...")
			userRecord[name] = 0
		case "WHO":
			printUsers()
		default:
			fmt.Printf("Received data: %v \n", rec)
			userRecord[name] = 1
		}

		if err != nil {
			fmt.Println("Error reading the data from the client,", err.Error())
			return //终止程序
		}

	}
}

func printUsers() {
	fmt.Println("This is the client list: 1:active, 0=inactive")
	for user, state := range userRecord {
		fmt.Printf("User %s is %d \n", user, state)
	}
}
