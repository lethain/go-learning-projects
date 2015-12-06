package main

import (
	"net"
	"log"
	"bufio"
	"io"
	"regexp"
	"fmt"
	"strconv"
)

var cmdFormat = regexp.MustCompile("^([a-z]+) ")
var incrDecrFormat = regexp.MustCompile("^([a-zA-Z0-9._-]+) ([0-9.]+)\r\n$")
var getFormat = regexp.MustCompile("^([a-zA-Z0-9._-]+)\r\n$")

var kvs = make(map[string]string)

func parseCommand(msg string) (string, string, error) {
	m := cmdFormat.FindStringSubmatch(msg)
	if len(m) < 2 {
		return "", "", fmt.Errorf("Couldn't extract command from %v", msg)
	}
	return m[1], msg[len(m[1])+1:], nil
}


func handleGet(msg string) string {
	m := getFormat.FindStringSubmatch(msg)
	if len(m) < 2 {
		return fmt.Sprintf("CLIENT_ERROR couln't extract key, value from %v\r\n", msg)
	}
	key, val := m[1], kvs[m[1]]
	resp := fmt.Sprintf("VALUE %v %v %v\r\n%v\r\nEND\r\n", key, 0, len(val), val)
	log.Print(resp)
	return resp
}


/*
Handle incrementing and decrementing operations.

incr <key> <value> [noreply]\r\n
decr <key> <value> [noreply]\r\n
*/
func handleIncrDecr(cmd string, msg string) string {
	m := incrDecrFormat.FindStringSubmatch(msg)
	if len(m) == 0 {
		return fmt.Sprintf("CLIENT_ERROR couln't extract key, value from %v\r\n", msg)
	}
	key, val := m[1], m[2]

	// increment/decrement existing value
	existingStr := kvs[key]
	var existing int64
	var err error
	if existingStr != "" {
		existing, err = strconv.ParseInt(existingStr, 10, 64)
		if err != nil {
			log.Printf("Couldn't parse existing: %v.\n", existingStr)
		}
	}
	mod, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Printf("%v", err)
		return fmt.Sprintf("ERROR %v\r\n", err)
	}
	
	if cmd == "incr" {
		kvs[key] = strconv.FormatInt(existing + mod, 10)
	} else if cmd == "decr" {
		kvs[key] = strconv.FormatInt(existing - mod, 10)
	}
	return fmt.Sprintf("%v\r\n", kvs[key])
}


func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// CMD, flags, TTL, size in bytes; "set some_key 0 0 10"
		// <command name> <key> <flags> <exptime> <bytes> [noreply]\r\n
		// cas <key> <flags> <exptime> <bytes> <cas unique> [noreply]\r\n
		// delete <key> [noreply]\r\n
		// touch <key> <exptime> [noreply]\r\n
		// get <key>*\r\n
		// gets <key>*\r\n
		// VALUE <key> <flags> <bytes> [<cas unique>]\r\n
		// <data block>\r\n
		msg, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			if err != io.EOF {
				log.Printf("Error parsing cmd: %v\n", err)
			}
			return
		}
		cmd, rest, _ := parseCommand(msg)
		log.Printf("Help %v\n", rest)
		switch cmd {
		case "set":
			log.Printf(msg)
			conn.Write([]byte("STORED\r\n"))
		case "get":
			conn.Write([]byte(handleGet(rest)))
		case "incr", "decr":
			conn.Write([]byte(handleIncrDecr(cmd, rest)))
		default:
			log.Printf("Unsupported command: %v\n", msg)
			conn.Write([]byte("ERROR unsupported command\r\n"))
		}
	}
}


func main() {
	loc := ":11211"
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
		go handleConnection(conn)
	}
}
