package painting

import "image/color"

type Color color.RGBA

type HexColor string

type Pixel struct {
	color Color
}

type Canvas struct {
	height int
	width  int
	Grid   [][]*Pixel
}

type PixelToPlace struct {
	x     int
	y     int
	pixel Pixel
}