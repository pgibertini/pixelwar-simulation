package painting

func (p *Canvas) GetWidth() int {
	return p.width
}

func (p *Canvas) GetHeight() int {
	return p.height
}

func NewCanvas(h int, w int) *Canvas {
	grid := make([][]*Pixel, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]*Pixel, w)
		for j := 0; j < w; j++ {
			grid[i][j] = NewPixel()
		}
	}

	return &Canvas{
		height: h,
		width:  w,
		Grid:   grid,
	}
}

func (p *CanvasHex) GetWidth() int {
	return p.width
}

func (p *CanvasHex) GetHeight() int {
	return p.height
}

func NewCanvasHex(h int, w int) *CanvasHex {
	grid := make([][]HexColor, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]HexColor, w)
		for j := 0; j < w; j++ {
			grid[i][j] = "#FFFFFF"
		}
	}

	return &CanvasHex{
		height: h,
		width:  w,
		Grid:   grid,
	}
}
