package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	target := "localhost:30001"
	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Fatal(err)
	}
	// 和服务端建立连接
	conn, err := net.DialTCP("tcp", nil, raddr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	//conn.SetNoDelay(false) // 如果打开这行代码，则客户端禁用TCP_NODELAY，打开Nagle算法
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请输入要计算的数据，例如1+2：")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
		}
		// 发送数据
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			break
		}
		fmt.Println(string(buf[0:length]))

	}
}
