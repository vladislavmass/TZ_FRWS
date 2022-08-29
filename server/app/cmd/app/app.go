package main

import (
	"fmt"
	"os"

	phraseservice "github.com/vladislavmass/pow_ddos/server/app/internal/phrase_service"
	tcpserver "github.com/vladislavmass/pow_ddos/server/app/internal/tcp_server"
)

func main() {
	outPort := getEnv("OutPort", "8080")
	outPort = fmt.Sprintf(":%s", outPort)
	tcpserver.NewServer(outPort, phraseservice.GetStringInWarcraft3).Run()
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
