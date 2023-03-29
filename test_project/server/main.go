package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		network = "tcp"
		address = "127.0.0.1:2022"
		// address = "127.0.0.1:8000"

	)

	// create listener
	var listener net.Listener
	if l, lErr := net.Listen(network, address); lErr != nil {
		fmt.Println("lError :", lErr)
	} else {
		listener = l
		defer listener.Close()
	}

	fmt.Println("listener address :", listener.Addr())

	// listen for new connection
	var connection net.Conn
	for {

		if c, aErr := listener.Accept(); aErr != nil {
			log.Println("aError :", aErr)

			continue
		} else {
			connection = c
		}

		fmt.Println("client address:", connection.RemoteAddr())

		// process request
		var data = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(data)
		if rErr != nil {
			log.Println("rError", rErr)

			continue
		}
		fmt.Println(numberOfReadBytes, string(data))

		if string(data[:5]) == "break" {

			break
		}

		data = []byte(`your message recieved "_"`)
		numberOfReadBytes, rErr = connection.Write(data)
		if rErr != nil {
			log.Println("rError", rErr)

			continue
		}
		fmt.Println(numberOfReadBytes, string(data))

		connection.Close()
	}

}
