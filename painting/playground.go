package painting

func (p *Playground) GetLength() int {
	return p.length
}

func (p *Playground) GetWidth() int {
	return p.width
}

func NewPlayground(l int, w int) *Playground {
	playzone := make([][]*Pixel, l)
	for i := 0; i < l; i++ {
		playzone[i] = make([]*Pixel, w)
		for j := 0; j < w; j++ {
			playzone[i][j] = NewPixel()
		}
	}
	return &Playground{
		length: l,
		width:  w,
		zone:   playzone,
	}
}
