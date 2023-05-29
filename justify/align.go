package align

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Process(input string, alignment string, banner string) (string, error) {
	for _, r := range input {
		if r < ' ' || r > '~' {
			return "", fmt.Errorf("invalid character: %c", r)
		}
	}

	template, err := GetAllLines(fmt.Sprintf("fonts/%s.txt", banner))

	if err != nil {
		return "", err
	}
	segmentedTemplate := segment(template)
	art, err := createArt(input, segmentedTemplate, alignment)
	if err != nil {
		return "", err
	}
	return art, nil
}

func createArt(input string, template [][]string, alignment string) (string, error) {
	var arr [][]string
	var result string
	wordArr := []string{"", "", "", "", "", "", "", ""}
	skipNext := false
	for i, r := range input {
		if skipNext {
			skipNext = false
			continue
		}
		if r == '\\' && len(input) > i+1 && input[i+1] == 'n' {
			arr = append(arr, wordArr)
			if alignment == "justify" {
				result += printJustify(arr)
			} else {
				for _, wordArr := range arr {
					result += print(wordArr, alignment)
				}
			}
			arr = [][]string{}
			wordArr = []string{"", "", "", "", "", "", "", ""}
			skipNext = true
			continue
		}
		if r == ' ' {
			arr = append(arr, wordArr)
			wordArr = []string{"          ", "          ", "          ", "          ", "          ", "          ", "          ", "          "}
			arr = append(arr, wordArr)
			wordArr = []string{"", "", "", "", "", "", "", ""}
			continue
		}
		wordArr = customAppend(wordArr, template[int(r)-32])
	}
	arr = append(arr, wordArr)

	lineArr := make([]string, 8)
	if alignment == "justify" {
		result += printJustify(arr)
	} else {
		for _, wordArr := range arr {
			for i := range lineArr {
				lineArr[i] += wordArr[i]
			}
		}
		result += print(lineArr, alignment)
	}
	return result, nil
}

func print(str []string, alignment string) string {
	var result string
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

func customAppend(str, item []string) []string {
	if str == nil {
		str = make([]string, 8)
	}
	for i := 0; i < 8; i++ {
		str[i] = str[i] + item[i]
	}
	return str
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
