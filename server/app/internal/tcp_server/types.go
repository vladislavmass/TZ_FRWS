package tcpserver

import "net"

type GetContent func() string

type ExecutorHandler func(net.Conn)
