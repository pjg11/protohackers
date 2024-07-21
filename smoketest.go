package main

import (
	"io"
	"log"
	"net"

	"protohackers/server"
)

func echo(conn net.Conn) {
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatal("copy: ", err.Error())
	}
}

func main() {
	server.Run(echo)
}
