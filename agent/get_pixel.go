package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (*Server) decodeGetPixelRequest(r *http.Request) (req getPixelRequest, err error) {
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

	// traitement de la requête
	color := srv.places[req.PlaceID].canvas.Grid[req.X][req.Y].GetColor().ToHex()
	fmt.Println(srv.places[req.PlaceID].canvas.Grid[req.X][req.Y].GetColor())

	resp := getPixelResponse{Color: color}
	w.WriteHeader(http.StatusOK)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
