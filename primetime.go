package main

import (
	"bufio"
	"encoding/json"
	"math/big"
	"net"

	"protohackers/server"
)

type Request struct {
	Method *string  `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func primetime(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var req Request
		data := scanner.Bytes()

		err := json.Unmarshal(data, &req)
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
		if req.Method == nil {
			conn.Write([]byte("Missing method\n"))
			return
		}
		if *req.Method != "isPrime" {
			conn.Write([]byte("Invalid method, expected 'isPrime'\n"))
			return
		}
		if req.Number == nil {
			conn.Write([]byte("Missing number\n"))
			return
		}

		resp, err := json.Marshal(&Response{
			Method: "isPrime",
			Prime:  big.NewInt(int64(*req.Number)).ProbablyPrime(0),
		})
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}

		conn.Write(append(resp, '\n'))
	}
}

func main() {
	server.Run(primetime)
}
