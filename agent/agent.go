package agent

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

type AgentWorker struct {
	id      string
	tab     []painting.PixelToPlace
	hobbies []string
	Cout    chan interface{}
	Srv     *Server
}

type AgentManager struct {
	id              string
	agts            []*AgentWorker
	hobby           string
	Srv             *Server
	bufferImgLayout []painting.PixelToPlace
	Cin             chan interface{}
	C_findWorkers   chan findWorkersResponse
}

func NewAgentWorker(idAgt string, hobbiesAgt []string, srv *Server) *AgentWorker {
	channel := make(chan interface{})
	return &AgentWorker{
		id:      idAgt,
		hobbies: hobbiesAgt,
		Cout:    channel,
		Srv:     srv}
}

func (aw *AgentWorker) Start() {
	aw.register()
}

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetHobbies() []string {
	return aw.hobbies
}

func (aw *AgentWorker) drawOnePixel(pixel painting.Pixel) {

}

func (aw *AgentWorker) register() {
	(aw.Srv).Cin <- aw
}

// ============================ AgentManager ============================

func NewAgentManager(idAgt string, hobbyAgt string, srv *Server) *AgentManager {
	cin := make(chan interface{})
	cout := make(chan findWorkersResponse)
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

func (am *AgentManager) getID() string {
	return am.id
}

func (am *AgentManager) register() {
	(am.Srv).Cin <- am
}

func (am *AgentManager) updateWorkers() {
	// Not sure of this. Can AgentWorkers change their hobbies?
	for k, v := range am.agts {
		if !containsHobby(v.hobbies, am.hobby) {
			am.agts = remove(am.agts, k)
		}
	}
	req := findWorkersRequest{am.id, am.hobby}
	(am.Srv).Cin <- req

	fmt.Println("Voici ma liste initiale de workers : ", am.agts)

	var wg sync.WaitGroup
	var resp findWorkersResponse

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("J'attends une réponse du serveur...")
		resp = <-am.C_findWorkers
	}()

	wg.Wait()
	fmt.Println("Réponse reçue ! : ", resp)

	for _, v := range resp.workers {
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
	fmt.Println("end : ", end)

	intervalSize := (end + 1) / numWorkers
	remainder := intervalSize - (end + 1)

	fmt.Println("intervaleSize : ", intervalSize)

	for i := 0; i < numWorkers; i++ {
		low := start + i*intervalSize
		high := low + intervalSize - 1

		fmt.Printf("interval: [%d, %d]\n", low, high)
	}
}

// ============================ Utilities ============================

func containsHobby(hobbies []string, hobby string) bool {
	for _, v := range hobbies {
		if v == hobby {
			return true
		}
	}
	return false
}

func exists(s []*AgentWorker, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func getManagerIndex(s []*AgentManager, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func remove[T comparable](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func MakeRandomSliceOfHobbies(hobbies []string) (result []string) {
	result = make([]string, 0)
	for i := 0; i < 1; i++ {
		k := rand.Intn(len(hobbies))
		result = append(result, hobbies[k])
	}
	return
}
