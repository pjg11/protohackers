package main

import (
	"log"
	"net"
	"fmt"
	"strings"

	"protohackers/server"
)

var db map[string]string

func unusualdatabase(conn net.PacketConn) {
	buf := make([]byte, 1000)
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		return
	}

	key, value, found := strings.Cut(string(buf[:n]), "=")
	key = strings.Trim(key, "\n")
	if found {
		// Insert
		if key != "version" {
			db[key] = value
		}
	} else {
		// Retrieve
		resp := fmt.Sprintf("%s=%s", key, db[key])
		conn.WriteTo([]byte(resp), addr)
	}
}

func main() {
	db = map[string]string{
		"version": "Ken's Key-Value Store 1.0",
	}
	err := server.RunUDP(unusualdatabase)
	if err != nil {
		log.Fatal(err)
	}
}
