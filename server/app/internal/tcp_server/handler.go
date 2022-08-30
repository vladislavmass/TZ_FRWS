package tcpserver

import (
	"fmt"
	"net"
	"sync"
)

func NewExecutorHandler(getContent GetContent) ExecutorHandler {
	return func(conn net.Conn, wg *sync.WaitGroup) {
		defer conn.Close()
		defer wg.Done()
		fmt.Fprintln(conn, getContent())
	}
}
