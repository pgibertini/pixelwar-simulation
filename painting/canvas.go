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
		grid:   grid,
	}
}
