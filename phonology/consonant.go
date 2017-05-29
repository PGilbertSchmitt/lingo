package phonology

// Voicing is the voicing of a consonant
type Voicing int

// The different voicings of a consonant. Nilvoice represents irrelevant voicing
const (
	Nilvoice Voicing = iota
	Voiced
	Unvoiced
)

// ArticulationPoint describes where the sound is made
type ArticulationPoint int

// This enum represents all reasonable points of articulation. Nopoint
// is used for filtering
const (
	Nopoint ArticulationPoint = iota
	Bilabial
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

// This enum represents all reasonable methods of articulation. Nomethod
// is used for filtering
const (
	Nomethod ArticulationMethod = iota
	Plosive
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
	voiced Voicing
}

// Consonants is extra self explanitory
type Consonants []Consonant

// NewConsonant does the obvious
func NewConsonant(code rune, voiced int, point ArticulationPoint, method ArticulationMethod) Consonant {
	return Consonant{
		code:   code,
		voiced: Voicing(voiced),
		point:  point,
		method: method,
	}
}

// Char returns the unicode code point for the appropriate
// IPA character
func (c Consonant) Char() rune {
	return c.code
}

// Filter takes attributes of the method or point of articulation, and voicing, and returns all consonants that share that attribute. Using all three attributes guarantees that you will only receive one consonant. This assumes the phonebank is untouched, as there are no two phones in that file that share all three attributes. This will change if more attributes are considered, such as aspiration.
func (consonants Consonants) Filter(attrs ...interface{}) Consonants {
	// Zero values indicate that the specific attribute is not set
	var point ArticulationPoint
	var method ArticulationMethod
	var voicing Voicing

	for _, attr := range attrs {
		switch attr.(type) {
		case Voicing:
			voicing = attr.(Voicing)
		case ArticulationPoint:
			point = attr.(ArticulationPoint)
		case ArticulationMethod:
			method = attr.(ArticulationMethod)
		}
	}

	var filtered Consonants

	for _, c := range consonants {
		// Filters
		if point > 0 && c.point != point {
			continue
		}
		if method > 0 && c.method != method {
			continue
		}
		if voicing > 0 && c.voiced != voicing {
			continue
		}

		filtered = append(filtered, c)
	}

	return filtered
}
