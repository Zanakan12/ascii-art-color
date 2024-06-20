package colors


import (
	"fmt"
	"math/rand" 
	"regexp"
	"strconv"
	"strings"
	"time"
)

////////////////////////////////////////////////////// color ///////////////////////////////////////////////////////////////////////////////////
const (
	Esc = "\033["

	// Reset
	Reset = "\033[0m"

	// Regular Colors
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Purple  = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Orange  = "\033[38;2;255;165;0m"
	Magenta = "\033[35m"
	Teal    = "\033[38;2;0;128;128m"
	Olive   = "\033[38;2;128;128;0m"
)

// Regular expressions for matching hex, RGB, and HSL color formats
var hexColorRegex = regexp.MustCompile(`^#?([a-fA-F0-9]{6})$`)
var rgbColorRegex = regexp.MustCompile(`^rgb\((\d{1,3}),\s*(\d{1,3}),\s*(\d{1,3})\)$`)
var hslColorRegex = regexp.MustCompile(`^hsl\((\d{1,3}),\s*(\d{1,3})%,\s*(\d{1,3})%\)$`)

// Converts a hex color code to RGB values
func hexToRGB(hex string) (int, int, int) {
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return int(r), int(g), int(b)
}

// Converts HSL color format to RGB values
func hslToRGB(h, s, l int) (int, int, int) {
	hF := float64(h) / 360
	sF := float64(s) / 100
	lF := float64(l) / 100

	var rF, gF, bF float64
	if s == 0 {
		rF, gF, bF = lF, lF, lF // achromatic
	} else {
		q := lF * (1 + sF)
		if lF >= 0.5 {
			q = lF + sF - lF*sF
		}
		p := 2*lF - q
		rF = hueToRGB(p, q, hF+1/3.0)
		gF = hueToRGB(p, q, hF)
		bF = hueToRGB(p, q, hF-1/3.0)
	}

	return int(rF * 255), int(gF * 255), int(bF * 255)
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
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

// Generates a random color in RGB format
func randomColor() (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(256), rand.Intn(256), rand.Intn(256)
}

// Applies the specified color to the text
func ApplyColor(text, color string) string {
	colorCode := ""
	switch strings.ToLower(color) {
	case "black":
		colorCode = Black
	case "red":
		colorCode = Red
	case "green":
		colorCode = Green
	case "yellow":
		colorCode = Yellow
	case "blue":
		colorCode = Blue
	case "purple":
		colorCode = Purple
	case "cyan":
		colorCode = Cyan
	case "white":
		colorCode = White
	case "orange":
		colorCode = Orange
	case "magenta":
		colorCode = Magenta
	case "teal":
		colorCode = Teal
	case "olive":
		colorCode = Olive
	case "random":
		r, g, b := randomColor()
		colorCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	default:
		// Check for hex color format
		if match := hexColorRegex.FindStringSubmatch(color); match != nil {
			r, g, b := hexToRGB(match[1])
			colorCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		// Check for RGB color format
		} else if match := rgbColorRegex.FindStringSubmatch(color); match != nil {
			r, _ := strconv.Atoi(match[1])
			g, _ := strconv.Atoi(match[2])
			b, _ := strconv.Atoi(match[3])
			colorCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		// Check for HSL color format
		} else if match := hslColorRegex.FindStringSubmatch(color); match != nil {
			h, _ := strconv.Atoi(match[1])
			s, _ := strconv.Atoi(match[2])
			l, _ := strconv.Atoi(match[3])
			r, g, b := hslToRGB(h, s, l)
			colorCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		} else {
			return text // If no match, return the original text
		}
	}
	return colorCode + text + Reset
}

