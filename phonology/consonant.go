package phonology

// ArticulationPoint describes where the sound is made
type ArticulationPoint int

// This enum represents all reasonable points of articulation
const (
	Bilabial ArticulationPoint = iota
	Labiodental
	Dental
	Alveolar
	Postalveolar
	Retroflex
	Palatal
	Velar
	Uvular
	Pharyngeal
	Glottal
)

// ArticulationMethod describes how the sound is made
type ArticulationMethod int

// This enum represents all reasonable methods of articulation
const (
	Plosive ArticulationMethod = iota
	Nasal
	Trill
	Flap
	Fricative
	LateralFricative
	Approximant
	LateralApproximant
)

// Consonant represents a minimal pulmonic consonant sound
// (as per Wikipedia's IPA table)
type Consonant struct {
	code   rune
	method ArticulationMethod
	point  ArticulationPoint
	voiced bool
}

// NewConsonant does the obvious
func NewConsonant(code rune, voiced bool, point ArticulationPoint, method ArticulationMethod) Consonant {
	return Consonant{
		code:   code,
		voiced: voiced,
		point:  point,
		method: method,
	}
}

// Char returns the unicode code point for the appropriate
// IPA character
func (c Consonant) Char() rune {
	return c.code
}
