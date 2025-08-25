package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print(">")
    text, _ := reader.ReadString('\n')
    _, err = conn.Write([]byte(text))
    if err != nil {
      log.Fatal("error while writing to conn", err)
    }
  }
}
