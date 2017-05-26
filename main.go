package main

import (
	ph "lingwish/phonology"
	"math/rand"
	"strings"
	"time"

	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func (comp composition) generateSyllable(consonants, vowels ph.WeightedPhones) ph.Syllable {
	copyComp := make(composition, len(comp))
	copy(copyComp, comp)

	var curComposite composite
	afterVowels := false

	var newOnset, newCoda ph.Consonants
	var newNucleus ph.Vowels

	for {
		// Shifting
		curComposite, copyComp = copyComp[0], copyComp[1:]
		if curComposite == c {
			if afterVowels {
				newOnset = append(newOnset, consonants.RandomPhone().(ph.Consonant))
			} else {
				newCoda = append(newCoda, consonants.RandomPhone().(ph.Consonant))
			}
		} else if curComposite == v {
			newNucleus = append(newNucleus, vowels.RandomPhone().(ph.Vowel))
		}
		if len(copyComp) == 0 {
			break
		}
	}

	return ph.NewSyllable(newOnset, newNucleus, newCoda)
}

func main() {
	// To ensure "true" randomness when pulling phones
	rand.Seed(int64(time.Now().Nanosecond()))

	/* Getting all info from input file */

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
	userConsonants := []string{}
	userVowels := []string{}

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
			userConsonants = append(userConsonants, line)
		case vowelMode:
			userVowels = append(userVowels, line)
		case syllableMode:
			var newRule composition
			for _, char := range line {
				newRule = append(newRule, composite(char))
			}
			syllableRules = append(syllableRules, newRule)
		}
	}

	/* Forming phonetic inventory */

	// Consonants
	allConsonants := ph.AllConsonants()

	consonantMap := make(map[string]ph.Consonant)

	for _, consonant := range allConsonants {
		consonantMap[string(consonant.Char())] = consonant
	}

	var weightedConsonants ph.WeightedPhones

	for _, consonant := range userConsonants {
		pieces := strings.Fields(consonant)
		char := pieces[0]
		weight, err := strconv.Atoi(pieces[1])
		if err != nil {
			log.Printf("Could not convert weight: %s", err)
			weight = 0
		}

		weightedConsonants.AddPhone(ph.Phone(consonantMap[char]), weight)
	}
	// weightedConsonants is now able to pull any consonant randomly

	// Vowels
	allVowels := ph.AllVowels()

	vowelMap := make(map[string]ph.Vowel)

	for _, vowel := range allVowels {
		vowelMap[string(vowel.Char())] = vowel
	}

	var weightedVowels ph.WeightedPhones

	for _, vowel := range userVowels {
		pieces := strings.Fields(vowel)
		char := pieces[0]
		weight, err := strconv.Atoi(pieces[1])
		if err != nil {
			log.Printf("Could not convert weight: %s", err)
			weight = 0
		}

		weightedVowels.AddPhone(ph.Phone(vowelMap[char]), weight)
	}
	// weightedVowels is now able to pull any vowel randomly

	/* Generating random syllables */

	for x := 0; x < 30; x++ {
		idx := int(rand.Float64() * float64(len(syllableRules)))
		comp := syllableRules[idx]

		syllable := comp.generateSyllable(weightedConsonants, weightedVowels)
		var word ph.Word
		word = append(word, syllable)
		fmt.Printf("/%s/\n", word.WordString())
	}
}
