package phonology

// Phone covers both vowels and consonants
type Phone interface {
	Char() rune
}
