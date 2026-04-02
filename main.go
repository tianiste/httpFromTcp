package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for line := range getLinesChannel(file) {
		fmt.Printf("Read: %s\n", line)
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
					fmt.Printf("read: %s\n", currentLine+parts[i])
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
