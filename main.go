package main

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent/server"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/launcher"
)

func main() {
	// server for canvas requests
	myServer := server.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	// Launch the pixel war
	launcher.LaunchPixelWar(
		false,
		true,
		1,
		5,
		500,
		1,
		1,
		true)
}
