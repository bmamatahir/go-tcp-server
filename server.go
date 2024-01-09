package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
)

func handleConnection(conn net.Conn, verbose bool) {
	defer conn.Close()

	var lastPacketTime time.Time

	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading:", err)
			break
		}

		currentTime := time.Now()
		duration := currentTime.Sub(lastPacketTime)

		receivedMessage := string(buffer)

		if verbose {
			log.Printf("Received data: %s\n", color.GreenString(receivedMessage))
			log.Printf(color.New(color.FgHiBlack).Sprintf("Duration since last packet: %s\n", duration.String()))
		} else {
			log.Printf("Received data: %s\n", receivedMessage)
		}

		lastPacketTime = currentTime
	}
}

func main() {
	port := flag.Int("port", 8080, "TCP port to listen on")
	verbose := flag.Bool("v", true, "Enable verbose logging")
	flag.Parse()

	if *verbose {
		log.SetFlags(log.Lmicroseconds)
		log.Println("Verbose logging enabled.")
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer listener.Close()
	log.Printf(color.New(color.FgHiBlack).Sprintf("Server listening on :%d\n", *port))

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
