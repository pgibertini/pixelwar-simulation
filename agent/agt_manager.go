package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"io"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

func NewAgentManager(id string, hobby string, conquestValue float64, chat *Chat) *AgentManager {
	cin := make(chan interface{})
	cout := make(chan FindWorkersResponse)
	return &AgentManager{
		id:            id,
		hobby:         hobby,
		Chat:          chat,
		Cout:          cin,
		Cin:           cout,
		conquestValue: conquestValue,
	}
}

// GETTERS

func (am *AgentManager) GetID() string {
	return am.id
}

func (am *AgentManager) GetHobby() string {
	return am.hobby
}

func (am *AgentManager) Start() {
	am.register()
	am.updateWorkers()
	am.LoadLayoutFromFile(filepath.Join("images", am.GetHobby()))

	// image de base à placer
	am.AddPixelsToPlace(painting.ImgLayoutToPixelList(
		am.ImgLayout,
		rand.Intn(am.Chat.GetWidth()-am.Painting.Width),
		rand.Intn(am.Chat.GetHeight()-am.Painting.Height),
	))

	workList := am.divideWork(am.pixelsToPlace)
	am.DistributeWork(workList)
	time.Sleep(time.Duration(am.Chat.cooldown*len(workList[0])) * time.Second)

	// maintient ou agrandit son territoire
	go func() {
		for {
			unplacedPixels := am.GetUnplacedPixels()
			if len(unplacedPixels) < int(float64(len(am.workers))*am.conquestValue) {
				am.AddPixelsToPlace(painting.ImgLayoutToPixelList(
					am.ImgLayout,
					rand.Intn(am.Chat.GetWidth()-am.Painting.Width),
					rand.Intn(am.Chat.GetHeight()-am.Painting.Height),
				))
			}
			workList := am.divideWork(am.GetUnplacedPixels())

			am.DistributeWork(workList)
			time.Sleep(time.Duration(am.Chat.cooldown*len(workList[0])) * time.Second)
		}
	}()
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
	if Debug {
		log.Printf("Manager %s now has %d workers", am.id, len(am.workers))
	}
}

// LoadLayoutFromFile load a layout from a given file
func (am *AgentManager) LoadLayoutFromFile(imgPath string) {
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

// AddPixelsToPlace add to pixels to the pixelsToPlace slice
func (am *AgentManager) AddPixelsToPlace(p []painting.HexPixel) {
	am.pixelsToPlace = append(am.pixelsToPlace, p...)
}

// divideWork divide the given slice of pixel to place in a number of slices corresponding to the number of workers
func (am *AgentManager) divideWork(pixels []painting.HexPixel) [][]painting.HexPixel {
	numWorkers := len(am.workers)
	workPerWorker := len(pixels) / numWorkers
	remainder := len(pixels) % numWorkers

	workList := make([][]painting.HexPixel, numWorkers)
	for i := 0; i < numWorkers; i++ {
		startIndex := i * workPerWorker
		endIndex := startIndex + workPerWorker + 1
		if endIndex > len(pixels) {
			endIndex = len(pixels)
		}
		workList[i] = pixels[startIndex:endIndex]
	}

	// add the remainder
	for i := 0; i < remainder; i++ {
		workList[i] = append(workList[i], pixels[(numWorkers)*workPerWorker+i])
	}

	return workList
}

// DistributeWork distribute the list of pixel to place that are not already placed
func (am *AgentManager) DistributeWork(workList [][]painting.HexPixel) {
	for i, agt := range am.workers {
		request := sendPixelsRequest{workList[i], am.id}
		agt.Cin <- request
		time.Sleep(time.Duration(2*rand.Float64()*(float64(am.Chat.GetCooldown())/float64(len(am.workers)))) * time.Second)
	}
}

// GetUnplacedPixels return the slice of am.PixelToPlace that are not already placed, using getCanvasRequest method
func (am *AgentManager) GetUnplacedPixels() []painting.HexPixel {
	unplacedPixels := make([]painting.HexPixel, 0)

	// Get the current state of the canvas
	grid, err := am.getCanvasRequest()
	if err != nil {
		log.Printf("Error getting canvas: %v\n", err)
		return unplacedPixels
	}

	// Check which pixels in pixelsToPlace have not been placed on the canvas
	for _, pixel := range am.pixelsToPlace {
		if grid[pixel.X][pixel.Y] != pixel.Color {
			unplacedPixels = append(unplacedPixels, pixel)
		}
	}
	return unplacedPixels
}

// GetUnplacedPixelsOneByOne return the slice of am.PixelToPlace that are not already placed, using getPixelRequest method
func (am *AgentManager) GetUnplacedPixelsOneByOne() []painting.HexPixel {
	unplacedPixels := make([]painting.HexPixel, 0)
	var wg sync.WaitGroup
	for _, pixel := range am.pixelsToPlace {
		wg.Add(1)
		go func(x, y int, color painting.HexColor) {
			defer wg.Done()
			c, err := am.getPixelRequest(x, y)
			if err != nil {
				log.Printf("Error getting pixel color: %v\n", err)
				return
			}
			if c != color {
				unplacedPixels = append(unplacedPixels, painting.HexPixel{X: x, Y: y, Color: color})
			}
		}(pixel.X, pixel.Y, pixel.Color)
	}
	wg.Wait()
	return unplacedPixels
}

// HTTP REQUESTS

// getPixelRequest do a getPixel request to the server and return the response color
func (am *AgentManager) getPixelRequest(x, y int) (color painting.HexColor, err error) {
	req := GetPixelRequest{
		PlaceID: am.Chat.GetPlaceID(),
		X:       x,
		Y:       y,
	}

	// sérialisation de la requête
	url := am.Chat.GetURL() + "/get_pixel"
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

// getCanvasRequest do a getCanvas request to the server and return the response grid
func (am *AgentManager) getCanvasRequest() (grid [][]painting.HexColor, err error) {
	req := GetCanvasRequest{
		PlaceID: am.Chat.GetPlaceID(),
	}

	// sérialisation de la requête
	url := am.Chat.GetURL() + "/get_canvas"
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

	return getCanvasResponse.Grid, nil
}
