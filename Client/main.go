package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var serverip string
	fmt.Println("Enter serverip: ")
	fmt.Scanln(&serverip)
	conn, err := net.Dial("tcp", serverip)
	if err != nil {
		log.Fatalf("%s", err.Error())
		os.Exit(1)
	}
	go readServerConnection(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))

		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}
}

func readServerConnection(conn net.Conn) {
	for {
		reader := bufio.NewReader((conn))
		msg, _ := reader.ReadString('\n')
		fmt.Printf("%s", msg)
	}
}
