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

type CanvasHex struct {
	height int
	width  int
	Grid   [][]HexColor
}

type PixelToPlace struct {
	x     int
	y     int
	pixel Pixel
}

type HexPixel struct {
	X     int      `json:"x"`
	Y     int      `json:"y"`
	Color HexColor `json:"c"`
}

type ManagerPainting struct {
	YOffset int
	XOffset int
	Height  int
	Width   int
	ImgPath string
}
