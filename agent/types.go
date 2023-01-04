package agent

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"sync"
	"time"
)

type Server struct {
	sync.Mutex
	identifier string
	address    string
	places     map[string]*Place
}

type Place struct {
	id         string
	canvas     *painting.CanvasHex
	lastAction map[string]time.Time
	cooldown   time.Duration
}

type NewPlaceRequest struct {
	Width    int           `json:"width"`
	Height   int           `json:"height"`
	Cooldown time.Duration `json:"cooldown"`
}

type NewPlaceResponse struct {
	PlaceID string `json:"place-id"`
}

type PaintPixelRequest struct {
	X       int               `json:"x"`
	Y       int               `json:"y"`
	Color   painting.HexColor `json:"color"`
	PlaceID string            `json:"place-id"`
	UserID  string            `json:"user-id"`
}

type GetPixelRequest struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	PlaceID string `json:"place-id"`
}

type GetPixelResponse struct {
	Color painting.HexColor `json:"color"`
}

type GetCanvasRequest struct {
	PlaceID string `json:"place-id"`
}

type GetCanvasResponse struct {
	Height int                   `json:"height"`
	Width  int                   `json:"width"`
	Grid   [][]painting.HexColor `json:"grid"`
}
