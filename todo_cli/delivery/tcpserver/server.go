package tcpserver

import "net"

type Server struct {
	listener   net.Listener
	connection net.Conn
}

func (s Server) CreateTask() {

}

func (s Server) ListTask() {

}
