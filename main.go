package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 8)
	for {
		current, err := file.Read(buffer)

		if current > 0 {
			fmt.Printf("read: %s\n", string(buffer[:current]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
