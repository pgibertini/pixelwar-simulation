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
	tab     []painting.HexPixel
	Hobbies []string
	Cout    chan interface{}
	Chat    *Chat
	srvUrl  string
	placeId string
}

type AgentManager struct {
	id              string
	agts            []*AgentWorker
	hobby           string
	Chat            *Chat
	bufferImgLayout []painting.HexPixel
	Cin             chan interface{}
	C_findWorkers   chan FindWorkersResponse
	srvUrl          string
	placeId         string
}
