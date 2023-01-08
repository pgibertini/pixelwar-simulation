package agent

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func NewAgentManager(idAgt string, hobbyAgt string, chat *Chat, placeID string, url string) *AgentManager {
	cin := make(chan interface{})
	cout := make(chan FindWorkersResponse)
	return &AgentManager{
		id:            idAgt,
		hobby:         hobbyAgt,
		Chat:          chat,
		Cin:           cin,
		C_findWorkers: cout,
		placeId:       placeID,
		srvUrl:        url,
	}
}

func (am *AgentManager) Start() {
	//am.register()
	//am.updateWorkers()
	am.convertImgToPixels("./usa", 0, 0)
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
	for k, v := range am.agts {
		if !ContainsHobby(v.Hobbies, am.hobby) {
			am.agts = remove(am.agts, k)
		}
	}
	req := FindWorkersRequest{am.id, am.hobby}
	(am.Chat).Cin <- req

	fmt.Println("Voici ma liste initiale de workers : ", am.agts)

	var wg sync.WaitGroup
	var resp FindWorkersResponse

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("J'attends une réponse du serveur...")
		resp = <-am.C_findWorkers
	}()

	wg.Wait()
	fmt.Println("Réponse reçue ! : ", resp)

	for _, v := range resp.Workers {
		if exists(am.agts, v.id) == -1 {
			am.agts = append(am.agts, v)
		}
	}

	fmt.Println("Voici ma liste finale de workers : ", am.agts)

}

// Shall we specify the offset right now or shall we make another function which adds the offset?
// Because at this point, the manager does not know the size of the image and where to place it
func (am *AgentManager) convertImgToPixels(img_path string, x_offset int, y_offset int) {
	f, err := os.Open(img_path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	//Regarde la première ligne pour obtenir les dimensions du tableau
	scanner.Scan()
	str := scanner.Text()
	width, err := strconv.Atoi(str)
	scanner.Scan()
	str = scanner.Text()
	height, err := strconv.Atoi(str)

	am.Painting.Width = width
	am.Painting.Height = height
	am.imgLayout = make([][]painting.HexColor, height)

	for i := 0; i < height; i++ {
		am.imgLayout[i] = make([]painting.HexColor, width)
		for j := 0; j < width; j++ {
			scanner.Scan()
			str = scanner.Text()
			am.imgLayout[i][j] = painting.HexColor(str)
		}
	}

	println(am.Painting.Width, ";", am.Painting.Height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			print(am.imgLayout[i][j], " ")
		}
		println()
	}
}

func (am *AgentManager) sendPixelsToWorkers() {
	numWorkers := len(am.agts)

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
		am.agts[i].Cout <- request
	}
}

func (am *AgentManager) AddPixelsToBuffer(p []painting.HexPixel) {
	am.pixelsToPlace = append(am.pixelsToPlace, p...)
}

func (am *AgentManager) divideWork() [][]painting.HexPixel {
	numWorkers := len(am.agts)
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

func (am *AgentManager) DistributeWork() {
	workList := am.divideWork()
	for i, agt := range am.agts {
		request := sendPixelsRequest{workList[i], am.id}
		agt.Cout <- request
		// TODO : have the channel saved directly
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
			if getCanvasResponse.Grid[i+am.Painting.XOffset][j+am.Painting.YOffset] != am.imgLayout[i][j] {
				am.pixelsToPlace = append(am.pixelsToPlace, painting.HexPixel{X: i, Y: j, Color: am.imgLayout[i][j]})
			}
		}
	}
}
