package main

import (
	"bufio"
	"strconv"
	"log"
	"io"
	"net"
	"flag"
	"fmt"
)

var loc = flag.String("loc", ":6379", "location:port to run server")


type RedisServer struct {
	Loc string
}

const RedisErrorFmt = "-Error %v\r\n"
const RedisOk = "+OK\r\n"


type Command struct {
	Type string
	Params string
}


type Scanner struct {
	buffer *bufio.Reader
}

func (s *Scanner) Scan() ([]byte, error) {
	buf := make([]byte, 0)

	for {
		next, err := s.buffer.ReadByte()
		if err != nil {
			return buf, err
		}
		switch next {
		case '\r':
			// peek to see if it's a \n
			forwardArr, err := s.buffer.Peek(1)
			if err != nil {
				return buf, err
			} else if len(forwardArr) == 0 {
				return buf, io.EOF
			}
			switch forwardArr[0] {
			case '\n':
				// discard the \n
				s.buffer.ReadByte()
				return buf, nil
			default:
				buf = append(buf, next)
			}
		default:
			buf = append(buf, next)
		}	
	}
}

func EnsureByte(char byte, chars []byte) error {
	if chars[0] != char {
		return fmt.Errorf("must start with %v, not %v", string(char), string(chars[0]))
	}
	return nil
}

func (s *Scanner) VerifyAndParse(char byte) (int, error) {
	next, err := s.Scan()
	if err != nil {
		return 0, err
	}	
	if err := EnsureByte(char, next); err != nil {
		return 0,  err
	}
	num, err := strconv.ParseInt(string(next[1:]), 10, 16)
	if err != nil {
		return 0, err
	}
	return int(num), nil	

}

func (s *Scanner) ParseNumCommands() (int, error) {
	return s.VerifyAndParse('*')
}

func (s *Scanner) ParseCmdSize() (int, error) {
	return s.VerifyAndParse('$')
}

func (s *Scanner) ScanCommand() ([][]byte, error) {
	parts := make([][]byte, 0)
	n, err := s.ParseNumCommands()
	if err != nil {
		return parts, err
	}

	for i := 0; i < n; i++ {
		size, err := s.ParseCmdSize()
		if err != nil {
			return parts, err
		}
		var  cmd []byte
		for len(cmd) < size {
			nextCmd, err := s.Scan()
			if err != nil {
				return parts, nil
			}
			cmd = append(cmd, nextCmd...)
		}
		parts = append(parts, cmd)
	}
	return parts, nil
}

func (rs *RedisServer) HandleConn(conn net.Conn) {
	defer conn.Close()
	scanner := Scanner{bufio.NewReader(conn)}
	for {
		parts, err := scanner.ScanCommand()
		
		if err != nil {
			if err != io.EOF {
				log.Printf("error reading connecting: %v", err)
			}
			return
		}

		for i, part := range parts {
			log.Printf("%v\t%v", i, string(part))
		}
		conn.Write([]byte(RedisOk))
	}
}

func (rs *RedisServer) ListenAndServe() {
	ln, err := net.Listen("tcp", rs.Loc)
	defer ln.Close()
	if err != nil {
		log.Fatalf("error initializing server: %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("error accepting connection: %v.\n", err)
			continue
		}
		go rs.HandleConn(conn)
	}	
}

func main() {
	flag.Parse()
	rs := RedisServer{Loc: *loc}
	rs.ListenAndServe()
}
