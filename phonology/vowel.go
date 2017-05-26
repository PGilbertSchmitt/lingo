package phonology

// Roundness
type Roundness int

const (
	Nilrounding Roundness = iota
	Rounded
	Unrounded
)

// Frontness describes the location of a vowel sound
// in the mound
type Frontness int

// All reasonble mouth locations
const (
	Front Frontness = iota
	NearFront
	Central
	NearBack
	Back
)

// Openness describes how open the mount is during
// pronounciation
type Openness int

// All reasonable levels of openness
const (
	Close Openness = iota
	NearClose
	CloseMid
	Mid
	OpenMid
	NearOpen
	Open
)

// Vowel describes all reasonable vowel sounds
type Vowel struct {
	code      rune
	frontness Frontness
	openness  Openness
	rounded   Roundness
}

// Vowels is self explanitory
type Vowels []Vowel

// NewVowel does what you think it does
func NewVowel(code rune, rounded int, frontness Frontness, openness Openness) Vowel {
	return Vowel{
		code:      code,
		rounded:   Roundness(rounded),
		frontness: frontness,
		openness:  openness,
	}
}

// Char returns the unicode code point for the appropriate
// IPA character
func (v Vowel) Char() rune {
	return v.code
}
