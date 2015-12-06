package main

import (
	"flag"
	"log"
	"net"
	"regexp"
	"strconv"
	"time"
)

type Counter struct {
	Name  string
	Value int64
}

var format = regexp.MustCompile(`^(?P<metric>[a-zA-Z0-9.-_]+):(?P<val>[0-9.]+)\|(?P<kind>[a-z]+)$`)

func handleStat(counterStream chan<- Counter, recv string) {
	matches := format.FindStringSubmatch(recv)
	if len(matches) < 4 {
		log.Fatalf("Received fewer than expect submatches for %v\n", recv)
	}
	metric := string(matches[1])
	kind := matches[3]
	switch kind {
	case "c":
		value, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			log.Printf("Failed to parse int from %v.\n", matches[2])
			return
		}
		counterStream <- Counter{Name: metric, Value: value}
	}
}

func flushStats(counterStream <-chan Counter, flushPeriod int) {
	ticker := time.NewTicker(time.Duration(flushPeriod) * time.Second)
	counters := make(map[string]int64)
	for {
		select {
		case <-ticker.C:
			log.Printf("Counters: %v\n", counters)
			counters = make(map[string]int64)
		case msg := <-counterStream:
			counters[msg.Name] += msg.Value
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func startServer(port int64, flushPeriod int) {
	portStr := ":" + strconv.FormatInt(port, 10)
	log.Printf("Server starting on %v.", portStr)
	ServerAddr, err := net.ResolveUDPAddr("udp", portStr)
	if err != nil {
		log.Fatal(err)
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer ServerConn.Close()

	counterStream := make(chan Counter)

	go flushStats(counterStream, flushPeriod)

	buf := make([]byte, 1024)
	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			go handleStat(counterStream, string(buf[0:n]))
		}
	}
}

func main() {
	port := flag.Int("port", 8125, "Port to bind to")
	flushPeriod := flag.Int("flush", 5, "Seconds between flushes")
	flag.Parse()
	log.Printf("Statsd binding to %v.\n", *port)
	startServer(int64(*port), *flushPeriod)
}
