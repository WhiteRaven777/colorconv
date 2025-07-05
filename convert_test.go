package main

import (
	"math"
	"testing"
)

func TestRGBToHSL(t *testing.T) {
	tests := []struct {
		name                string
		r, g, b             uint8
		wantH, wantS, wantL float64
	}{
		{
			name: "Pure Red",
			r:    255, g: 0, b: 0,
			wantH: 0.0, wantS: 1.00, wantL: 0.50,
		},
		{
			name: "Pure Green",
			r:    0, g: 255, b: 0,
			wantH: 120.0, wantS: 1.00, wantL: 0.50,
		},
		{
			name: "Pure Blue",
			r:    0, g: 0, b: 255,
			wantH: 240.0, wantS: 1.00, wantL: 0.50,
		},
		{
			name: "Pure Black",
			r:    0, g: 0, b: 0,
			wantH: 0.0, wantS: 0.00, wantL: 0.00,
		},
		{
			name: "Pure White",
			r:    255, g: 255, b: 255,
			wantH: 0.0, wantS: 0.00, wantL: 1.00,
		},
		{
			name: "Glay",
			r:    128, g: 128, b: 128,
			wantH: 0.0, wantS: 0.00, wantL: 0.50,
		},
		{
			name: "Skayblue",
			r:    139, g: 184, b: 232,
			wantH: 211.0, wantS: 0.669, wantL: 0.727,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, s, l := RGBToHSL(tt.r, tt.g, tt.b)

			if math.Abs(h-tt.wantH) > 0.01 {
				t.Errorf("RGBToHSL() h = %f, want %f", h, tt.wantH)
			}
			if math.Abs(s-tt.wantS) > 0.01 {
				t.Errorf("RGBToHSL() s = %f, want %f", s, tt.wantS)
			}
			if math.Abs(l-tt.wantL) > 0.01 {
				t.Errorf("RGBToHSL() l = %f, want %f", l, tt.wantL)
			}
		})
	}
}

func TestHSLToRGB(t *testing.T) {
	tests := []struct {
		name                string
		h, s, l             float64
		wantR, wantG, wantB uint8
	}{
		{
			name: "Pure Red",
			h:    0.0, s: 100.0, l: 50.0,
			wantR: 255, wantG: 0, wantB: 0,
		},
		{
			name: "Pure Green",
			h:    120.0, s: 100.0, l: 50.0,
			wantR: 0, wantG: 255, wantB: 0,
		},
		{
			name: "Pure Blue",
			h:    240.0, s: 100.0, l: 50.0,
			wantR: 0, wantG: 0, wantB: 255,
		},
		{
			name: "Pure Black",
			h:    0.0, s: 0.0, l: 0.0,
			wantR: 0, wantG: 0, wantB: 0,
		},
		{
			name: "Pure White",
			h:    0.0, s: 0.0, l: 100.0,
			wantR: 255, wantG: 255, wantB: 255,
		},
		{
			name: "Glay",
			h:    0.0, s: 0.0, l: 50.0,
			wantR: 128, wantG: 128, wantB: 128,
		},
		{
			name: "Skayblue",
			h:    211.0, s: 66.9, l: 72.7,
			wantR: 139, wantG: 184, wantB: 232,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := HSLToRGB(tt.h, tt.s, tt.l)

			if math.Abs(float64(r)-float64(tt.wantR)) > 1 {
				t.Errorf("HSLToRGB() r = %d, want %d", r, tt.wantR)
			}
			if math.Abs(float64(g)-float64(tt.wantG)) > 1 {
				t.Errorf("HSLToRGB() g = %d, want %d", g, tt.wantG)
			}
			if math.Abs(float64(b)-float64(tt.wantB)) > 1 {
				t.Errorf("HSLToRGB() b = %d, want %d", b, tt.wantB)
			}
		})
	}
}
