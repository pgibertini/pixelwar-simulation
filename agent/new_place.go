package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"net/http"
)

func (*Server) decodeNewPlaceRequest(r *http.Request) (req newPlaceRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (srv *Server) doNewPlace(w http.ResponseWriter, r *http.Request) {
	srv.Lock()
	defer srv.Unlock()

	// vérification de la méthode de la requête
	if !srv.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := srv.decodeNewPlaceRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	id := fmt.Sprintf("place%d", len(srv.places)+1)
	place := Place{
		id:     id,
		canvas: painting.NewCanvas(req.Height, req.Width),
	}
	srv.places[id] = &place

	resp := newPlaceResponse{PlaceID: id}
	w.WriteHeader(http.StatusCreated)

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
