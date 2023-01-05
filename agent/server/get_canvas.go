package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"net/http"
)

func (*Server) decodeGetCanvasRequest(r *http.Request) (req agt.GetCanvasRequest, err error) {
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

	// check if the place-id exists
	if _, exists := srv.places[req.PlaceID]; exists {
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid place-id")
		return
	}

	// traitement de la requête
	gridHeight := srv.places[req.PlaceID].canvas.GetHeight()
	gridWidth := srv.places[req.PlaceID].canvas.GetWidth()
	grid := &srv.places[req.PlaceID].canvas.Grid

	if debug {
		log.Printf("get_canvas: place-id=%s\n", req.PlaceID)
	}

	if req.ResetDiff {
		srv.places[req.PlaceID].diff = make(map[int]painting.HexPixel)
		log.Println("diff reset")
	}

	log.Println(req.ResetDiff)
	resp := agt.GetCanvasResponse{Height: gridHeight, Width: gridWidth, Grid: *grid}
	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
