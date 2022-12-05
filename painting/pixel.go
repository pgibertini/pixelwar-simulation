package painting

// colors list from https://lospec.com/palette-list/r-place-2022-day3
// TODO: make a const slice "colorPalette" instead of these 32 const string, and use the type "color" (or better: colorRGBA)
const (
	c1  = "6d001a"
	c2  = "be0039"
	c3  = "ff4500"
	c4  = "ffa800"
	c5  = "ffd635"
	c6  = "fff8b8"
	c7  = "00a368"
	c8  = "00cc78"
	c9  = "7eed56"
	c10 = "00756f"
	c11 = "009eaa"
	c12 = "00ccc0"
	c13 = "2450a4"
	c14 = "3690ea"
	c15 = "51e9f4"
	c16 = "493ac1"
	c17 = "6a5cff"
	c18 = "94b3ff"
	c19 = "811e9f"
	c20 = "b44ac0"
	c21 = "e4abff"
	c22 = "de107f"
	c23 = "ff3881"
	c24 = "ff99aa"
	c25 = "6d482f"
	c26 = "9c6926"
	c27 = "ffb470"
	c28 = "000000"
	c29 = "515252"
	c30 = "898d90"
	c31 = "d4d7d9"
	c32 = "ffffff"
)

func (p *Pixel) GetColor() Color {
	return p.color
}

func (p *Pixel) SetColor(color Color) {
	p.color = color
}

func NewPixel() *Pixel {
	return &Pixel{
		color: c28,
	}
}
