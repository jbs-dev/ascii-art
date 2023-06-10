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

func main() {

	// hasFlagOption := false

	for _, arg := range os.Args {
		switch {
		case strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-"):
			// hasFlagOption = true

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

	/* if !hasFlagOption && (len(os.Args) != 3 || !ban.IsValidBanner(os.Args[2])) {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		os.Exit(0)
	} */

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
		// New case to handle both --output and --color flags
	case *outputFlag != "" && *colorFlag != "":
		inputString, banner := ban.ParseArgs(flag.Args())
		asciiArt, err := ascii_art.Generate(inputString, banner)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if *colorFlag != "" {
			colors := strings.Split(*colorFlag, ",")
			asciiArt, err = color.ApplyColor(asciiArt, colors)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		output.SaveProcess(*outputFlag, asciiArt)

		// New case to handle both --align and --color flags
	case *alignFlag != "" && *colorFlag != "":
		args := flag.Args()
		if len(args) < 1 {
			fmt.Println("Invalid command line arguments. Usage: go run main.go --align=<alignment> --color=<color> <string> [banner]")
			os.Exit(1)
		}
		inputString, banner := ban.ParseArgs(args)
		// Process alignment
		asciiArt, err := align.Process(inputString, *alignFlag, banner)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Process colorization
		colors := strings.Split(*colorFlag, ",")
		asciiArt, err = color.ApplyColor(asciiArt, colors)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(asciiArt)

	case *alignFlag != "":
		args := flag.Args()
		inputString, banner := ban.ParseArgs(args)
		asciiArt, err := align.Process(inputString, *alignFlag, banner)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(asciiArt)

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

	case *outputFlag != "":
		inputString, banner := ban.ParseArgs(flag.Args())
		if !ban.IsValidBanner(banner) {
			fmt.Println("Invalid banner type.")
			os.Exit(1)
		}
		output.Process(*outputFlag, []string{inputString}, banner)

	default:
		inputString, banner := ban.ParseArgs(flag.Args())
		err := ascii_art.TerminalPrint(inputString, banner)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		/* default:
		args := flag.Args()
		if len(args) > 0 {
			inputString := args[0]
			banner := "standard" // set default banner

			if len(args) > 1 {
				if !ban.IsValidBanner(args[1]) {
					fmt.Println("Invalid banner type. Please use a valid banner or leave it blank for standard banner.")
					os.Exit(1)
				}
				banner = args[1] // if banner is specified and valid, use it
			}

			err := ascii_art.TerminalPrint(inputString, banner)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Usage: go run . [STRING] [BANNER]")
			fmt.Println("EX: go run . something standard")
			os.Exit(1)
		} */
	}
}
