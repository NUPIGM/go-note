package main

import (
	"fmt"
	"net"
)

func tcp() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8080")
	listen, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			return
		}
		go handle(conn)
	}

}

func handle(conn *net.TCPConn) {
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		_, _ = conn.Write(b)
		if err != nil {
			break
		}
		fmt.Println(conn.RemoteAddr(), "connect", string(b[0:n]))

	}
}
