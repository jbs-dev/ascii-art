package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Mapping of color names to their ANSI escape sequences
var ColorsMap = map[string]string{
	"red":    "\u001b[38;2;255;0;0m",
	"green":  "\u001b[38;2;0;255;0m",
	"yellow": "\u001b[38;2;255;255;0m",
	"blue":   "\u001b[38;2;0;0;255m",
	"purple": "\u001b[38;2;161;32;255m",
	"cyan":   "\u001b[38;2;0;183;235m",
	"white":  "\u001b[38;2;255;255;255m",
	"pink":   "\u001b[38;2;255;0;255m",
	"grey":   "\u001b[38;2;128;128;128m",
	"black":  "\u001b[38;2;0;0;0m",
	"brown":  "\u001b[38;2;160;128;96m",
	"orange": "\u001b[38;2;255;160;16m",
}

func GetRGBColor(color string) (string, error) {
	var selectedColor string

	if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
		rgbValues := strings.TrimPrefix(color, "rgb(")
		rgbValues = strings.TrimSuffix(rgbValues, ")")
		rgbValuesSplit := strings.Split(rgbValues, ",")
		if len(rgbValuesSplit) != 3 {
			return "", fmt.Errorf("invalid RGB color value")
		}
		r, err := strconv.Atoi(strings.TrimSpace(rgbValuesSplit[0]))
		if err != nil {
			return "", fmt.Errorf("invalid RGB color value: %v", err)
		}
		g, err := strconv.Atoi(strings.TrimSpace(rgbValuesSplit[1]))
		if err != nil {
			return "", fmt.Errorf("invalid RGB color value: %v", err)
		}
		b, err := strconv.Atoi(strings.TrimSpace(rgbValuesSplit[2]))
		if err != nil {
			return "", fmt.Errorf("invalid RGB color value: %v", err)
		}
		selectedColor = fmt.Sprintf("\u001b[38;2;%d;%d;%dm", r, g, b)
	}

	return selectedColor, nil
}
