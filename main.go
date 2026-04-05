package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Connection accepted")
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		conn.Close()
		fmt.Println("Connection closed")
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	buffer := make([]byte, 8)
	currentLine := ""
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			chunk, err := file.Read(buffer)

			if chunk > 0 {
				parts := strings.Split(string(buffer[:chunk]), "\n")
				for i := 0; i < len(parts)-1; i++ {
					out <- currentLine + parts[i]
					currentLine = ""
				}
				currentLine += parts[len(parts)-1]
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		if currentLine != "" {
			out <- currentLine
		}
	}()
	return out

}
