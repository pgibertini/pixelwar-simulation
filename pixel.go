package main

type Color string

type Pixel struct {
	color Color
}

func (p *Pixel) GetPixelColor() Color {
	return p.color
}

func (p *Pixel) SetPixelColor(color Color) {
	p.color = color
}

func NewPixel() *Pixel {
	return &Pixel{
		color: "white",
	}
}
