package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"net"
	"sync"
)

var connMap = &sync.Map{}

func main() {
	l, err := net.Listen("tcp", "192.168.0.111:6969")
	if err != nil {
		fmt.Println("Unable to start listener")
		return
	}

	fmt.Println("Successfully started listener")

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}

		fmt.Println("Accepted connection")

		go handleUserConnection(conn)
	}

}

func handleUserConnection(c net.Conn) {
	id := uuid.New().String()

	defer func() {
		c.Close()
		connMap.Delete(id)
	}()

	connMap.Store(id, c)

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return
		}

		fmt.Printf("Accepted user input: %s", userInput)

		connMap.Range(func(key, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				conn.Write([]byte(userInput))
			}

			return true
		})
	}
}
