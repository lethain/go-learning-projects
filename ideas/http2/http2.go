package main

import (
	"bufio"
	"flag"
	"log"
	"io"
	"strconv"
	"net"
)

var switchProto = "HTTP/1.1 101 Switching Protocols\nConnection: Upgrade\nUpgrade: h2c\n\n"

func handle(conn net.Conn) {
	log.Printf("Handling %v\n", conn)
	reader := bufio.NewReader(conn)

	buf := make([]byte, 1024)
	for {
		numBytes, err := reader.Read(buf)		
		log.Printf("Read %v bytes\n", numBytes)
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection closed.\n")
				return
			}
			log.Printf("Failed reading into buffer: %v\n", err)
			continue
		}
		log.Printf("Buffer: %v\n", string(buf))
		conn.Write([]byte(switchProto))
	}
}

func start(port int64) {
	loc := ":" + strconv.FormatInt(port, 10)
	ln, err := net.Listen("tcp", loc)
	if err != nil {
		log.Fatalf("Failed to connecting to %v.\n", loc)
	}
	defer ln.Close()	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v.\n", err)
			continue
		}
		go handle(conn)
	}
}

func main() {
	port := flag.Int64("port", 9000, "Port to bind to")
	flag.Parse()
	log.Printf("Starting on %v.\n", *port)
	start(*port)
}
