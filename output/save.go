package output

import (
	"fmt"
	"os"
	"path/filepath"
)

// Process function to generate and output ASCII art
func SaveProcess(outputFileName string, asciiArt string) {
	// Check if outputFileName is provided
	if outputFileName != "" {
		// Write the ASCII art to the output file
		outputDirectory := "tests"
		err := os.MkdirAll(outputDirectory, 0755) // 0755 sets read/write permissions for owner and read for others
		if err != nil {
			fmt.Println("Error creating directory:", err)
			os.Exit(1)
		}

		fullPath := filepath.Join(outputDirectory, outputFileName)

		// Write the ASCII art to the file.
		err = os.WriteFile(fullPath, []byte(asciiArt), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}

		// Open the file.
		file, err := os.OpenFile(fullPath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()

		// Get the file size.
		stat, err := file.Stat()
		if err != nil {
			fmt.Println("Error getting file stats:", err)
			os.Exit(1)
		}

		// Check the last byte for a newline.
		lastByte := make([]byte, 1)
		_, err = file.ReadAt(lastByte, stat.Size()-1)
		if err != nil {
			fmt.Println("Error reading from file:", err)
			os.Exit(1)
		}

		// If the last byte is a newline, truncate the file by one byte.
		if lastByte[0] == '\n' {
			err = file.Truncate(stat.Size() - 1)
			if err != nil {
				fmt.Println("Error truncating file:", err)
				os.Exit(1)
			}
		}
	} else {
		// If no outputFileName is provided, print to the terminal
		fmt.Println(asciiArt)
	}
}
