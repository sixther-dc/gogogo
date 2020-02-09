package routine

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//RunClock fuc
func RunClock() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("时间服务器正在运行...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerConn(conn)
	}
}

func handlerConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Fprintf(c, "%s: %s\n", input.Text(), time.Now().Format("15:04:04"))
	}
	//记录协议
	log.Printf("%s\n", c.RemoteAddr().Network())
	//记录客户端地址
	log.Printf("%s\n", c.RemoteAddr().String())
}
