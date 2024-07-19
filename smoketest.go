package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal("listen: ", err.Error())
	}

	log.Println("listening on port 10000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("accept: ", err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatal("copy: ", err.Error())
	}
}
