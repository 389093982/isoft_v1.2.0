package test

import (
	"bufio"
	"fmt"
	net "net"
	"os"
	"strings"
)

//客户端
//向8888端口发送消息
func TestClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("客户端dial 失败，err=%v\n", err)
		return
	}

	for true {
		reader := bufio.NewReader(os.Stdin) //终端输入
		fmt.Printf("断点1\n")
		line, err := reader.ReadString('\n')
		fmt.Printf("断点2\n")
		if err != nil {
			fmt.Printf("ReadString err:%v\n", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Printf("客户端exit")
			break
		}

		//将line发给服务端
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Printf("发送失败，err=%v\n", err)
		}
	}
}

//服务端
//监听8888端口，并接收端口发来的消息
func TestServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("监听失败，err:%v", err)
		return
	}
	defer listen.Close()

	for true {
		fmt.Printf("等待客户端来连接。。。")
		fmt.Printf("断点1\n")
		conn, err := listen.Accept()
		fmt.Printf("断点2\n")
		if err != nil {
			fmt.Printf("接受数据失败，err:%v", err)
			return
		} else {
			fmt.Printf("Accept() success conn=%v, ip=%v\n", conn, conn.RemoteAddr().String())
		}

		//准备一个协程，为客户端做准备
		go process(conn)
	}
}

//接收客户端消息并打印
func process(conn net.Conn) {
	defer conn.Close()
	//客户端通过conn 发送消息
	fmt.Printf("客户端地址%v,正在发来消息\n", conn.RemoteAddr())
	for true {
		//创建一个新的切片
		buf := make([]byte, 1024)
		fmt.Printf("断点3\n")
		n, err := conn.Read(buf)
		fmt.Printf("断点4\n")
		if err != nil {
			fmt.Printf("客户端退出了，err:%v", err)
			return
		}
		fmt.Printf("=====> %v\n", string(buf[:n]))
	}
}
