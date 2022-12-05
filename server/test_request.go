package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Méthode de la classe "Serveur" permettant de décoder une requête POST en JSON
func (*Server) decodeTestRequest(r *http.Request) (req testRequest, err error) {

	// Lecture de la requête POST
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)

	return
} // Retourne une requête de type "testRequest", et éventuellement une erreur

func (srv *Server) makeTestRequest(w http.ResponseWriter, r *http.Request) {

	// Verrouillage du serveur le temps de créer la requête
	srv.Lock()
	defer srv.Unlock()

	// Décodage de la requête POST
	req, err := srv.decodeTestRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Request OK. The value entered is: %s", req.Value)
	w.Write([]byte(msg))
	return
}
