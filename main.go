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

	// PARAMETERS
	// nbWorkerPerManager (string): can be set to "random" (random between 5 and 50), "proportional" (proportional to the image size) or you can pass an "int"
	// cooldown (int): the cooldown (in seconds) between 2 pixels painted by a worker
	// size (int): the size (height and width) of the canvas
	// conquestValue (int): value piloting how managers will be likely to expend their territory by drawing their image at different places (0: won't expand)

	//launcher.LaunchPixelWar("proportional", 1, 500, 0, true)
	//launcher.LaunchPixelWar("random", 1, 500, 5, true)
	//launcher.LaunchPixelWar("50", 1, 200, 10, true)
	launcher.LaunchPixelWar("50", 1, 500, 10, true)
}
