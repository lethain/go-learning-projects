package main

import (
	"flag"
	"log"
	"net"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var format = regexp.MustCompile(`^(?P<metric>[a-zA-Z0-9.-_]+):(?P<val>[0-9.]+)\|(?P<kind>[a-z]+)$`)

type Counters struct {
	sync.RWMutex
	Metrics map[string]int64
}

func handleStat(counters *Counters, recv string) {
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
		counters.Lock()
		counters.Metrics[metric] += value
		counters.Unlock()
	default:
		log.Printf("Unknown kind: %v.\n", kind)
	}
}

func flushStats(counters *Counters, flushPeriod int) {
	ticker := time.NewTicker(time.Duration(flushPeriod) * time.Second)
	for range ticker.C {
		counters.Lock()
		log.Printf("Counters: %v\n", counters.Metrics)
		counters.Metrics = make(map[string]int64)
		counters.Unlock()
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

	counters := Counters{Metrics: make(map[string]int64)}

	go flushStats(&counters, flushPeriod)

	buf := make([]byte, 1024)
	for {
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			go handleStat(&counters, string(buf[0:n]))
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
