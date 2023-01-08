package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"net/http"
)

func NewAgentWorker(idAgt string, hobbiesAgt []string, chat *Chat, placeID string, url string) *AgentWorker {
	channel := make(chan interface{})
	return &AgentWorker{
		id:      idAgt,
		Hobbies: hobbiesAgt,
		Cin:     channel,
		Chat:    chat,
		placeId: placeID,
		srvUrl:  url,
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
				aw.paintPixel(pixel)
				aw.tab = aw.tab[1:]
				//time.Sleep(time.Second)
			}
			aw.mu.Unlock()
		}
	}()
}

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetHobbies() []string {
	return aw.Hobbies
}

func (aw *AgentWorker) paintPixel(pixel painting.HexPixel) {
	req := PaintPixelRequest{
		PlaceID: aw.placeId,
		UserID:  aw.id,
		X:       pixel.X,
		Y:       pixel.Y,
		Color:   pixel.Color,
	}

	// sérialisation de la requête
	url := aw.srvUrl + "/paint_pixel"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return
	}

	log.Printf("%s painted pixel (%d, %d) with color %s", aw.id, req.X, req.Y, req.Color)
	return
}

func (aw *AgentWorker) register() {
	(aw.Chat).Cin <- aw
}
