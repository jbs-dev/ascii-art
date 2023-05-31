package color

import (
	"strings"
	"student/utils"
)

func ApplyColor(input string, colors []string) (string, error) {
	var selectedColors []string

	for _, color := range colors {
		var selectedColor string
		var err error
		if strings.HasPrefix(color, "rgb(") && strings.HasSuffix(color, ")") {
			selectedColor, err = utils.GetRGBColor(color)
			if err != nil {
				return "", err
			}
		} else {
			if val, ok := utils.ColorsMap[color]; ok {
				selectedColor = val
			} else {
				selectedColor = colorReset
			}
		}

		selectedColors = append(selectedColors, selectedColor)
	}
	// Create color queue
	colorQueue := NewQueue()

	// Push all colors into queue
	for _, color := range selectedColors {
		colorQueue.Push(color)
	}

	coloredArt := colorArt(input, colorQueue)

	return coloredArt, nil
}

func colorArt(input string, colorQueue *Queue) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		color := colorQueue.Pop().(string)
		colorQueue.Push(color) // Requeue the color
		lines[i] = color + line
	}
	return strings.Join(lines, "\n")
}
