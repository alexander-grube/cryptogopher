package utils

import "math"

// RgbToHSL converts a color from RGB to HSL.
func RgbToHSL(r, g, b uint8) (float64, float64, float64) {
	rr := float64(r) / 255
	gg := float64(g) / 255
	bb := float64(b) / 255

	max := math.Max(rr, math.Max(gg, bb))
	min := math.Min(rr, math.Min(gg, bb))

	l := (max + min) / 2

	if max == min {
		return 0, 0, l
	}

	var h, s float64
	d := max - min
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}

	switch max {
	case rr:
		h = (gg - bb) / d
		if g < b {
			h += 6
		}
	case gg:
		h = (bb-rr)/d + 2
	case bb:
		h = (rr-gg)/d + 4
	}
	h /= 6

	return h, s, l
}

// HslToRGB converts a color from HSL to RGB.
func HslToRGB(h, s, l float64) (uint8, uint8, uint8) {
	var r, g, b float64
	if s == 0 {
		v := clamp(l * 255)
		return v, v, v
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r = hueToRGB(p, q, h+1/3.0)
	g = hueToRGB(p, q, h)
	b = hueToRGB(p, q, h-1/3.0)

	return clamp(r * 255), clamp(g * 255), clamp(b * 255)
}

// clamp rounds and clamps float64 value to fit into uint8.
func clamp(x float64) uint8 {
	v := int64(x + 0.5)
	if v > 255 {
		return 255
	}
	if v > 0 {
		return uint8(v)
	}
	return 0
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1/2.0 {
		return q
	}
	if t < 2/3.0 {
		return p + (q-p)*(2/3.0-t)*6
	}
	return p
}
