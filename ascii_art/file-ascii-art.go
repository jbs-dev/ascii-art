package ascii_art

import (
	"fmt"
	"os"
	"strings"
	"student/utils"
)

// Print ascii-art to file
func Generate(input, banner string) (string, error) {
	var asciiArt strings.Builder
	for _, r := range input {
		if r < ' ' || r > '~' {
			return "", fmt.Errorf("invalid character: %c", r)
		}
	}

	// read the file from the fonts folder
	bytes, err := os.ReadFile(fmt.Sprintf("fonts/%s.txt", banner))
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	var lines []string
	if banner == "thinkertoy" {
		lines = strings.Split(string(bytes), "\r\n")
	} else {
		lines = strings.Split(string(bytes), "\n")
	}

	var arr []rune
	Newline := false

	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" || banner == "colossal" {
		for i, r := range input {
			if Newline {
				Newline = false
				art(&asciiArt, arr, lines)
				arr = []rune{}
				continue
			}
			if r == '\\' && len(input) != i+1 {
				if input[i+1] == 'n' {
					Newline = true
					continue
				}
			}
			arr = append(arr, r)
		}
		art(&asciiArt, arr, lines)
	} else {
		var lineCount int
		var offset int

		bannerDetails, err := utils.GetBannerDetails(banner)
		if err != nil {
			return "", fmt.Errorf("unknown banner: %v", err)
		}

		lineCount = bannerDetails.LineCount
		offset = bannerDetails.Offset

		for i, r := range input {
			if Newline {
				Newline = false
				printArt(&asciiArt, arr, lines, lineCount, offset)
				arr = []rune{}
				continue
			}

			if r == '\\' && i != len(input)-1 {
				if input[i+1] == 'n' {
					Newline = true
					continue
				}
			}
			arr = append(arr, r)
		}
		printArt(&asciiArt, arr, lines, lineCount, offset)
	}
	// fmt.Print(asciiArt.String()) commented out to check the
	return asciiArt.String(), nil
}

func printArt(asciiArt *strings.Builder, arr []rune, lines []string, lineCount int, offset int) {
	if len(arr) != 0 {
		for line := 1; line <= lineCount; line++ {
			for _, r := range arr {
				skip := (r * rune(lineCount)) - rune(offset)
				asciiArt.WriteString(lines[line+int(skip)])
			}
			fmt.Fprintln(asciiArt)
		}
	} else {
		fmt.Fprintln(asciiArt)
	}
}

func art(asciiArt *strings.Builder, arr []rune, lines []string) {
	if len(arr) != 0 {
		for line := 1; line <= 8; line++ {
			for _, r := range arr {
				skip := (r - 32) * 9
				asciiArt.WriteString(lines[line+int(skip)])
			}
			fmt.Fprintln(asciiArt)
		}
	} else {
		fmt.Fprintln(asciiArt)
	}
}
