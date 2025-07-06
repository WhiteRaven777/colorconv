package main

import "math"

func RGBToHSL(r, g, b uint8) (float64, float64, float64) {
	rr := float64(r) / 255
	gg := float64(g) / 255
	bb := float64(b) / 255

	max := math.Max(math.Max(rr, gg), bb)
	min := math.Min(math.Min(rr, gg), bb)
	d := max - min
	var (
		h float64           // Hue
		s float64           // Saturation
		l = (max + min) / 2 // Lightness
	)

	if d == 0 {
		// Grayscale, no hue or saturation
		return 0, 0, math.Round(l*1000) / 1000
	}

	// Calculate saturation
	if l > 0.5 {
		s = d / (2 - max - min)
	} else {
		s = d / (max + min)
	}

	// Calculate hue
	switch max {
	case rr:
		h = (gg - bb) / d
	case gg:
		h = (bb-rr)/d + 2
	case bb:
		h = (rr-gg)/d + 4
	}
	h *= 60
	if h < 0 {
		h += 360
	}

	return math.Round(h*10) / 10, math.Round(s*1000) / 1000, math.Round(l*1000) / 1000
}

func HSLToRGB(h, s, l float64) (uint8, uint8, uint8) {
	s /= 100
	l /= 100

	var (
		c = (1 - math.Abs(2*l-1)) * s
		x = c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
		m = l - c/2
	)

	var rr, gg, bb float64
	switch {
	case h >= 0 && h < 60:
		rr, gg, bb = c, x, 0
	case h >= 60 && h < 120:
		rr, gg, bb = x, c, 0
	case h >= 120 && h < 180:
		rr, gg, bb = 0, c, x
	case h >= 180 && h < 240:
		rr, gg, bb = 0, x, c
	case h >= 240 && h < 300:
		rr, gg, bb = x, 0, c
	case h >= 300 && h < 360:
		rr, gg, bb = c, 0, x
	default:
		rr, gg, bb = 0, 0, 0
	}

	return uint8(math.Round((rr + m) * 255)), uint8(math.Round((gg + m) * 255)), uint8(math.Round((bb + m) * 255))
}

func RGBToHSV(r, g, b uint8) (float64, float64, float64) {
	rr := float64(r) / 255
	gg := float64(g) / 255
	bb := float64(b) / 255

	max := math.Max(math.Max(rr, gg), bb)
	min := math.Min(math.Min(rr, gg), bb)
	d := max - min

	var (
		h float64 // Hue
		s float64 // Saturation
		v = max   // Value
	)

	if max > 0 {
		s = d / max
	} else {
		s = 0
	}

	if d != 0 {
		switch max {
		case rr:
			h = (gg - bb) / d
		case gg:
			h = (bb-rr)/d + 2
		case bb:
			h = (rr-gg)/d + 4
		}
		h *= 60
		if h < 0 {
			h += 360
		}
	}

	return math.Round(h*10) / 10, math.Round(s*1000) / 1000, math.Round(v*1000) / 1000
}

func HSVToRGB(h, s, v float64) (uint8, uint8, uint8) {
	s /= 100
	v /= 100

	var (
		c = v * s
		x = c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
		m = v - c
	)

	var rr, gg, bb float64
	switch {
	case h >= 0 && h < 60:
		rr, gg, bb = c, x, 0
	case h >= 60 && h < 120:
		rr, gg, bb = x, c, 0
	case h >= 120 && h < 180:
		rr, gg, bb = 0, c, x
	case h >= 180 && h < 240:
		rr, gg, bb = 0, x, c
	case h >= 240 && h < 300:
		rr, gg, bb = x, 0, c
	case h >= 300 && h < 360:
		rr, gg, bb = c, 0, x
	default:
		rr, gg, bb = 0, 0, 0
	}

	return uint8(math.Round((rr + m) * 255)), uint8(math.Round((gg + m) * 255)), uint8(math.Round((bb + m) * 255))
}
