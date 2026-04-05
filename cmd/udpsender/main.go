package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	address, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		if _, err := conn.Write([]byte(line)); err != nil {
			log.Println(err)
			continue
		}
	}
}
