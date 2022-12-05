package painting

type Pixel struct {
	color Color
}

type Color string // TODO: use image.color.RGBA data structures instead of a string

type Playground struct {
	length int
	width  int
	zone   [][]*Pixel
}
