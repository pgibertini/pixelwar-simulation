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

	diff := srv.places[req.PlaceID].canvas.Diff(srv.places[req.PlaceID].lastCanvas)

	lastCanvas := painting.NewCanvasHex(
		srv.places[req.PlaceID].canvas.GetHeight(),
		srv.places[req.PlaceID].canvas.GetWidth(),
	)

	for i := range srv.places[req.PlaceID].canvas.Grid {
		lastCanvas.Grid[i] = make([]painting.HexColor, srv.places[req.PlaceID].canvas.GetWidth())
		copy(lastCanvas.Grid[i], srv.places[req.PlaceID].canvas.Grid[i])
	}

	srv.places[req.PlaceID].lastCanvas = lastCanvas
	if debug {
		log.Printf("get_diff: place-id=%s\n", req.PlaceID)
	}

	resp := agt.GetDiffResponse{Diff: diff}
	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
