package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 连接服务器，IP地址本地，端口8888
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	for {
		// 向服务器发送数据
		_, err = conn.Write([]byte("我是客户端小白菜"))
		if err != nil {
			fmt.Println("发送数据失败", err)
			return
		}
		msg := make([]byte, 512)
		// 接受服务器消息
		cnt, err := conn.Read(msg)
		if err != nil {
			fmt.Println("接受数据失败", err)
			return
		}
		fmt.Println("server", string(msg[:cnt]))

		time.Sleep(time.Second)
	}
}
