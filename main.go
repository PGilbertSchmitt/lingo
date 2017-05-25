package main

import (
	"bufio"
	"log"
	"os"
)

type mode int

const (
	emptyMode mode = iota
	syllableMode
	consonantMode
	vowelMode
)

func whichMode(line string) mode {
	switch line {
	case "# SYLLABLES":
		return syllableMode
	case "# CONSONANTS":
		return consonantMode
	case "# VOWELS":
		return vowelMode
	default:
		return emptyMode
	}
}

type composite int

// These make up a composition
const (
	c composite = 99
	v composite = 118
)

// This represents a syllable's composition
type composition []composite

func main() {
	inputName := os.Args[1]
	if inputName == "" {
		log.Fatal("No file given")
	}
	configFile, err := os.Open(inputName)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	mode := emptyMode
	syllableRules := []composition{}
	consonants := []string{}
	vowels := []string{}

	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			mode = emptyMode
		}
		switch mode {
		case emptyMode:
			mode = whichMode(line)
		case consonantMode:
			consonants = append(consonants, line)
		case vowelMode:
			vowels = append(vowels, line)
		case syllableMode:
			var newRule composition
			for _, char := range line {
				newRule = append(newRule, composite(char))
			}
			syllableRules = append(syllableRules, newRule)
		}
	}
}
