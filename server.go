package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn, verbose bool) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading:", err)
		return
	}

	log.Printf("Received data: %s\n", string(buffer))

	if verbose {
		log.Printf("Connection established at %s\n", time.Now().Format("2006-01-02 15:04:05"))
	}
}

func main() {
	port := flag.Int("port", 8080, "TCP port to listen on")
	verbose := flag.Bool("v", false, "Enable verbose logging")
	flag.Parse()

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)
		log.Println("Verbose logging enabled.")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer listener.Close()
	log.Printf("Server listening on :%d\n", *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		if *verbose {
			log.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		}

		go handleConnection(conn, *verbose)
	}
}
