package main

import "fmt"

type Color string

type Pixel struct {
	color Color
}

type Playground struct {
	length int
	width  int
	zone   [][]*Pixel
}

func NewPixel() *Pixel {
	return &Pixel{
		color: "white",
	}
}

func (p *Pixel) GetPixelColor() Color {
	return p.color
}

func (p *Pixel) SetPixelColor(color Color) {
	p.color = color
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

func (p *Playground) GetLength() int {
	return p.length
}

func (p *Playground) GetWidth() int {
	return p.width
}

func main() {
	myPlay := NewPlayground(100, 100)

	for i := 0; i < myPlay.GetLength(); i++ {
		for j := 0; j < myPlay.GetWidth(); j++ {
			fmt.Print(myPlay.zone[i][j].GetPixelColor(), " ")
		}
		fmt.Println("")
	}
}
