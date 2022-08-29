package tcpserver

import (
	"net"
)

type Server struct {
	port    string
	handler ExecutorHandler
}

func NewServer(port string, getContent GetContent) *Server {
	return &Server{
		port:    port,
		handler: NewExecutorHandler(getContent),
	}
}

func (server Server) Run() error {
	tcpServer, err := net.Listen("tcp", server.port)
	if err != nil {
		return err
	}
	defer tcpServer.Close()
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			return err
		}
		go server.handler(conn)

	}
	return nil
}
