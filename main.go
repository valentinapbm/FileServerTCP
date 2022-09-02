package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		fmt.Println("unable to start server:", err)
	}
	defer listener.Close()
	log.Println("server started on :8888")

	//endless loop for acepting incoming connection
	for{
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("unable to accept connection:", err)
			continue
		}

		c := s.newClient(conn)
		go c.readInput()
	}
	
}