package main

import (
	"fmt"
	"os"

	powwork "github.com/vladislavmass/pow_ddos/client/internal/pow_work"
	tcpclient "github.com/vladislavmass/pow_ddos/client/internal/tcp_client"
)

func main() {
	netAddress := getEnv("NetAddress", "127.0.0.1")
	netPort := getEnv("NetPort", "8081")

	tcpcl := tcpclient.NewClient(netAddress, netPort, powwork.DoTheJob)
	mess, err := tcpcl.SendRequest()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mess)
}

// Get environment variable or default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
