package main

import (
	"student/ascii_art"
	"testing"
)

func TestProcess(t *testing.T) {
	input := "Hello, World!"
	banner := "standard"

	err := ascii_art.Process(input, banner)
	if err != nil {
		t.Errorf("Process failed: %v", err)
	}
}
