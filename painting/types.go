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
	grid   [][]*Pixel
}
