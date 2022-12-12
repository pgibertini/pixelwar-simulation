package painting

import "strconv"

func (h HexColor) ToRGB() (Color, error) {
	return Hex2RGB(h)
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
