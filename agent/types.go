package agent

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

type Chat struct {
	Ams []*AgentManager
	Aws []*AgentWorker
	Cin chan interface{}
}

type AgentWorker struct {
	id      string
	tab     []painting.PixelToPlace
	Hobbies []string
	Cout    chan interface{}
	Chat    *Chat
}

type AgentManager struct {
	id              string
	agts            []*AgentWorker
	hobby           string
	Chat            *Chat
	bufferImgLayout []painting.PixelToPlace
	Cin             chan interface{}
	C_findWorkers   chan FindWorkersResponse
}
