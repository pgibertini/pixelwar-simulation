package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/launcher"
	"net/http"
)

func (*Server) decodeLaunchRequest(r *http.Request) (req agt.LaunchRequest, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (srv *Server) launch(w http.ResponseWriter, r *http.Request) {
	srv.Lock()
	defer srv.Unlock()

	// vérification de la méthode de la requête
	if !srv.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := srv.decodeLaunchRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	placeID := launcher.LaunchPixelWar(req.NbWorkersPerManager, req.Cooldown, req.Size, 1, false)

	serial, _ := json.Marshal(placeID)
	w.Write(serial)

	w.WriteHeader(http.StatusOK)
}
