package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

//tcp server端

func processConn(conn net.Conn) {
	defer conn.Close()
	var tmp [1024]byte
	for {
		//3. 与客户端通信
		n, err := conn.Read(tmp[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from conn failed, err", err)
			return
		}
		fmt.Println(string(tmp[:n]))

		//向客户端发送回复
		content := string(tmp[:n])
		content_sli := strings.Split(content, "+")
		if len(content_sli) != 2 {
			conn.Write([]byte("0"))
			continue
		}

		a, err := strconv.Atoi(content_sli[0])
		if err != nil {
			conn.Write([]byte("0"))
			continue
		}
		b, err := strconv.Atoi(content_sli[1])
		if err != nil {
			conn.Write([]byte("0"))
			continue
		}
		conn.Write([]byte(strconv.Itoa(a + b)))
	}
}

func setNoDelay(conn net.Conn) error {
	switch conn := conn.(type) {
	case *net.TCPConn:
		var err error
		if err = conn.SetNoDelay(false); err != nil {
			return err
		}
		return err

	default:
		return fmt.Errorf("unknown connection type %T", conn)
	}
}

func main() {
	network := "tcp"
	address := "127.0.0.1:30001"
	listener, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("start tcp at 127.0.0.1:30001 failed, err:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		//设置服务端禁用TCP_NODELAY，打开Nagle算法
		setNoDelay(conn)
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go processConn(conn)
	}

}
