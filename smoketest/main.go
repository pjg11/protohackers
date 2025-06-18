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
		log.Fatal(err)
	}
}

func main() {
	err := server.RunTCP(echo)
	if err != nil {
		log.Fatal(err)
	}
}
