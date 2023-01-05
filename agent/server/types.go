package server

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
