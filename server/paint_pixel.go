package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (*Server) decodePaintPixelRequest(r *http.Request) (req paintPixelRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (srv *Server) doPaintPixel(w http.ResponseWriter, r *http.Request) {
	srv.Lock()
	defer srv.Unlock()

	// vérification de la méthode de la requête
	if !srv.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := srv.decodePaintPixelRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	rgb, err := req.Color.ToRGB()
	srv.places[req.PlaceID].canvas.Grid[req.X][req.Y].SetColor(rgb)

	w.WriteHeader(http.StatusOK)
}
