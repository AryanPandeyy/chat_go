package main

import (
	"time"
	"log"
	"net"
)

func handleRequest(conn net.Conn) {
	b := make([]byte, 10);
	a, err := conn.Read(b);

	if (err != nil) {
		log.Fatal(err);
	}

	log.Print(a);
	log.Print("CHRACTER ", string(b));

	// s := []byte{'a', '\n'};
	// log.Print(s);
	conn.Write(b);
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080");
	if (err != nil) {
		log.Fatal(err);
	}

	conn, err := listener.Accept();
	log.Print(conn.RemoteAddr());
	if (err != nil) {
		log.Fatal(err);
	}

	for {
		handleRequest(conn);
	}

	time.Sleep(time.Second);
}
