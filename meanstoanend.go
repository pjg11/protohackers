package main

import (
	"encoding/binary"
	"log"
	"net"

	"protohackers/server"
)

type Message struct {
	Type byte
	One  int32
	Two  int32
}

func meanstoanend(conn net.Conn) {
	prices := map[int32]int32{}

	for {
		var message Message
		err := binary.Read(conn, binary.BigEndian, &message)
		if err != nil {
			return
		}

		switch message.Type {

		case 'I':
			if _, ok := prices[message.One]; ok {
				continue
			}
			prices[message.One] = message.Two

		case 'Q':
			sum, count := 0, 0
			if message.One > message.Two {
				binary.Write(conn, binary.BigEndian, int32(sum))
				continue
			}
			for timestamp, price := range prices {
				if timestamp >= message.One && timestamp <= message.Two {
					sum += int(price)
					count += 1
				}
			}
			if count > 1 {
				sum = sum / count
			}
			binary.Write(conn, binary.BigEndian, int32(sum))

		default:
			return
		}
	}
}

func main() {
	err := server.RunTCP(meanstoanend)
	if err != nil {
		log.Fatal(err)
	}
}
