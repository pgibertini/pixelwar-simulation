package main

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent/server"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/launcher"
	"time"
)

func main() {
	// server for canvas requests
	myServer := server.NewServer("TEST", "127.0.0.1:5555")
	go myServer.Start()

	time.Sleep(time.Second)

	// Launch the pixel war
	launcher.LaunchPixelWar(
		false,
		true,
		100,
		0,
		500,
		0,
		0,
		true)
}
