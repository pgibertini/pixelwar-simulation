package agent

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"time"
)

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
	PlaceID   string `json:"place-id"`
	ResetDiff bool   `json:"reset-diff"`
}

type GetCanvasResponse struct {
	Height int                   `json:"height"`
	Width  int                   `json:"width"`
	Grid   [][]painting.HexColor `json:"grid"`
}

type GetDiffResponse struct {
	Diff []painting.HexPixel `json:"diff"`
}

type sendPixelsRequest struct {
	pixels []painting.HexPixel
	id_am  string
}

type FindWorkersRequest struct {
	Id_manager string
	Hobby      string
}

type FindWorkersResponse struct {
	Workers []*AgentWorker
	//places  map[string]*Place
}
