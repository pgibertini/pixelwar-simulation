package agent

import (
	"bufio"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"os"
	"sync"
)

func NewAgentManager(idAgt string, hobbyAgt string, srv *Server) *AgentManager {
	cin := make(chan interface{})
	cout := make(chan FindWorkersResponse)
	return &AgentManager{
		id:            idAgt,
		hobby:         hobbyAgt,
		Srv:           srv,
		Cin:           cin,
		C_findWorkers: cout}
}

func (am *AgentManager) Start() {
	am.register()
	am.updateWorkers()
	am.convertImgToPixels(".\\usa", 0, 0)
	am.sendPixelsToWorkers()
}

func (am *AgentManager) GetID() string {
	return am.id
}

func (am *AgentManager) register() {
	(am.Srv).Cin <- am
}

func (am *AgentManager) updateWorkers() {
	// Not sure of this. Can AgentWorkers change their hobbies?
	for k, v := range am.agts {
		if !ContainsHobby(v.Hobbies, am.hobby) {
			am.agts = remove(am.agts, k)
		}
	}
	req := FindWorkersRequest{am.id, am.hobby}
	(am.Srv).Cin <- req

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

	for scanner.Scan() {
		str := scanner.Text()
		if str != "!" {
			tmpPixel := painting.NewPixelLocal(painting.StringToColor(str))
			ptp := painting.NewPixelToPlaceLocal(tmpPixel, x_offset, y_offset)
			am.bufferImgLayout = append(am.bufferImgLayout, ptp)
			x_offset++
		} else {
			y_offset++
		}
	}
}

func (am *AgentManager) sendPixelsToWorkers() {
	numWorkers := len(am.agts)

	start := 0
	end := len(am.bufferImgLayout) - 1

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

		var pixelsToSend []painting.PixelToPlace
		for j := low; j <= high; j++ {
			pixelsToSend = append(pixelsToSend, am.bufferImgLayout[j])
			//TODO : send pixels to workers channels
		}
		request := sendPixelsRequest{pixelsToSend, am.id}
		am.agts[i].Cout <- request
	}
}
