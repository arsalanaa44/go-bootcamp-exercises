package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	const (
		network = "tcp"
		address = "127.0.0.1:2022"
	)

	var connection net.Conn
	if c, dErr := net.Dial(network, address); dErr != nil {
		fmt.Println("dError", dErr)
	} else {
		connection = c
	}
	fmt.Println("client and server address :", connection.LocalAddr(), connection.RemoteAddr())

	fmt.Println("command :", os.Args[0])
	message := "default"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	fmt.Println(connection.Write([]byte(message)))

	// process request
	var data = make([]byte, 1024)
	numberOfReadBytes, rErr := connection.Read(data)
	if rErr != nil {
		log.Fatalln("rError", rErr)
	}
	fmt.Println(numberOfReadBytes, string(data))
}
