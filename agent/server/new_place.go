package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"net/http"
	"time"
)

func (*Server) decodeNewPlaceRequest(r *http.Request) (req agt.NewPlaceRequest, err error) {
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
		id:         id,
		canvas:     painting.NewCanvasHex(req.Height, req.Width),
		lastAction: make(map[string]time.Time),
		cooldown:   req.Cooldown * time.Second,
		diff:       make(map[int]painting.HexPixel),
	}
	srv.places[id] = &place

	resp := agt.NewPlaceResponse{PlaceID: id}
	w.WriteHeader(http.StatusCreated)

	if Debug {
		log.Printf("new_place: place-id=%s\n", id)
	}

	serial, _ := json.Marshal(resp)
	w.Write(serial)
}
