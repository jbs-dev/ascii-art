package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"student/ascii_art"
	"student/color"
	"student/justify"
	"student/output"
	"student/reverse"
)

var validBanners = []string{
	"standard",
	"shadow",
	"thinkertoy",
	"colossal",
	"graffiti",
	"cards",
	"isometric",
	"matrix",
}

func main() {
	if len(os.Args) != 3 || !isValidBanner(os.Args[2]) {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	}

	reverseFlag := flag.String("reverse", "", "Path to the target file to process")
	outputFlag := flag.String("output", "", "File path to output the ASCII art")
	colorFlag := flag.String("color", "", "Color for ASCII art")
	alignFlag := flag.String("align", "", "Alignment for ASCII art")
	flag.Parse()

	switch {
	case *reverseFlag != "":
		err := reverse.Process(*reverseFlag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case *outputFlag != "":
		output.Process(flag.Args(), *outputFlag)
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
		err := justify.Process()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		args := flag.Args()
		if len(args) > 0 {
			inputString := args[0]
			banner := "standard" // set default banner

			if len(args) > 1 && isValidBanner(args[1]) {
				banner = args[1] // if banner is specified and valid, use it
			}

			err := ascii_art.Process(inputString, banner)
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

func isValidBanner(banner string) bool {
	for _, validBanner := range validBanners {
		if banner == validBanner {
			return true
		}
	}
	return false
}
