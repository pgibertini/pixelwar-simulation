package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (*Server) decodeGetPixelRequest(r *http.Request) (req GetPixelRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (srv *Server) doGetPixel(w http.ResponseWriter, r *http.Request) {
	srv.Lock()
	defer srv.Unlock()

	// vérification de la méthode de la requête
	if !srv.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := srv.decodeGetPixelRequest(r)
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

	// check if the coordinates are in the canvas
	if req.X < 0 || req.X >= srv.places[req.PlaceID].canvas.GetWidth() || req.Y < 0 || req.Y >= srv.places[req.PlaceID].canvas.GetHeight() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid coordinates")
		return
	}

	// traitement de la requête
	color := srv.places[req.PlaceID].canvas.Grid[req.X][req.Y]
	//fmt.Println(srv.places[req.PlaceID].canvas.Grid[req.X][req.Y].GetColor())

	resp := GetPixelResponse{Color: color}
	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
