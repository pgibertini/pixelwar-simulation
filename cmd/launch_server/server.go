package main

import "fmt"
import "pixelwar/server"

func main() {
	myServer := server.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	fmt.Scanln()
}
