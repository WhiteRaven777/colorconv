package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch strings.ToLower(cmd) {
	case "rgb":
		handleRGB(args)
	case "hsl":
		handleHSL(args)
	case "hsv":
		handleHSV(args)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Usage:
  colorconv rgb <red> <green> <blue>
  colorconv rgb <hex>
  colorconv hsl <hue> <saturation> <lightness>
  colorconv hsv <hue> <saturation> <value>

Note:
  - <red>, <green>, <blue>: integers from 0 to 255
  - <hex>: 6-digit hex (e.g., '#ff00ff' or 'ff00ff')
  - <hue>: degrees from 0.0 to 360.0
  - <saturation>, <lightness>, <value>: percentage from 0.0 to 100.0`)
}

func handleRGB(args []string) {
	switch len(args) {
	case 1:
		// HEX mode
		hex := strings.TrimPrefix(args[0], "#")
		if len(hex) != 6 {
			fmt.Println("Invalid HEX format. Use 6-digit hex (e.g., '#ff00ff' or 'ff00ff')")
			os.Exit(1)
		}
		var r, g, b uint8
		if _, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b); err != nil {
			fmt.Println("Invalid HEX value.")
			os.Exit(1)
		} else {
			printResult(r, g, b)
		}
	case 3:
		// RGB mode
		r := parseUint8Arg(args[0], "Red")
		g := parseUint8Arg(args[1], "Green")
		b := parseUint8Arg(args[2], "Blue")
		printResult(r, g, b)
	default:
		fmt.Println("Invalid arguments for 'rgb'.")
		printUsage()
		os.Exit(1)
	}
}

func handleHSL(args []string) {
	switch len(args) {
	case 3:
		// HSL mode
		h := parseFloat64Arg(args[0], "Hue", 0.0, 360.0)
		s := parseFloat64Arg(args[1], "Saturation", 0.0, 100.0)
		l := parseFloat64Arg(args[2], "Lightness", 0.0, 100.0)

		printResult(HSLToRGB(h, s, l))
	default:
		fmt.Println("Invalid arguments for 'hsl'.")
		printUsage()
		os.Exit(1)
	}
}

func handleHSV(args []string) {
	switch len(args) {
	case 3:
		// HSV mode
		h := parseFloat64Arg(args[0], "Hue", 0.0, 360.0)
		s := parseFloat64Arg(args[1], "Saturation", 0.0, 100.0)
		v := parseFloat64Arg(args[2], "Value", 0.0, 100.0)

		printResult(HSVToRGB(h, s, v))
	default:
		fmt.Println("Invalid arguments for 'hsv'.")
		printUsage()
		os.Exit(1)
	}
}

func printResult(r, g, b uint8) {
	h, s, l := RGBToHSL(r, g, b)
	hsvH, hsvS, hsvV := RGBToHSV(r, g, b)
	fmt.Printf(`RGB: %d, %d, %d [ #%02x%02x%02x ]
HSL: %.1f, %.3f, %.3f [ hsl(%.1f, %.1f%%, %.1f%%) ]
HSV: %.1f, %.3f, %.3f [ hsv(%.1f, %.1f%%, %.1f%%) ]
`,
		r, g, b, r, g, b,
		h, s, l, h, s*100, l*100,
		hsvH, hsvS, hsvV, hsvH, hsvS*100, hsvV*100)
}

func parseUint8Arg(arg, name string) uint8 {
	v, err := strconv.ParseUint(arg, 10, 8)
	if err != nil {
		fmt.Printf("Invalid %s value: %s\n", name, arg)
		os.Exit(1)
	}
	return uint8(v)
}

func parseFloat64Arg(arg, name string, min, max float64) float64 {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("Invalid %s value: %s\n", name, arg)
		os.Exit(1)
	}
	if v < min || v > max {
		fmt.Printf("%s value out of range. Must be between %.2f and %.2f.\n", name, min, max)
		os.Exit(1)
	}
	return v
}
