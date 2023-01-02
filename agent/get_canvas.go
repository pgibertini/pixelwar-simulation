package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"net/http"
)

func (*Server) decodeGetCanvasRequest(r *http.Request) (req getCanvasRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (srv *Server) doGetCanvas(w http.ResponseWriter, r *http.Request) {
	srv.Lock()
	defer srv.Unlock()

	// vérification de la méthode de la requête
	if !srv.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := srv.decodeGetCanvasRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	if req.PlaceID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("incorrect place ID in getCanvasRequest")
		return
	}

	gridHeight := srv.places[req.PlaceID].canvas.GetHeight()
	gridWidth := srv.places[req.PlaceID].canvas.GetWidth()

	// Créé un tableau avec l'hexa de chaque pixel
	grid := make([][]painting.HexColor, gridHeight)
	for i := 0; i < gridHeight; i++ {
		grid[i] = make([]painting.HexColor, gridWidth)
		for j := 0; j < gridWidth; j++ {
			grid[i][j] = srv.places[req.PlaceID].canvas.Grid[i][j].GetColor().ToHex()
		}
	}

	resp := getCanvasResponse{gridHeight, gridWidth, grid}
	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
