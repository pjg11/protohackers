package server

import (
	"net"
	"log"
)

func Run(handle func(net.Conn)) (error) {
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("listening on port 10000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			defer conn.Close()
			handle(conn)
		}()
	}
}
