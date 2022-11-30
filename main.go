package main

import (
	"fmt"
	"image"
	"image/png"

	"io"
	"os"
	"strconv"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("./place_2022.png")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file, 256, 1524, 1524+16, 12)

	fmt.Println(pixels)
}

//256, 1524
// area = 192 pixels

// Get the bi-dimensional pixel array
func getPixels(file io.Reader, begin_X int, begin_Y int, end_Y int, length_line int) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	var pixels [][]Pixel
	for y := begin_Y; y <= end_Y; y++ {
		var row []Pixel
		for x := begin_X; x < begin_X+length_line; x++ {
			row = append(row, rgbaToHex(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToHex(r uint32, g uint32, b uint32, a uint32) Pixel {
	hex := strconv.FormatUint(uint64(r), 16) + strconv.FormatUint(uint64(g), 16) + strconv.FormatUint(uint64(b), 16)
	return Pixel{Color(hex)}
}
