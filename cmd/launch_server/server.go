package main

import "fmt"
import "pixelwar/server"

func main() {
	myServer := server.NewServer("TEST", "127.0.0.1:5555")
	go myServer.Start()

	fmt.Scanln()
}
