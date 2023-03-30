package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	var response = `HTTP/1.1 200 OK
Date: Mon, 27 Jul 2009 12:28:53 GMT
Content-Length: 12
Content-Type: text/plane
Connection: Closed

Hey there its Me`

	const (
		network = "tcp"
		address = ":8080"
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
		if _, rErr := connection.Read(data); rErr != nil {
			log.Println("rError", rErr)

			continue
		}
		method, path := parseHTTPRequest(string(data))
		fmt.Println(method, path)
		if method == "GET" {
			if path != "/" {
				var file *os.File
				if f, oErr := os.OpenFile(strings.Replace(path, "/", "", 1)+".html", os.O_RDWR, 0777); oErr != nil {
					fmt.Println("error open file :", oErr)

					continue
				} else {
					file = f
				}
				var data = make([]byte, 1024)
				n, _ := file.Read(data)
				response = fmt.Sprintf(`HTTP/1.1 200 OK
				Content-Length: %v
				Content-Type: text/html

				%s`, n, data)

			}
		}
		if _, wErr := connection.Write([]byte(response)); wErr != nil {
			log.Println("wError", wErr)

			continue
		}

		connection.Close()
	}

}

func parseHTTPRequest(data string) (string, string) {
	dataLines := strings.Split(data, "\n")
	line0 := strings.Split(dataLines[0], " ")
	method, path := string(line0[0]), string(line0[1])

	return method, path
}
