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

	launcher.LaunchPixelWar("proportional", 1, 500, 0, true)
	//launcher.LaunchPixelWar("random", 1, 500, 5, true)
	//launcher.LaunchPixelWar("50", 1, 200, 10, true)
}
