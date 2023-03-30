package main

import (
	"fmt"
	"net/http"
)

var counter = 0

func main() {
	server := httpServer{}
	http.ListenAndServe(":8090", server)

}

type httpServer struct {
}

func (hs httpServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		rw.Write([]byte(`I am hello`))
	}
	counter += 1
	fmt.Println(counter, "request arrived !")

}
