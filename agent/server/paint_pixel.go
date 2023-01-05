package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"log"
	"net/http"
	"time"
)

func (*Server) decodePaintPixelRequest(r *http.Request) (req agt.PaintPixelRequest, err error) {
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

	// check if the user-id is not empty
	if req.UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, fmt.Sprintf("empty user-id"))
		return
	}

	userLastAction, exists := srv.places[req.PlaceID].lastAction[req.UserID]
	// check if the user is already in the map of cooldown
	if exists {
		// Vérifie que le cooldown a été respecté
		if wait := userLastAction.Add(srv.places[req.PlaceID].cooldown).Sub(time.Now()).Seconds(); wait > 0 {
			w.WriteHeader(http.StatusTooEarly)
			fmt.Fprint(w, fmt.Sprintf("Please wait %f seconds", wait))
			return
		}
	}

	// Check if the color is valid
	if !req.Color.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid color")
		return
	}

	// traitement de la requête
	if debug {
		log.Printf("paint_pixel: place-id=%s ; user-id=%s ; coord=(%d, %d) ; color=%s\n", req.PlaceID, req.UserID, req.X, req.Y, req.Color)
	}

	// Update ou rajoute l'user dans le map
	srv.places[req.PlaceID].lastAction[req.UserID] = time.Now()

	// rgb, err := req.Color.ToRGB()
	srv.places[req.PlaceID].canvas.Grid[req.X][req.Y] = req.Color

	w.WriteHeader(http.StatusOK)
}
