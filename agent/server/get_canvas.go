package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
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
	if _, exists := srv.Places[req.PlaceID]; exists {
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid place-id")
		return
	}

	gridHeight := srv.Places[req.PlaceID].Canvas.GetHeight()
	gridWidth := srv.Places[req.PlaceID].Canvas.GetWidth()
	grid := &srv.Places[req.PlaceID].Canvas.Grid

	if debug {
		log.Printf("get_canvas: place-id=%s\n", req.PlaceID)
	}

	resp := agt.GetCanvasResponse{Height: gridHeight, Width: gridWidth, Grid: *grid}
	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
