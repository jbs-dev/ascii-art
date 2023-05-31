package main

import (
	"fmt"
	"student/color"
)

func main() {
	asciiArt := `
	 _                  _    
	| |                | |   
	| |_    ___   ___  | |_  
	| __|  / _ \ / __| | __| 
	\ |_  |  __/ \__ \ \ |_  
	 \__|  \___| |___/  \__| 
							 
							 
	 `
	colors := []string{"red"}

	result, err := color.ApplyColor(asciiArt, colors)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
