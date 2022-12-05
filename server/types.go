package server

import "sync"

type Server struct {
	sync.Mutex
	identifier string
	address    string
}

type testRequest struct {
	Value string `json:"value"`
}
