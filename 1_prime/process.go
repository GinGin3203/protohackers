package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"

	"github.com/GinGin3203/protohackers/pkg/must"
)

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func processTCPConn(c net.Conn) {
	defer c.Close()

	s := bufio.NewScanner(c)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		b := s.Bytes()
		var req Request
		if err := json.Unmarshal(b, &req); err != nil || !req.isWellFormed() {
			log.Printf("malformed request or err: %s", err)
			if _, err := c.Write([]byte("invalid")); err != nil {
				log.Println(err)
				return
			}
			return
		}
		log.Println(req)

		var re Response
		if n, err := req.Number.Int64(); err != nil {
			re = Response{
				Method: "isPrime",
				Prime:  false,
			}
		} else {
			re = Response{
				Method: "isPrime",
				Prime:  isPrime(n),
			}
		}
		log.Println(re)
		if _, err := c.Write(append(must.NotFail(json.Marshal(re)), byte('\n'))); err != nil {
			log.Println(err)
			return
		}
	}
}

func isPrime(n int64) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}

	var i int64
	for i = 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
