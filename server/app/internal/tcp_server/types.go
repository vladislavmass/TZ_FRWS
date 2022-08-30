package tcpserver

import (
	"net"
	"sync"
)

type GetContent func() string

type ExecutorHandler func(net.Conn, *sync.WaitGroup)
