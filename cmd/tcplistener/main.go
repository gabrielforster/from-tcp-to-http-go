package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string)

	go func() {
		var line string

		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					close(channel)
					break
				}
				log.Fatal("error", "error", err)
			}

			index := strings.Index(string(buffer), "\n")
			if index != -1 {
				line += string(buffer[:index])
        channel <- line
				line = ""
				if index+1 < n {
					line += string(buffer[index+1 : n])
				}
			} else {
				line += string(buffer[:n])
			}
		}
	}()

	return channel
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
  if err != nil {
    log.Fatal("error", "error", err)
  }
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
    fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Printf("%s\n", line)
		}
    conn.Close()
	}
}
