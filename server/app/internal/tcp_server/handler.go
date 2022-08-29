package tcpserver

import (
	"fmt"
	"net"
)

func NewExecutorHandler(getContent GetContent) ExecutorHandler {
	return func(conn net.Conn) {
		defer conn.Close()
		fmt.Fprintln(conn, getContent())
	}
}
