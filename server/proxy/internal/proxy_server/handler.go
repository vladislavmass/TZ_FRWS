package proxyserver

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

func (server ProxyServer) Handler(inConn net.Conn) {
	/*
		I know that you need to close the connection in the same place where you opened it. But I don't know how to do this
	*/
	defer inConn.Close()

	//sending authorization information
	fmt.Fprintln(inConn, server.POW.GetWork())

	//Waiting for response code
	inConn.SetReadDeadline(time.Now().Add(time.Duration(server.readTimeOut) * time.Second))
	message, err := bufio.NewReader(inConn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	//Purification and validation
	message = strings.Replace(message, "\n", "", -1)
	if !server.POW.ValidateWork(message) {
		fmt.Println("Key not valid")
		return
	}
	//Establishing a connection to the server
	outConn, err := net.Dial("tcp", server.outAdress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outConn.Close()

	outConn.SetDeadline(time.Now().Add(time.Duration(server.readTimeOut) * time.Second))
	_, err = io.Copy(inConn, outConn)
	if err != nil {
		fmt.Println(err)
		return
	}
}
