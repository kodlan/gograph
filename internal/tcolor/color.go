package tcolor

import "image/color"

func Hsv(h float64) color.Color {
	r, g, b := 0.0, 0.0, 0.0
	i := int(h*6) % 6
	f := h*6 - float64(i)
	q := 1 - f
	switch i {
	case 0:
		r, g = 1, f
	case 1:
		r, g, b = q, 1, 0
	case 2:
		g, b = 1, f
	case 3:
		g, b, r = q, 1, 0
	case 4:
		b, r = 1, f
	case 5:
		b, r, g = 1, q, 0
	}
	return color.RGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), 255}
}
