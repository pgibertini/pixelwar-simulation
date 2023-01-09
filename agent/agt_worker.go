package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"net/http"
	"time"
)

func NewAgentWorker(id string, hobbies []string, chat *Chat) *AgentWorker {
	channel := make(chan interface{})
	return &AgentWorker{
		id:      id,
		Hobbies: hobbies,
		Cin:     channel,
		Chat:    chat,
	}
}

func (aw *AgentWorker) Start() {
	aw.register()

	// écoute les instructions
	go func() {
		for {
			value := <-aw.Cin
			switch value.(type) {
			case sendPixelsRequest:
				aw.mu.Lock()
				aw.tab = append(aw.tab, value.(sendPixelsRequest).pixels...)
				aw.mu.Unlock()
			default:
				fmt.Println("Error: bad request")
			}
		}
	}()

	// place des pixels
	go func() {
		for {
			aw.mu.Lock()
			if len(aw.tab) > 0 {
				pixel := aw.tab[0]
				err := aw.paintPixelRequest(pixel)
				aw.tab = aw.tab[1:]
				if err != nil {
					log.Println(err, pixel)
				}
				time.Sleep(time.Second * time.Duration(aw.Chat.GetCooldown()))
			}
			aw.mu.Unlock()
		}
	}()
}

// GETTERS

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetHobbies() []string {
	return aw.Hobbies
}

func (aw *AgentWorker) register() {
	(aw.Chat).Cin <- aw
}

// HTTP REQUEST

func (aw *AgentWorker) paintPixelRequest(pixel painting.HexPixel) (err error) {
	req := PaintPixelRequest{
		PlaceID: aw.Chat.GetPlaceID(),
		UserID:  aw.id,
		X:       pixel.X,
		Y:       pixel.Y,
		Color:   pixel.Color,
	}

	// sérialisation de la requête
	url := aw.Chat.GetURL() + "/paint_pixel"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	//log.Printf("%s painted pixel (%d, %d) with color %s", aw.id, req.X, req.Y, req.Color)
	return
}
