package main

import (
	"student/ascii_art"
	"testing"
)

func TestProcessAscii(t *testing.T) {
	testCases := []struct {
		input  string
		banner string
	}{
		{input: "ABC", banner: "card"},
		{input: "Testing", banner: "colossal"},
		{input: "Testing-123", banner: "graffiti"},
		{input: "Hello, World!", banner: "matrix"},
		{input: "ASCII Art", banner: "metric"},
		{input: "ABC", banner: "rev"},
		{input: "Test1ng-123", banner: "shadow"},
		{input: "Hello, World!", banner: "standard"},
		{input: "ASCII", banner: "thinkertoy"},
	}

	for _, tc := range testCases {
		err := ascii_art.Process(tc.input, tc.banner)
		if err != nil {
			t.Errorf("Process failed for input '%s' and banner '%s': %v", tc.input, tc.banner, err)
		}
	}
}
