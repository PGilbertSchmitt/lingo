package phonology

import (
	"bytes"
)

// Syllable contains the broad definition of a syllable
type Syllable struct {
	onset   Consonants
	nucleus Vowels
	coda    Consonants
}

// NewSyllable creates a new syllable from its constituent parts
func NewSyllable(onset Consonants, nucleus Vowels, coda Consonants) Syllable {
	return Syllable{
		onset:   onset,
		nucleus: nucleus,
		coda:    coda,
	}
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

func (w Word) WordString() string {
	phoneList := w.PhoneList()
	var buffer bytes.Buffer
	var phone Phone

	for len(phoneList) > 0 {
		phone, phoneList = phoneList[0], phoneList[1:]
		buffer.WriteString(string(phone.Char()))
	}

	return buffer.String()
}
