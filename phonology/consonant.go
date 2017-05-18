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
	code    rune
	method  ArticulationMethod
	point   ArticulationPoint
	voicing bool
}

// New returns a new Consonant
func New(method ArticulationMethod, point ArticulationPoint, voicing bool) Consonant {
	return Consonant{
		method:  method,
		point:   point,
		voicing: voicing,
	}
}

// Char returns the unicode code point for the appropriate
// IPA character
func (c Consonant) Char() rune {
	return c.code
}
