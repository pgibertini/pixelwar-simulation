package painting

import "image/color"

type Color color.RGBA // TODO: use image.color.RGBA data structures instead of a string

type Pixel struct {
	color Color
}

type Canvas struct {
	height int
	width  int
	grid   [][]*Pixel
}
