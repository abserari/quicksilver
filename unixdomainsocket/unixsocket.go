package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var file  = "test.sock"
start:
	lis, err := net.Listen("unix", file)
	if err != nil {
		log.Println("UNIX Domain Socket 创 建失败，正在尝试重新创建 -> ", err)
		err = os.Remove(file)
		if err != nil {
			log.Fatalln("删除 sock 文件失败！程序退出 -> ", err)
		}
		goto start
	} else {
		fmt.Println("创建 UNIX Domain Socket 成功")
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("请求接收错误 -> ", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err == io.EOF { //当对端退出后会报这么一个错误
			fmt.Println("go : 对端已接 收全部数据")
			break
		} else if err != nil { //处理完客户端关闭的错误正常错误还是要处理的
			log.Println(err)
			break
		}
		fmt.Println("go : 对端发送数据 -> ", msg)
		conn.Write([]byte("服务端已接收数\n"))
	}
}