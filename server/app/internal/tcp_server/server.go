package tcpserver

import (
	"fmt"
	"net"
	"sync"
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

func (server Server) Run() {
	tcpServer, err := net.Listen("tcp", server.port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tcpServer.Close()
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Add(1)
		go server.handler(conn, wg)

	}
}
