package server

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"sync"
)

type Server struct {
	sync.Mutex
	identifier string
	address    string
	places     map[string]*Place
}

type Place struct {
	id     string
	canvas *painting.Canvas
}

type testRequest struct {
	Value string `json:"value"`
}

type newPlaceRequest struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type newPlaceResponse struct {
	PlaceID string `json:"place-id"`
}

type paintPixelRequest struct {
	X       int               `json:"x"`
	Y       int               `json:"y"`
	Color   painting.HexColor `json:"color"`
	PlaceID string            `json:"place-id"`
}
