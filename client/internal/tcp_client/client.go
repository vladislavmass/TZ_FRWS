package tcpclient

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Work func()

type Client struct {
	address string
	port    string
	work    Job
}

func NewClient(address string, port string, work Job) *Client {
	return &Client{
		address: address,
		port:    port,
		work:    work,
	}
}

func (client Client) SendRequest() (string, error) {

	tcpConn, err := net.Dial("tcp", client.address+":"+client.port)
	if err != nil {
		return "", err
	}
	defer tcpConn.Close()

	message, err := bufio.NewReader(tcpConn).ReadString('\n')
	if err != nil {
		return "", err
	}
	taskData := strings.Split(message, ":")
	taskData[1] = strings.Replace(taskData[1], "\n", "", -1)
	complexity, err := strconv.Atoi(taskData[1])
	if err != nil {
		return "", err
	}
	workRezult, err := client.work(taskData[0], complexity)
	if err != nil {
		return "", err
	}
	_, err = fmt.Fprintln(tcpConn, workRezult)
	if err != nil {
		return "", err
	}

	message, err = bufio.NewReader(tcpConn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return message, nil
}
