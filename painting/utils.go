package painting

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

func (h HexColor) ToRGB() (Color, error) {
	return Hex2RGB(h)
}

func (c Color) ToHex() HexColor {
	return RGB2Hex(c)
}

func (h HexColor) IsValid() bool {
	hexColorRegexp := regexp.MustCompile("#[0-9a-fA-F]{6}")
	return hexColorRegexp.MatchString(string(h))
}

func Hex2RGB(hex HexColor) (rgb Color, err error) {
	values, err := strconv.ParseUint(string(hex), 16, 32)

	if err != nil {
		return Color{}, err
	}

	rgb = Color{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
		A: 255,
	}

	return
}

func RGB2Hex(rgb Color) (hex HexColor) {
	hex = HexColor(fmt.Sprintf("#%02x%02x%02x", rgb.R, rgb.G, rgb.B))
	return
}

func ShuffleHexPixels(slice []HexPixel) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func FileToLayout(filePath string) (int, int, [][]HexColor, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, 0, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	// Get the dimensions of the layout
	scanner.Scan()
	str := scanner.Text()
	width, err := strconv.Atoi(str)
	if err != nil {
		return 0, 0, nil, err
	}
	scanner.Scan()
	str = scanner.Text()
	height, err := strconv.Atoi(str)
	if err != nil {
		return 0, 0, nil, err
	}

	// Create the layout
	layout := make([][]HexColor, height)
	for i := 0; i < height; i++ {
		// Create each row of the layout
		layout[i] = make([]HexColor, width)
		for j := 0; j < width; j++ {
			scanner.Scan()
			str = scanner.Text()
			layout[i][j] = HexColor(str)
		}
	}

	return width, height, layout, nil
}

func ImgLayoutToPixelList(layout [][]HexColor, xOffset, yOffset int) []HexPixel {
	pixels := make([]HexPixel, 0)
	for y, row := range layout {
		for x, color := range row {
			pixel := HexPixel{
				X:     x + xOffset,
				Y:     y + yOffset,
				Color: color,
			}
			pixels = append(pixels, pixel)
		}
	}
	return pixels
}
