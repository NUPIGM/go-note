package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "gpapi.biligame.net:443")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	// reader := bufio.NewReader(os.Stdin)
	bt := make([]byte, 1024)
	// for {
	// byt, _, _ := reader.ReadLine()
	// conn.Write(byt)
	n, _ := conn.Read(bt)
	fmt.Println("接收到的返回:", string(bt[0:n]))
	// }

}
