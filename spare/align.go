package align

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Process creates ASCII art from the input string using the specified banner font,
// aligns the resulting ASCII art according to the alignment parameter,
// and returns the final ASCII art as a string.
func Process(input string, alignment string, banner string) (string, error) {
	for _, r := range input {
		if r < ' ' || r > '~' {
			return "", fmt.Errorf("invalid character: %c", r)
		}
	}

	bannerTemplate, err := GetAllLines(fmt.Sprintf("fonts/%s.txt", banner))
	if err != nil {
		return "", err
	}

	segmentedBanner := segment(bannerTemplate)
	art, err := createArt(input, segmentedBanner, alignment)
	if err != nil {
		return "", err
	}

	return art, nil
}

// createArt transforms the input string into ASCII art and aligns it.
func createArt(input string, bannerTemplate [][]string, alignment string) (string, error) {
	var arr [][]string
	var result string

	wordArr := make([]string, 8)
	newline := false

	// Process each character in the input string.
	for i, r := range input {
		if newline {
			newline = false
			continue
		}

		// If we encounter a '\n', add the current word to the list, add a newline, and start a new word.
		if r == '\\' && len(input) != i+1 && input[i+1] == 'n' {
			arr, wordArr, result = processWord(arr, wordArr, alignment, result)
			arr = append(arr, []string{"\n"})
			newline = true
			continue
		}

		// If we encounter a space, add the current word to the list, add a space block, and start a new word.
		if r == ' ' {
			arr, wordArr, result = processWord(arr, wordArr, alignment, result)
			arr = append(arr, makeSpace())
			continue
		}

		// Append the ASCII art for the current character to wordArr.
		wordArr = appendASCIIArt(wordArr, bannerTemplate[int(r)-32])
	}

	// Add the last word to the list.
	arr = append(arr, wordArr)

	// Generate the output from the list of words.
	result = generateOutput(arr, alignment, result)

	return result, nil
}

// processWord adds the current word to the list of ASCII art words and returns a new word array.
func processWord(arr [][]string, wordArr []string, alignment string, result string) ([][]string, []string, string) {
	arr = append(arr, wordArr)

	// If the alignment is "justify", print the current line immediately.
	if alignment == "justify" {
		result += printJustify(arr)
		arr = nil
	}

	// Start a new word.
	wordArr = make([]string, 8)

	return arr, wordArr, result
}

// generateOutput concatenates the ASCII art words into a single string.
func generateOutput(arr [][]string, alignment string, result string) string {
	if alignment == "justify" {
		// If the alignment is "justify", print the last line.
		result += printJustify(arr)
	} else {
		// Otherwise, print all lines.
		for _, wordArr := range arr {
			result += print(wordArr, alignment)
		}
	}

	return result
}

// makeSpace creates an ASCII art space.
func makeSpace() []string {
	return []string{"          ", "          ", "          ", "          ", "          ", "          ", "          ", "          "}
}

// appendASCIIArt appends the ASCII art for a character to the current word.
func appendASCIIArt(wordArr []string, asciiArt []string) []string {
	for i := 0; i < 8; i++ {
		wordArr[i] += asciiArt[i]
	}

	return wordArr
}

func print(str []string, alignment string) string {
	var result string

	// If the string is a newline, just print a newline
	if len(str) == 1 && str[0] == "\n" {
		return "\n"
	}

	n := len(str)
	artlen := len(str[0])
	width := getTerminalWidth()

	switch alignment {
	case "center":
		diff := width/2 - artlen/2
		for i := 0; i < n; i++ {
			for j := 0; j < diff; j++ {
				result += " "
			}
			result += str[i] + "\n"
		}
	case "right":
		diff := width - artlen
		for i := 0; i < n; i++ {
			for j := 0; j < diff; j++ {
				result += " "
			}
			result += str[i] + "\n"
		}
	case "left":
		for i := 0; i < n; i++ {
			result += str[i] + "\n"
		}
	}
	return result
}

func segment(template []string) [][]string {
	var result [][]string
	for i := 0; i < len(template)-1; i = i + 9 {
		temp := template[i+1 : i+9]
		result = append(result, temp)
	}
	return result
}

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	width, _ := strconv.Atoi(strings.Fields(string(out))[1])
	return width
}

func GetAllLines(fontFileName string) ([]string, error) {
	file, err := os.Open(fontFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func printJustify(arr [][]string) string {
	var result string
	n := len(arr)
	if n == 0 {
		return result
	}
	m := len(arr[0])
	totalWidth := 0
	for _, wordArr := range arr {
		totalWidth += len(wordArr[0])
	}
	totalSpaces := getTerminalWidth() - totalWidth
	spacePerWord := totalSpaces / (n - 1)
	for i := 0; i < m; i++ {
		for j, wordArr := range arr {
			result += wordArr[i]
			if j < n-1 {
				result += strings.Repeat(" ", spacePerWord)
			}
		}
		result += "\n"
	}
	return result
}
