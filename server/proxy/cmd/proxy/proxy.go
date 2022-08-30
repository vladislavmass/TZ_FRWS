package main

import (
	"fmt"
	"os"
	"strconv"

	powservice "github.com/vladislavmass/pow_ddos/server/proxy/internal/pow_service"
	proxyserver "github.com/vladislavmass/pow_ddos/server/proxy/internal/proxy_server"
)

func main() {
	inPort := getEnv("ProxyInPort", "8081")
	inPort = fmt.Sprintf(":%s", inPort)
	outAddress := getEnv("ProxyOutAddress", "127.0.0.1")
	outPort := getEnv("ProxyOutPort", "8080")
	outAddress = fmt.Sprintf("%s:%s", outAddress, outPort)
	complexity, err := strconv.Atoi(getEnv("Complexity", "5"))
	if err != nil {
		complexity = 6
	}
	timeOut, err := strconv.Atoi(getEnv("timeOut", "15"))
	if err != nil {
		timeOut = 15
	}

	pow := powservice.CreatePOWService(complexity)

	proxy := proxyserver.NewProxyServer(inPort, outAddress, complexity, timeOut, pow)
	proxy.Run()

}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
