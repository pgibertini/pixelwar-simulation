package main

import "fmt"
import srv "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"

func main() {
	myServer := srv.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	fmt.Scanln()
}
