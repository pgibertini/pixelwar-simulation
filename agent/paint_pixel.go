package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

	userLastAction, exists := srv.places[req.PlaceID].lastAction[req.UserID]
	// Regarde si l'utilisateur est déjà dans le map
	if exists {
		// Vérifie que le cooldown a été respecté
		if time.Now().Before(userLastAction.Add(srv.places[req.PlaceID].cooldown)) {
			w.WriteHeader(http.StatusTooEarly)
			fmt.Fprint(w, "Trop tôt")
			return
		}
	}
	// Update ou rajoute l'user dans le map
	srv.places[req.PlaceID].lastAction[req.UserID] = time.Now()

	// traitement de la requête
	rgb, err := req.Color.ToRGB()
	srv.places[req.PlaceID].canvas.Grid[req.X][req.Y].SetColor(rgb)

	w.WriteHeader(http.StatusOK)
}
