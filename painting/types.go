package painting

import "image/color"

type Color color.RGBA

type Pixel struct {
	color Color
}

type Canvas struct {
	height int
	width  int
	grid   [][]*Pixel
}

type PixelToPlace struct {
	x     int
	y     int
	pixel Pixel
}
