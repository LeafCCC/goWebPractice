package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	//打开连接:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, err2 := inputReader.ReadString('\n')

	if err2 != nil {
		fmt.Println("Error namel reading,", err2.Error())
		return
	}

	// fmt.Printf("CLIENTNAME %s", clientName)
	//移除掉CRLF linux下是LF
	trimmedClient := strings.Trim(clientName, "\r\n") // Windows 平台下用 "\r\n"，Linux平台下使用 "\n"
	// 给服务器发送信息直到程序退出：
	for {
		fmt.Println("What to send to the server? Type Q to quit, and SH to shut down the server.")
		input, err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Println("Error content reading,", err2.Error())
			return
		}

		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--s%--", input)
		// fmt.Printf("trimmedInput:--s%--", trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte("From " + trimmedClient + ": " + trimmedInput))

		if err != nil {
			fmt.Println("Error sending messages,", err2.Error())
			return
		}
	}
}
