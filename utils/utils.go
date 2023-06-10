package utils

import (
	"fmt"
	"os"
)

var ValidBanners = []string{
	"standard",
	"shadow",
	"thinkertoy",
	"colossal",
	"graffiti",
	"metric",
	"matrix",
	"rev",
	"card",
}

func IsValidBanner(banner string) bool {
	for _, validBanner := range ValidBanners {
		if banner == validBanner {
			return true
		}
	}
	return false
}

// Function to parse string and banner arguments
func ParseArgs(args []string) (string, string) {
	// Set default banner
	banner := "standard"

	if len(args) == 0 {
		fmt.Println("Please provide at least one argument: the string to print.")
		os.Exit(1)
	}

	inputString := args[0]

	// Check for banner argument and validate it
	if len(args) > 1 {
		if !IsValidBanner(args[1]) {
			fmt.Println("Invalid banner type. Please use a valid banner or leave it blank for standard banner.")
			os.Exit(1)
		}
		banner = args[1]
	}

	return inputString, banner
}
