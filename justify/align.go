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

	var builder strings.Builder
	err = createArt(&builder, input, segmentedTemplate, alignment)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}

func createArt(builder *strings.Builder, input string, template [][]string, alignment string) error {
	// log.Println("alignment:", alignment)

	var arr [][]string
	var wordArr []string
	skipChar := false
	newLineCount := 0
	for i, r := range input {
		if skipChar {
			skipChar = false
			continue
		}

		if r == '\\' && len(input) > i+1 && input[i+1] == 'n' {
			skipChar = true
			newLineCount++
			if len(wordArr) != 0 {
				arr = append(arr, wordArr)
				if alignment == "justify" {
					printJustify(builder, arr)
				} else {
					print(builder, arr, alignment)
				}
				arr = [][]string{}
				wordArr = []string{}
			}

			if newLineCount > 1 {
				builder.WriteString("\n")
			}

			continue
		}

		// Reset newLineCount if character isn't a newline
		newLineCount = 0

		// If there's a space, it indicates end of a word
		if r == ' ' {
			if len(wordArr) != 0 {
				arr = append(arr, wordArr)
				wordArr = []string{}
			}
			wordArr = customAppend(wordArr, []string{"        ", "        ", "        ", "        ", "        ", "        ", "        ", "        "})
			continue
		}
		wordArr = customAppend(wordArr, template[int(r)-32])
	}
	if len(wordArr) != 0 {
		arr = append(arr, wordArr)
	}

	if alignment == "justify" {
		printJustify(builder, arr)
	} else {
		print(builder, arr, alignment)
	}
	// log.Println("Final wordArr length:", len(wordArr))
	// log.Println("Final arr length:", len(arr))

	return nil
}

func printJustify(builder *strings.Builder, arr [][]string) {
	n := len(arr)
	m := len(arr[0])
	totalWidth := 0
	for _, wordArr := range arr {
		totalWidth += len(wordArr[0])
	}
	totalSpaces := getTerminalWidth() - totalWidth

	// If there's only one word, it's not justified
	if n == 1 {
		print(builder, arr, "left") // pass the entire arr which is a [][]string
		return
	}

	spacesBetweenWords := totalSpaces / (n - 1)
	remainder := totalSpaces % (n - 1)

	// Print variables here
	// log.Println("n:", n)
	// log.Println("totalWidth:", totalWidth)
	// log.Println("totalSpaces:", totalSpaces)
	// log.Println("spacesBetweenWords:", spacesBetweenWords)
	// log.Println("remainder:", remainder)

	for i := 0; i < m; i++ {
		for j, wordArr := range arr {
			builder.WriteString(wordArr[i])
			if j < n-1 {
				spaces := spacesBetweenWords
				if remainder > 0 {
					spaces++
					remainder--
				}
				builder.WriteString(strings.Repeat(" ", spaces))
			}
		}
		builder.WriteString("\n")
	}
}

func print(builder *strings.Builder, arr [][]string, alignment string) {
	m := len(arr[0])
	width := getTerminalWidth()

	switch alignment {
	case "center":
		totalWidth := 0
		for _, wordArr := range arr {
			totalWidth += len(wordArr[0])
		}
		diff := width/2 - totalWidth/2
		for i := 0; i < m; i++ {
			for j := 0; j < diff; j++ {
				builder.WriteString(" ")
			}
			for _, wordArr := range arr {
				builder.WriteString(wordArr[i])
			}
			builder.WriteString("\n")
		}
	case "right":
		totalWidth := 0
		for _, wordArr := range arr {
			totalWidth += len(wordArr[0])
		}
		diff := width - totalWidth
		for i := 0; i < m; i++ {
			for j := 0; j < diff; j++ {
				builder.WriteString(" ")
			}
			for _, wordArr := range arr {
				builder.WriteString(wordArr[i])
			}
			builder.WriteString("\n")
		}
	default:
		for i := 0; i < m; i++ {
			for _, wordArr := range arr {
				builder.WriteString(wordArr[i])
			}
			builder.WriteString("\n")
		}
	}
}

func customAppend(str, item []string) []string {
	for i := 0; i < 8; i++ {
		if i < len(str) {
			str[i] += item[i]
		} else {
			str = append(str, item[i])
		}
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
	// If we're debugging in VS Code, return a hard-coded terminal width
	if _, ok := os.LookupEnv("VSCODE_GO_DEBUGGING"); ok {
		return 80 // Or whatever value is suitable for your needs
	}

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
