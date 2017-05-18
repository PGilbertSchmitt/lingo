package phonology

// Syllable contains the broad definition of a syllable
type Syllable struct {
	onset   []Consonant
	nucleus []Vowel
	coda    []Consonant
}

// Word can be any number of syllables
type Word []Syllable

// PhoneList returns all phones that constitute a word in order
func (w Word) PhoneList() []Phone {
	phones := []Phone{}

	for _, syl := range w {
		for _, c := range syl.onset {
			phones = append(phones, c)
		}

		for _, v := range syl.nucleus {
			phones = append(phones, v)
		}

		for _, c := range syl.coda {
			phones = append(phones, c)
		}
	}

	return phones
}
