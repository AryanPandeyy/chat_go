package main

import (
	"log"
	"net"
)

func handleRequest(conn net.Conn) {
	b := make(chan []byte);
	conn.Read(b);
	s := []byte{'a', '\n'};
	conn.Write(s);
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080");
	if (err != nil) {
		log.Fatal(err);
	}
	conn, err := listener.Accept();
	if (err != nil) {
		log.Fatal(err);
	}
	go handleRequest(conn);
}
