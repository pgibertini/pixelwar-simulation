package painting

func (p *Canvas) GetWidth() int {
	return p.width
}

func (p *Canvas) GetHeight() int {
	return p.height
}

func NewCanvas(l int, w int) *Canvas {
	playzone := make([][]*Pixel, l)
	for i := 0; i < l; i++ {
		playzone[i] = make([]*Pixel, w)
		for j := 0; j < w; j++ {
			playzone[i][j] = NewPixel()
		}
	}

	return &Canvas{
		height: l,
		width:  w,
		zone:   playzone,
	}
}
