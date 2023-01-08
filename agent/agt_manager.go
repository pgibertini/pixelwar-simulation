package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"io"
	"log"
	"net/http"
	"sync"
)

func NewAgentManager(id string, hobby string, chat *Chat, placeID string, url string) *AgentManager {
	cin := make(chan interface{})
	cout := make(chan FindWorkersResponse)
	return &AgentManager{
		id:      id,
		hobby:   hobby,
		Chat:    chat,
		Cout:    cin,
		Cin:     cout,
		placeId: placeID,
		srvUrl:  url,
	}
}

func (am *AgentManager) Start() {
	am.register()
	am.updateWorkers()
	//am.ConvertImgToPixels("./images/BlueMario", 0, 0)
	//am.sendPixelsToWorkers()
}

func (am *AgentManager) GetID() string {
	return am.id
}

func (am *AgentManager) GetHobby() string {
	return am.hobby
}

func (am *AgentManager) register() {
	(am.Chat).Cin <- am
}

func (am *AgentManager) updateWorkers() {
	// Not sure of this. Can AgentWorkers change their hobbies?
	for k, v := range am.workers {
		if !ContainsHobby(v.Hobbies, am.hobby) {
			am.workers = remove(am.workers, k)
		}
	}
	req := FindWorkersRequest{am.id, am.hobby}
	(am.Chat).Cin <- req

	//fmt.Println("Voici ma liste initiale de workers : ", am.workers)

	var wg sync.WaitGroup
	var resp FindWorkersResponse

	wg.Add(1)
	go func() {
		defer wg.Done()
		//fmt.Println("J'attends une réponse du serveur...")
		resp = <-am.Cin
	}()

	wg.Wait()
	//fmt.Println("Réponse reçue ! : ", resp)

	for _, v := range resp.Workers {
		if exists(am.workers, v.id) == -1 {
			am.workers = append(am.workers, v)
		}
	}

	//fmt.Println("Voici ma liste finale de workers : ", am.workers)
	log.Printf("Manager %s now has %d workers", am.id, len(am.workers))
}

func (am *AgentManager) ConvertImgToPixels(imgPath string) {
	width, height, layout, err := painting.FileToLayout(imgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Set the painting dimensions
	am.Painting.Width = width
	am.Painting.Height = height
	// Set the layout
	am.ImgLayout = layout
	// Convert the layout to a list of HexPixels
}

func (am *AgentManager) sendPixelsToWorkers() {
	numWorkers := len(am.workers)

	start := 0
	end := len(am.pixelsToPlace) - 1

	intervalSize := (end + 1) / numWorkers
	remainder := (end + 1) % numWorkers

	var plusOne int
	for i := 0; i < numWorkers; i++ {
		if remainder != 0 {
			remainder--
			plusOne = 1
		}

		low := start + i*intervalSize
		high := low + intervalSize - 1 + plusOne

		fmt.Printf("interval: [%d, %d]\n", low, high)
		fmt.Println("----- interval length: ", (high+1)-low)

		start += plusOne
		plusOne = 0

		var pixelsToSend []painting.HexPixel
		for j := low; j <= high; j++ {
			pixelsToSend = append(pixelsToSend, am.pixelsToPlace[j])
			//TODO : send pixels to workers channels
		}
		request := sendPixelsRequest{pixelsToSend, am.id}
		am.workers[i].Cin <- request
	}
}

func (am *AgentManager) AddPixelsToPlace(p []painting.HexPixel) {
	am.pixelsToPlace = append(am.pixelsToPlace, p...)
}

func (am *AgentManager) divideWork() [][]painting.HexPixel {
	numWorkers := len(am.workers)
	workPerWorker := len(am.pixelsToPlace) / numWorkers
	remainder := len(am.pixelsToPlace) % numWorkers

	workList := make([][]painting.HexPixel, numWorkers)
	for i := 0; i < numWorkers; i++ {
		startIndex := i * workPerWorker
		endIndex := startIndex + workPerWorker

		if i == numWorkers-1 {
			endIndex += remainder
		}

		workList[i] = am.pixelsToPlace[startIndex:endIndex]
	}

	return workList
}

// DistributeWork distribute the list of pixel to place that are not already placed
func (am *AgentManager) DistributeWork() {
	workList := am.divideWork()
	for i, agt := range am.workers {
		request := sendPixelsRequest{workList[i], am.id}
		agt.Cin <- request
	}
}

func (am *AgentManager) GetPixelsToPlace() {
	req := GetCanvasRequest{
		PlaceID: am.placeId,
	}

	// sérialisation de la requête
	url := am.srvUrl + "/get_canvas"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal the response into a NewPlaceResponse struct
	var getCanvasResponse GetCanvasResponse
	err = json.Unmarshal(body, &getCanvasResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	//Regarde si des pixels sur le canvas ne sont pas comme ceux du modèle
	for i := 0; i < am.Painting.Width; i++ {
		for j := 0; j < am.Painting.Height; j++ {
			if getCanvasResponse.Grid[i+am.Painting.XOffset][j+am.Painting.YOffset] != am.ImgLayout[i][j] {
				am.pixelsToPlace = append(am.pixelsToPlace, painting.HexPixel{X: i, Y: j, Color: am.ImgLayout[i][j]})
			}
		}
	}
}

func (am *AgentManager) getPixelRequest(x, y int) (color painting.HexColor, err error) {
	req := GetPixelRequest{
		PlaceID: am.placeId,
		X:       x,
		Y:       y,
	}

	// sérialisation de la requête
	url := am.srvUrl + "/paint_pixel"
	data, err := json.Marshal(req)
	if err != nil {
		return
	}

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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal the response into a NewPlaceResponse struct
	var getPixelResponse GetPixelResponse
	err = json.Unmarshal(body, &getPixelResponse)
	if err != nil {
		log.Printf("Error unmarshalling response: %v\n", err)
		return
	}

	return getPixelResponse.Color, nil
}

// GetUnplacedPixels return the slice of am.PixelToPlace that are not already placed, using getPixelRequest method
func (am *AgentManager) GetUnplacedPixels() []painting.HexPixel {
	unplacedPixels := make([]painting.HexPixel, 0)
	for _, pixel := range am.pixelsToPlace {
		color, err := am.getPixelRequest(pixel.X, pixel.Y)
		if err != nil {
			log.Printf("Error getting pixel color: %v\n", err)
			continue
		}
		if color != pixel.Color {
			unplacedPixels = append(unplacedPixels, pixel)
		}
	}
	return unplacedPixels
}
