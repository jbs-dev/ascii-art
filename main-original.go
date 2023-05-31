package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"student/ascii_art"
	"student/color"
	align "student/justify"
	"student/output"
	"student/reverse"
	ban "student/utils"
)

func main_test() {
	hasFlagOption := false

	for _, arg := range os.Args {
		switch {
		case strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-"):
			hasFlagOption = true

			switch {
			case strings.HasPrefix(arg, "--reverse") && !strings.Contains(arg, "="):
				fmt.Println("Usage: go run . [OPTION]")
				fmt.Println("EX: go run . --reverse=<fileName>")
				os.Exit(0)
			case strings.HasPrefix(arg, "--color") && !strings.Contains(arg, "="):
				fmt.Println("Usage: go run . [OPTION] [STRING]")
				fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
				os.Exit(0)
			case strings.HasPrefix(arg, "--output") && !strings.Contains(arg, "="):
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
				fmt.Println("EX: go run . --output=<fileName.txt> something standard")
				os.Exit(0)
			case strings.HasPrefix(arg, "--fs") && !strings.Contains(arg, "="):
				fmt.Println("Usage: go run . [STRING] [BANNER]")
				fmt.Println("EX: go run . something standard")
				os.Exit(0)
			case strings.HasPrefix(arg, "--justify") && !strings.Contains(arg, "="):
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
				fmt.Println("Example: go run . --align=right something standard")
				os.Exit(0)
			}
		}
	}

	if !hasFlagOption && (len(os.Args) != 3 || !ban.IsValidBanner(os.Args[2])) {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	}

	reverseFlag := flag.String("reverse", "", "Path to the target file to process")
	outputFlag := flag.String("output", "", "File path to output the ASCII art")
	alignFlag := flag.String("align", "", "Alignment for ASCII art")
	colorFlag := flag.String("color", "", "Color for ASCII art")
	flag.Parse()

	if *colorFlag == "" && strings.HasPrefix(os.Args[1], "--color") {
		fmt.Println("No color specified. Usage: go run . --color=<color> \"something\"")
		os.Exit(1)
	}

	switch {
	case *reverseFlag != "":
		err := reverse.Process(*reverseFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case *outputFlag != "":
		if len(os.Args) < 4 {
			fmt.Println("Invalid command line arguments. Usage: go run main.go --output=<file_path> <string> <banner>")
			os.Exit(1)
		}
		inputString := os.Args[2]
		banner := os.Args[3]
		if !ban.IsValidBanner(banner) {
			fmt.Println("Invalid banner type.")
			os.Exit(1)
		}
		output.Process(*outputFlag, []string{inputString}, banner)

	case *colorFlag != "":
		args := flag.Args()
		if len(args) == 0 {
			fmt.Println("Please provide at least one argument: the string to print.")
			return
		}

		colorWord := args[0]
		sentence := colorWord

		if len(args) > 1 {
			sentence = strings.Join(args[1:], " ")
		}

		colors := strings.Split(*colorFlag, ",")
		err := color.Process(sentence, colors, "", colorWord)
		if err != nil {
			fmt.Println(err)
			return
		}
	case *alignFlag != "":
		args := flag.Args()
		if len(args) != 2 {
			fmt.Println("Invalid command line arguments. Usage: go run main.go --align=<alignment> <string> <banner>")
			os.Exit(1)
		}
		inputString := args[0]
		banner := args[1]
		if !ban.IsValidBanner(banner) {
			fmt.Println("Invalid banner type.")
			os.Exit(1)
		}
		asciiArt, err := align.Process(inputString, *alignFlag, banner)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(asciiArt)

	default:
		args := flag.Args()
		if len(args) > 0 {
			inputString := args[0]
			banner := "standard" // set default banner

			if len(args) > 1 && ban.IsValidBanner(args[1]) {
				banner = args[1] // if banner is specified and valid, use it
			}

			err := ascii_art.TerminalPrint(inputString, banner)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Please provide a valid option or a string to generate ASCII art from.")
			os.Exit(1)
		}
	}
}
