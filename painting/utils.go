package painting

import (
	"fmt"
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
