package pixelwar

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strconv"
)

// Creates a layer file of a drawing's pixels
func GetLayerFromPNG(filename string, first_pixel_x int, first_pixel_y int, length int, width int) (err error) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	path := "./pixelwar/" + filename
	file, err := os.Open("./place_2022.png")

	if err != nil {
		return
	}

	defer file.Close()

	pixels, err := getPixels(file, first_pixel_x, first_pixel_y, length, width)

	if err != nil {
		return
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}

	defer f.Close()

	for _, v := range pixels {
		for _, v2 := range v {
			f.WriteString(string(v2))
			f.WriteString(" ")
		}
		f.WriteString("\n")
	}
	fmt.Println("The file was successfuly created. The path is:", path)

	return
} // Returns eventually an error

// Iterates for each pixel of the area calculated with the four last parameters
func getPixels(file io.Reader, first_pixel_x int, first_pixel_y int, length int, width int) ([][]Color, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	var pixels [][]Color
	for y := first_pixel_y; y <= first_pixel_y+width; y++ {
		var row []Color
		for x := first_pixel_x; x < first_pixel_x+length; x++ {
			row = append(row, rgbaToHex(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
} // Returns a 2x2 matrix of Color, and eventually an error

// Converts RGBA values to a single string representing hexadecimal value of a pixel
func rgbaToHex(r uint32, g uint32, b uint32, a uint32) Color {
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
	/* Don't uncomment until the value is required
	a_h := strconv.FormatUint(uint64(a), 16)
	if a_h == "0" {
		a_h = "0000"
	}
	*/

	fmt.Println(r, g, b)

	hex := r_h[0:2] + g_h[0:2] + b_h[0:2] // + a_h[0:2]

	return Color(hex)
} // Returns a Color in hexadecimal
