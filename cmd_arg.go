package main

import (
	"errors"
	"flag"
)

type options struct {
	filename      string
	wordCount     int
	syllableCount int
}

// grab() parses the arguments and returns a filled settings struct.
func grabOptions() (options, error) {
	filename := flag.String("f", "", "Name of configuration file (Required)")
	wordCount := flag.Int("w", 30, "Number of words to generate")
	syllableCount := flag.Int("s", 1, "Number of syllables to generate per word")

	flag.Parse()

	if *filename == "" {
		return options{}, errors.New("No file given")
	}

	return options{
		filename:      *filename,
		wordCount:     *wordCount,
		syllableCount: *syllableCount,
	}, nil
}
