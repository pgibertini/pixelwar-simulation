package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strconv"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("./painting/img/place_2022.png")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file, 251, 1516, 16, 12)

	var filename string

	painting.StringToColor("#10000000")

	fmt.Scanf("%s", &filename)

	path := "./painting/img/" + filename

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, v := range pixels {
		for _, v2 := range v {
			f.WriteString(string(v2))
			f.WriteString(" ")
		}
		f.WriteString("!\n")
	}
}

func getPixels(file io.Reader, begin_X int, begin_Y int, nb_lines int, nb_columns int) ([][]string, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	var pixels [][]string
	for y := begin_Y; y <= begin_Y+nb_lines; y++ {
		var row []string
		for x := begin_X; x < begin_X+nb_columns; x++ {
			row = append(row, rgbaToString(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// Converts RGBA values to a single string representing hexadecimal value of a pixel
func rgbaToString(r uint32, g uint32, b uint32, a uint32) string {
	r_h := strconv.FormatUint(uint64(r), 16)
	if r_h == "0" {
		r_h = "0000"
	}
	g_h := strconv.FormatUint(uint64(g), 16)
	if g_h == "0" {
		g_h = "0000"
	}
	b_h := strconv.FormatUint(uint64(b), 16)
	if b_h == "0" {
		b_h = "0000"
	}
	a_h := strconv.FormatUint(uint64(a), 16)

	fmt.Println(r, g, b)

	hex := "#" + r_h[0:2] + g_h[0:2] + b_h[0:2] + a_h[0:2]

	return hex
} // Returns a string representing a RGBA color in hexadecimal
