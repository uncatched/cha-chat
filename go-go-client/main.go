package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.111:6969")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Connected")
	fmt.Println("Enter your messages:")

	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		username := "uncatched"
		message := username + ":" + text
		data := []byte(message)
		_, err := conn.Write(data)
		if err != nil {
			fmt.Println("Failed to send message")
		}
	}
}
