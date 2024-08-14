package server

import (
	"log"
	"net"
)

func RunTCP(handle func(net.Conn)) error {
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Println("listening on port 10000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go func() {
			defer conn.Close()
			handle(conn)
		}()
	}
}

func RunUDP(handle func(net.PacketConn)) error {
	ln, err := net.ListenPacket("udp", ":10000")
	if err != nil {
		return err
	}
	defer ln.Close()

	log.Println("listening on port 10000")

	for {
		handle(ln)
	}
}
