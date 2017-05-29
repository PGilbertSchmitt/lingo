package phonology

import (
	"bytes"
)

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
