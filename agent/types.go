package agent

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
	// TODO: add map of ID/timestamp to know when the last pixel has been placed by an agent
	// TODO: add an attribute defining the cooldown between the placement of 2 pixels
}

type newPlaceRequest struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	// TODO: add a parameter defining the cooldown between the placement of 2 pixels
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

type getPixelRequest struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	PlaceID string `json:"place-id"`
}

type getPixelResponse struct {
	Color painting.HexColor `json:"color"`
}
