package main

import (
	"fmt"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
)

func main() {
	myServer := agent.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	fmt.Scanln()
}
