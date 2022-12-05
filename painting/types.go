package painting

type Color string // TODO: use image.color.RGBA data structures instead of a string

type Pixel struct {
	color Color
}

type Canvas struct {
	height int
	width  int
	zone   [][]*Pixel
}
