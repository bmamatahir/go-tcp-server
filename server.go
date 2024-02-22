package main

import (
	"encoding/hex"
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
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading:", err)
			break
		}

		currentTime := time.Now()
		duration := currentTime.Sub(lastPacketTime)

		receivedHex := hex.EncodeToString(buffer[:n])

		if verbose {
			log.Printf("Received data (string): %s\n", buffer[:n])
			log.Printf("Received data (bytes): %v\n", buffer[:n])
			log.Printf("Received data (hex): %s\n", receivedHex)
			log.Printf(color.New(color.FgHiBlack).Sprintf("Duration since last packet: %s\n", duration.String()))
		} else {
			log.Printf("Received data: %s\n", receivedHex)
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
