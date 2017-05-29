package phonology

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

// A Composite is either a 'c' or 'v', and represent the breakup of a syllable
type Composite int

// Two possible values
const (
	C Composite = 99  // Byte for 'c'
	V Composite = 118 // Byte for 'v'
)

// A Composition is just a series of Composites
type Composition []Composite

// GenerateSyllable takes a weighted set of consonants and vowels, and mixes them based on a given composition
func (comp Composition) GenerateSyllable(consonants, vowels WeightedPhones) Syllable {
	copyComp := make(Composition, len(comp))
	copy(copyComp, comp)

	var curComposite Composite
	afterVowels := false

	var newOnset, newCoda Consonants
	var newNucleus Vowels

	for {
		// Shifting
		curComposite, copyComp = copyComp[0], copyComp[1:]
		if curComposite == C {
			if afterVowels {
				newCoda = append(newCoda, consonants.RandomPhone().(Consonant))
			} else {
				newOnset = append(newOnset, consonants.RandomPhone().(Consonant))
			}
		} else if curComposite == V {
			newNucleus = append(newNucleus, vowels.RandomPhone().(Vowel))
		}
		if len(copyComp) == 0 {
			break
		}
	}

	return NewSyllable(newOnset, newNucleus, newCoda)
}
