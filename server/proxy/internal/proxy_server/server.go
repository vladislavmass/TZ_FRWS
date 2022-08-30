package proxyserver

import (
	"fmt"
	"net"
	"sync"
)

type ProxyServer struct {
	inPort      string
	outAdress   string
	complexity  int
	readTimeOut int
	POW         POWInterface
}

func NewProxyServer(InPort, OutAdress string, complexity int, readTimeOut int, pow POWInterface) *ProxyServer {
	return &ProxyServer{
		inPort:      InPort,
		outAdress:   OutAdress,
		complexity:  complexity,
		readTimeOut: readTimeOut,
		POW:         pow,
	}
}

func (proxy ProxyServer) Run() {
	tcpServer, err := net.Listen("tcp", proxy.inPort)
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
		go proxy.Handler(conn, wg)
	}

}
