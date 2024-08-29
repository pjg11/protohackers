package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	"protohackers/server"
)

var (
	re            = regexp.MustCompile(`(^|\s)7[A-Za-z0-9]{25,34}(\s|$)`)
	protocol      = "tcp"
	upstreamAddr  = "chat.protohackers.com:16963"
	tonyBoguscoin = "7YWHMfk9JZe0LM0g1ZauHuiSxhI"
)

func main() {
	err := server.RunTCP(mobinthemiddle)
	if err != nil {
		log.Fatal(err)
	}
}

func mobinthemiddle(conn net.Conn) {
	upstream, err := net.Dial(protocol, upstreamAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer upstream.Close()

	go hackmessages(upstream, conn)
	hackmessages(conn, upstream)
}

func hackmessages(src, dst net.Conn) {
	reader := bufio.NewReader(src)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.ReplaceAll(msg, " ", "  ")
		msg = re.ReplaceAllString(msg, fmt.Sprintf("${1}%s${2}", tonyBoguscoin))
		fmt.Fprintf(dst, strings.ReplaceAll(msg, "  ", " "))
	}
}
