package main

import (
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent/server"
)

func main() {
	myServer := server.NewServer("TEST", "127.0.0.1:5555")
	go myServer.Start()

	fmt.Scanln()
}
