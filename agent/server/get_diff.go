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

func (*Server) decodeGetDiffRequest(r *http.Request) (req agt.GetCanvasRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (srv *Server) doGetDiff(w http.ResponseWriter, r *http.Request) {
	srv.Lock()
	defer srv.Unlock()

	// vérification de la méthode de la requête
	if !srv.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := srv.decodeGetDiffRequest(r)
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

	if Debug {
		log.Printf("get_diff: place-id=%s\n", req.PlaceID)
	}

	// Retrieve the diff from the map
	diff := make([]painting.HexPixel, 0, len(srv.places[req.PlaceID].diff))
	for _, pixel := range srv.places[req.PlaceID].diff {
		diff = append(diff, pixel)
	}

	// Clear the diff map
	srv.places[req.PlaceID].diff = make(map[int]painting.HexPixel)

	// Send the diff
	resp := agt.GetDiffResponse{Diff: diff}

	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
