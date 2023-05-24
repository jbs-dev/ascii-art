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
		{input: "Hello, World!", banner: "standard"},
		{input: "Test1ng-123", banner: "shadow"},
		{input: "ASCII", banner: "thinkertoy"},
		{input: "Hello", banner: "standard"},
		{input: "Testing", banner: "colossal"},
		{input: "ASCII Art", banner: "metric"},
		{input: "Hello, World!", banner: "matrix"},
		{input: "Testing-123", banner: "graffiti"},
		{input: "ABC", banner: "rev"},
	}

	for _, tc := range testCases {
		err := ascii_art.Process(tc.input, tc.banner)
		if err != nil {
			t.Errorf("Process failed for input '%s' and banner '%s': %v", tc.input, tc.banner, err)
		}
	}
}
