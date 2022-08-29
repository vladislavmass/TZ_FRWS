package proxyserver

import "net"

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

func (proxy ProxyServer) Run() error {
	tcpServer, err := net.Listen("tcp", proxy.inPort)
	if err != nil {
		return err
	}
	defer tcpServer.Close()
	for {
		conn, err := tcpServer.Accept()
		if err != nil {
			return err
		}
		go proxy.Handler(conn)
	}
}
