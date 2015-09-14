package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	port := flag.String("p", "2195", "port to listen on")
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", *port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening on port %s\n", *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("connection accepted")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		header := make([]byte, 5)
		_, err := conn.Read(header)
		if err != nil {
			conn.Close()
			return
		}

		fmt.Println("header:", header)

		body := make([]byte, header[4])
		_, err = conn.Read(body)
		if err != nil {
			conn.Close()
			return
		}
		fmt.Println("body:", string(body))
	}
}
