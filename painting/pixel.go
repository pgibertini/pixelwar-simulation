package painting

// colors list from https://lospec.com/palette-list/r-place-2022-day3
var colorPalette = [...]Color{
	{109, 0, 26, 255},
	{190, 0, 57, 255},
	{255, 0, 69, 255},
	{255, 168, 0, 255},
	{255, 214, 53, 255},
	{255, 248, 184, 255},
	{0, 163, 104, 255},
	{0, 204, 120, 255},
	{126, 237, 86, 255},
	{0, 117, 103, 255},
	{0, 158, 170, 255},
	{0, 204, 192, 255},
	{36, 80, 164, 255},
	{54, 144, 234, 255},
	{81, 233, 244, 255},
	{73, 58, 193, 255},
	{106, 92, 255, 255},
	{148, 179, 255, 255},
	{129, 30, 159, 255},
	{180, 74, 192, 255},
	{228, 171, 255, 255},
	{222, 16, 127, 255},
	{255, 56, 129, 255},
	{255, 153, 170, 255},
	{109, 72, 47, 255},
	{156, 105, 38, 255},
	{255, 180, 112, 255},
	{0, 0, 0, 255},
	{81, 82, 82, 255},
	{137, 141, 144, 255},
	{212, 215, 217, 255},
	{255, 255, 255, 255},
}

func NewPixel() *Pixel {
	return &Pixel{
		color: Color{255, 255, 255, 255},
	}
}

func NewPixelLocal(c Color) Pixel {
	return Pixel{
		color: Color{c.R, c.G, c.G, c.A},
	}
}

func (p *Pixel) GetColor() Color {
	return p.color
}

func (p *Pixel) SetColor(color Color) {
	p.color = color
}

func NewPixelToPlaceLocal(p Pixel, x int, y int) PixelToPlace {
	return PixelToPlace{
		x:     x,
		y:     y,
		pixel: p,
	}
}
